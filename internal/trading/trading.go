package trading

import "fmt"

type TradeMaker struct {
	FearAndGreedScore int
}

func (tm *TradeMaker) SummariseNoAction() {
	fmt.Printf("Today's F&G score was %d, no trade was made.\n", tm.FearAndGreedScore)
}

func (tm *TradeMaker) Sell() {

}

func (tm *TradeMaker) Buy() {

}

func (tm *TradeMaker) Act() {
	if FearBuyThreshold < tm.FearAndGreedScore && GreedSellThreshold > tm.FearAndGreedScore {
		tm.SummariseNoAction()
		return
	}

	if FearBuyThreshold > tm.FearAndGreedScore {
		tm.Buy()
		return
	}

	if GreedSellThreshold < tm.FearAndGreedScore {
		tm.Sell()
	}
}

func NewTradeMaker(fearAndGreedScore int) *TradeMaker {
	return &TradeMaker{FearAndGreedScore: fearAndGreedScore}
}
