package trading

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math"
	"strconv"
)

type TraderAPI interface {
	GetWallets([]CoinbaseWalletName) ([]SimpleAccount, error)
	MarketBuy(MarketPair, string) (CreateOrderResponse, error)
	MarketSell(MarketPair, string) (CreateOrderResponse, error)
}

type TradeReporter interface {
	ReportNoAction(context.Context, int) error
	ReportSale(context.Context, string) error
	ReportPurchase(context.Context, string) error
	ReportError(context.Context, error) error
}

type TradeMakerOptions struct {
	API           TraderAPI
	TradeReporter TradeReporter
	MakeMinTrades bool
}

type TradeMaker struct {
	API           TraderAPI
	TradeReporter TradeReporter
	MakeMinTrades bool
}

type RequiredWallets struct {
	GBPWallet *SimpleAccount
	ETHWallet *SimpleAccount
}

type EthGbpWallet struct {
	Eth *SimpleAccount
	Gbp *SimpleAccount
}

func (tm *TradeMaker) GetEthGbpWallets(wallets []SimpleAccount) (EthGbpWallet, error) {
	walletPair := EthGbpWallet{}

	for _, wallet := range wallets {
		if wallet.Name == GBPWallet && wallet.GoodToTrade() {
			walletPair.Gbp = &wallet
		}
		if wallet.Name == ETHWallet && wallet.GoodToTrade() {
			walletPair.Eth = &wallet
		}
	}

	if walletPair.Gbp == nil {
		return EthGbpWallet{}, errors.New("GetEthGbpWallets: GBP Wallet Not Found and/or not in valid state to trade")
	}

	if walletPair.Eth == nil {
		return EthGbpWallet{}, errors.New("GetEthGbpWallets: ETH Wallet Not Found and/or not in valid state to trade")
	}

	return walletPair, nil
}

func (tm *TradeMaker) getSaleAmount(balance string) (string, error) {
	floatBalance, err := strconv.ParseFloat(balance, 64)

	if err != nil {
		return "", err
	}

	// Sell slightly below account balance as sometimes reports insufficient funds.
	// Exact prceision of this amount should be calibrated vs currency. %.5f works for ETH.
	floatBalance -= 0.00001

	// Return zero if subtraction takes balance below zero.
	floatBalance = math.Max(floatBalance, 0)

	if floatBalance == 0 {
		return "", errors.New("`SellEthGbp` cannot sell ETH due to insufficient balance")
	}

	if tm.MakeMinTrades {
		return "0.00010", nil // Approx £0.10 - £0.20 worth of ETH.
	}

	return fmt.Sprintf("%.5f", math.Max(floatBalance, 0)), nil
}

func (tm *TradeMaker) getPurchaseAmount(balance string) (string, error) {
	floatBalance, err := strconv.ParseFloat(balance, 64)

	// £1.00 seems to be min trade amount Coinbase will allow.
	minTradeAmount := 1.00

	if err != nil {
		return "", err
	}

	// Spend £0.01 less than balance as sometimes reports insufficient funds at max.
	floatBalance -= 0.01

	// If balance has gone below zero, return 0.
	floatBalance = math.Max(floatBalance, 0)

	if floatBalance == 0 || floatBalance < minTradeAmount {
		return "", ErrInsufficientEthGbp
	}

	if tm.MakeMinTrades {
		return strconv.FormatFloat(minTradeAmount, 'f', 2, 64), nil
	}

	return fmt.Sprintf("%.2f", floatBalance), nil
}

func (tm *TradeMaker) SellEthGbp(walletPair EthGbpWallet) (string, error) {
	saleAmount, err := tm.getSaleAmount(walletPair.Eth.Balance)

	if err != nil {
		return "", err
	}

	_, err = tm.API.MarketSell(ETH_GBP, saleAmount)

	if err != nil {
		return "", err
	}

	slog.Info("Sold ETH", slog.String("ETH-sold", saleAmount))

	return saleAmount, nil
}

func (tm *TradeMaker) BuyEthGbp(walletPair EthGbpWallet) (string, error) {
	purchaseAmount, err := tm.getPurchaseAmount(walletPair.Gbp.Balance)

	if err != nil {
		return "", err
	}

	_, err = tm.API.MarketBuy(ETH_GBP, purchaseAmount)

	if err != nil {
		return "", err
	}

	slog.Info("Bought ETH", slog.String("GBP-spent", purchaseAmount))

	return purchaseAmount, nil
}

type ActOptions struct {
	ForceSell         bool
	ForceBuy          bool
	FearAndGreedScore int
}

func (tm *TradeMaker) Act(ctx context.Context, options ActOptions) error {
	if options.ForceSell && options.ForceBuy {
		return errors.New("`ForceSell` and `ForceBuy` are both set to `true`, does not make sense to trade when both are true")
	}

	walletsToQuery := []CoinbaseWalletName{ETHWallet, GBPWallet}
	wallets, err := tm.API.GetWallets(walletsToQuery)

	if err != nil {
		return err
	}

	walletPair, err := tm.GetEthGbpWallets(wallets)

	if err != nil {
		return err
	}

	if GreedSellThreshold <= options.FearAndGreedScore || options.ForceSell {
		saleAmount, err := tm.SellEthGbp(walletPair)
		if err != nil {
			return err
		}
		tm.TradeReporter.ReportSale(ctx, saleAmount)
		return nil
	}

	if FearBuyThreshold >= options.FearAndGreedScore || options.ForceBuy {
		purchaseAmount, err := tm.BuyEthGbp(walletPair)

		if err != nil {
			return err
		}

		tm.TradeReporter.ReportPurchase(ctx, purchaseAmount)
		return nil
	}

	if FearBuyThreshold < options.FearAndGreedScore && GreedSellThreshold > options.FearAndGreedScore {
		err := tm.TradeReporter.ReportNoAction(ctx, options.FearAndGreedScore)

		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func NewTradeMaker(options TradeMakerOptions) *TradeMaker {
	return &TradeMaker{
		MakeMinTrades: options.MakeMinTrades,
		API:           options.API,
		TradeReporter: options.TradeReporter,
	}
}
