package main

import (
	"log"
	"log/slog"

	"github.com/itsindigo/get-rich-or-die-crying/internal/app_config"
	"github.com/itsindigo/get-rich-or-die-crying/internal/reporting"
	"github.com/itsindigo/get-rich-or-die-crying/internal/scraping"
	"github.com/itsindigo/get-rich-or-die-crying/internal/trading"
)

func main() {
	config := app_config.LoadConfig()
	slog.Debug("Config", slog.String("config", config.String()))

	score, err := scraping.ParseSentimentScore()

	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	tradeReporter := reporting.NewTradeReporter()
	tm := trading.NewTradeMaker(
		trading.TradeMakerOptions{
			FearAndGreedScore: score,
			MakeMinTrades:     config.ShouldMakeMinTrades,
			API: trading.NewCoinbaseAPI(
				trading.CoinbaseAPIConfig{
					KeyName: config.Coinbase.ApiKeyName,
					Secret:  config.Coinbase.Secret,
				},
			),
			TradeReporter: tradeReporter,
		},
	)

	err = tm.Act(trading.ActOptions{ForceSell: config.ForceSell, ForceBuy: config.ForceBuy})

	if err != nil {
		tradeReporter.ReportError(err)
		log.Fatalf(err.Error())
	}
}
