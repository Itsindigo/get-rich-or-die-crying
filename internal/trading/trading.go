package trading

import (
	"errors"
	"fmt"
)

type TraderAPI interface {
	GetWallets([]CoinbaseWalletName) ([]SimpleAccount, error)
	MarketBuy() (interface{}, error)
	MarketSell() (interface{}, error)
}

type TradeMaker struct {
	FearAndGreedScore int
	API               TraderAPI
}

type TradeMakerOptions struct {
	FearAndGreedScore int
	API               TraderAPI
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

func (tm *TradeMaker) SummariseNoAction() {
	fmt.Printf("Today's F&G score was %d, no trade was made.\n", tm.FearAndGreedScore)
}

func (tm *TradeMaker) SellEthGbp(walletPair EthGbpWallet) error {
	if walletPair.Eth.Balance == 0 {
		return errors.New("SellEthGbp: cannot sell ETH as balance is 0")
	}
	_, _ = tm.API.MarketSell()
	return nil
}

func (tm *TradeMaker) BuyEthGbp(walletPair EthGbpWallet) error {
	if walletPair.Gbp.Balance == 0 {
		return errors.New("BuyEthGbp: cannot sell GBP as balance is 0")
	}
	_, _ = tm.API.MarketBuy()

	return nil
}

type ActOptions struct {
	ForceSell bool
	ForceBuy  bool
}

func (tm *TradeMaker) Act(options ActOptions) error {
	walletsToQuery := []CoinbaseWalletName{ETHWallet, GBPWallet}
	wallets, err := tm.API.GetWallets(walletsToQuery)

	if err != nil {
		return err
	}

	walletPair, err := tm.GetEthGbpWallets(wallets)

	if err != nil {
		return err
	}

	if GreedSellThreshold <= tm.FearAndGreedScore || options.ForceSell {
		err := tm.SellEthGbp(walletPair)
		if err != nil {
			return err
		}
		return nil
	}

	if FearBuyThreshold >= tm.FearAndGreedScore || options.ForceBuy {
		tm.BuyEthGbp(walletPair)
		return nil
	}

	if FearBuyThreshold < tm.FearAndGreedScore && GreedSellThreshold > tm.FearAndGreedScore {
		tm.SummariseNoAction()
		return nil
	}

	return nil
}

func NewTradeMaker(options TradeMakerOptions) *TradeMaker {
	return &TradeMaker{FearAndGreedScore: options.FearAndGreedScore, API: options.API}
}
