package trading

import (
	"fmt"
)

type TraderAPI interface {
	GetWallets([]string) ([]SimpleAccount, error)
	CreateOrder() (interface{}, error)
}

type TradeMaker struct {
	FearAndGreedScore int
	API               TraderAPI
}

type TradeMakerOptions struct {
	FearAndGreedScore int
	API               TraderAPI
}

func NewTradeMaker(options TradeMakerOptions) *TradeMaker {
	return &TradeMaker{FearAndGreedScore: options.FearAndGreedScore, API: options.API}
}

func (tm *TradeMaker) SummariseNoAction() {
	fmt.Printf("Today's F&G score was %d, no trade was made.\n", tm.FearAndGreedScore)
}

func (tm *TradeMaker) Sell() {

}

func (tm *TradeMaker) Buy() {

}

func (tm *TradeMaker) Act() error {
	wallets := []string{"ETH Wallet", "BTC Wallet", "GBP Wallet"}
	accounts, err := tm.API.GetWallets(wallets)

	if err != nil {
		return fmt.Errorf("err: %w", err)
	}

	fmt.Printf("The Accs: %v", accounts)

	_, err = tm.API.CreateOrder()

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
