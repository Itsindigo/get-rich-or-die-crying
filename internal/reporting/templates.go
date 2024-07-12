package reporting

import (
	"fmt"

	"github.com/itsindigo/get-rich-or-die-crying/internal/slack"
)

func GetNoActionMessageBlocks(score int) (slack.Blocks, error) {
	message := fmt.Sprintf("*No Trade Today:* Fear & Greed index score was *%d*", score)
	blocks := []slack.SectionBlock{
		{Type: "section", Text: slack.TextBlock{Type: "mrkdwn", Text: message}},
	}

	return slack.NewBlocksMap(blocks)
}

func GetInsufficientFundsMessageBlocks(score int) (slack.Blocks, error) {
	message := fmt.Sprintf("*No Trade Today:* Fear score is *%d*, but funds are empty, no purchase was made.", score)
	blocks := []slack.SectionBlock{
		{Type: "section", Text: slack.TextBlock{Type: "mrkdwn", Text: message}},
	}

	return slack.NewBlocksMap(blocks)
}

func GetSaleMadeMessageBlocks(amount string) (slack.Blocks, error) {
	message := fmt.Sprintf("*Sale completed:* sold *%s ETH*", amount)
	blocks := []slack.SectionBlock{
		{Type: "section", Text: slack.TextBlock{Type: "mrkdwn", Text: message}},
	}

	return slack.NewBlocksMap(blocks)
}

func GetPurchaseMadeMessageBlocks(amount string) (slack.Blocks, error) {
	message := fmt.Sprintf("*Purchase completed:* spent *£%s*", amount)
	blocks := []slack.SectionBlock{
		{Type: "section", Text: slack.TextBlock{Type: "mrkdwn", Text: message}},
	}

	return slack.NewBlocksMap(blocks)
}

func GetErrorMessageBlocks(error string) (slack.Blocks, error) {
	message := fmt.Sprintf("*Error occurred while making trade:* %s", error)
	blocks := []slack.SectionBlock{
		{Type: "section", Text: slack.TextBlock{Type: "mrkdwn", Text: message}},
	}

	return slack.NewBlocksMap(blocks)
}
