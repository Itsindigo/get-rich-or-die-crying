package reporting

import (
	"context"
	"fmt"

	"github.com/itsindigo/get-rich-or-die-crying/internal/slack"
)

type Slack interface {
	SendMessage(ctx context.Context, message slack.Blocks) (string, error)
}

type TradeReporter struct {
	Slack Slack
}

func (tr *TradeReporter) ReportNoAction(ctx context.Context, score int) error {
	message, err := NoActionMessage(score)

	if err != nil {
		return fmt.Errorf("could not create no action message: %w", err)
	}

	_, err = tr.Slack.SendMessage(ctx, message)

	if err != nil {
		return err
	}

	return nil
}

func (tr *TradeReporter) ReportSale(ctx context.Context) {

}

func (tr *TradeReporter) ReportBuy(ctx context.Context) {

}

func (tr *TradeReporter) ReportError(ctx context.Context, err error) {

}

func NewTradeReporter(slack Slack) *TradeReporter {
	return &TradeReporter{
		Slack: slack,
	}
}
