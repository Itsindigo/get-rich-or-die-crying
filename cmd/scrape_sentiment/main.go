package main

import (
	"fmt"
	"log"

	"github.com/itsindigo/get-rich-or-die-crying/internal/app_config"
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

	tm := trading.NewTradeMaker(score)
	tm.Act()
}
