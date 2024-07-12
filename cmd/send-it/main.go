package main

import (
	"context"
	"errors"
	"log"
	"log/slog"

	"github.com/itsindigo/get-rich-or-die-crying/internal/app_config"
	"github.com/itsindigo/get-rich-or-die-crying/internal/reporting"
	"github.com/itsindigo/get-rich-or-die-crying/internal/scraping"
	"github.com/itsindigo/get-rich-or-die-crying/internal/slack"
	"github.com/itsindigo/get-rich-or-die-crying/internal/trading"
)

func main() {
	ctx := context.Background()
	config := app_config.ConfigureApp()
	slog.Debug("Config", slog.String("config", config.String()))

	slackClient := slack.NewSlack(config.Slack.WebhookID)

	tradeReporter := reporting.NewTradeReporter(slackClient)
	tm := trading.NewTradeMaker(
		trading.TradeMakerOptions{
			MakeMinTrades: config.ShouldMakeMinTrades,
			API: trading.NewCoinbaseAPI(
				trading.CoinbaseAPIConfig{
					KeyName: config.Coinbase.ApiKeyName,
					Secret:  string(config.Coinbase.Secret),
				},
			),
			TradeReporter: tradeReporter,
		},
	)

	score, err := scraping.ParseSentimentScore()

	if err != nil {
		tradeReporter.ReportError(ctx, err)

		slog.Error(err.Error())
		log.Fatalf("Could not parse sentiment score, exiting.")
		return
	}

	err = tm.Act(ctx, trading.ActOptions{FearAndGreedScore: score, ForceSell: config.ForceSell, ForceBuy: config.ForceBuy})

	if err != nil {
		if errors.Is(err, trading.ErrInsufficientEthGbp) {
			tradeReporter.ReportInsufficientFunds(ctx, score)
			return
		}

		tradeReporter.ReportError(ctx, err)

		log.Fatalf(err.Error())
		return
	}

	slog.Info("Job finished, exiting.")
}
