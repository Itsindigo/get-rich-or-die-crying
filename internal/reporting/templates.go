package reporting

import (
	"fmt"

	"github.com/itsindigo/get-rich-or-die-crying/internal/slack"
)

func NoActionMessage(score int) (slack.Blocks, error) {
	message := fmt.Sprintf("*No Trade Today:* Score was *%d*", score)
	blocks := []slack.SectionBlock{
		{Type: "section", Text: slack.TextBlock{Type: "mrkdwn", Text: message}},
	}

	return slack.NewBlocksMap(blocks)
}
