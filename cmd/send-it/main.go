package main

import (
	"log"

	"github.com/itsindigo/get-rich-or-die-crying/internal/app_config"
	"github.com/itsindigo/get-rich-or-die-crying/internal/scraping"
	"github.com/itsindigo/get-rich-or-die-crying/internal/trading"
)

func main() {
	config := app_config.LoadConfig()
	score, err := scraping.ParseSentimentScore()

	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	coinbaseAPI := trading.NewCoinbaseAPI(trading.CoinbaseAPIConfig{KeyName: config.Coinbase.ApiKeyName, Secret: config.Coinbase.Secret})
	tm := trading.NewTradeMaker(trading.TradeMakerOptions{FearAndGreedScore: score, API: coinbaseAPI})

	err = tm.Act(trading.ActOptions{ForceSell: false, ForceBuy: false})

	if err != nil {
		log.Fatalf(err.Error())
	}
}
