package main

import (
	"fmt"
	"log"

	"github.com/itsindigo/get-rich-or-die-crying/internal/app_config"
	"github.com/itsindigo/get-rich-or-die-crying/internal/coinbase"
	"github.com/itsindigo/get-rich-or-die-crying/internal/scraping"
	"github.com/itsindigo/get-rich-or-die-crying/internal/trading"
)

func main() {
	config := app_config.LoadConfig()

	fmt.Printf("Config: %v\n", config)

	score, err := scraping.ParseSentimentScore()

	if err != nil {
		log.Fatalf("Could not parse sentiment score: %v\n", err)
		return
	}

	coinbaseAPI := coinbase.NewCoinbaseAPI(coinbase.CoinbaseAPIConfig{KeyName: config.Coinbase.ApiKeyName, Secret: config.Coinbase.Secret})
	tm := trading.NewTradeMaker(trading.TradeMakerOptions{FearAndGreedScore: score, API: coinbaseAPI})
	tm.Act()
}
