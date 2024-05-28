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

func (tm *TradeMaker) assertWalletsFound(wallets []SimpleAccount) error {
	foundWallets := make(map[CoinbaseWalletName]*SimpleAccount)

	for _, wallet := range wallets {
		if wallet.Name == GBPWallet {
			foundWallets[GBPWallet] = &wallet
		}
		if wallet.Name == ETHWallet {
			foundWallets[ETHWallet] = &wallet
		}
	}

	if foundWallets[GBPWallet] == nil {
		return errors.New("assertWalletsFound: GBP Wallet Not Found")
	}

	if foundWallets[ETHWallet] == nil {
		return errors.New("assertWalletsFound: ETH Wallet Not Found")
	}

	return nil
}

func (tm *TradeMaker) SummariseNoAction() {
	fmt.Printf("Today's F&G score was %d, no trade was made.\n", tm.FearAndGreedScore)
}

func (tm *TradeMaker) Sell() {
	_, _ = tm.API.MarketSell()
}

func (tm *TradeMaker) Buy() {
	_, _ = tm.API.MarketBuy()
}

func (tm *TradeMaker) Act() error {
	walletsToQuery := []CoinbaseWalletName{ETHWallet, GBPWallet}
	wallets, err := tm.API.GetWallets(walletsToQuery)

	if err != nil {
		return err
	}

	if err := tm.assertWalletsFound(wallets); err != nil {
		return err
	}

	if FearBuyThreshold < tm.FearAndGreedScore && GreedSellThreshold > tm.FearAndGreedScore {
		tm.SummariseNoAction()
		return nil
	}

	if FearBuyThreshold > tm.FearAndGreedScore {
		tm.Buy()
		return nil
	}

	if GreedSellThreshold < tm.FearAndGreedScore {
		tm.Sell()
		return nil
	}

	return nil
}

func NewTradeMaker(options TradeMakerOptions) *TradeMaker {
	return &TradeMaker{FearAndGreedScore: options.FearAndGreedScore, API: options.API}
}
