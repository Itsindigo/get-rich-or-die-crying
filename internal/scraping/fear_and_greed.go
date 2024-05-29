package scraping

import (
	"log/slog"
	"strconv"

	"github.com/gocolly/colly/v2"
)

const CoinMarketCapHomePageURL = "https://coinmarketcap.com/"

func ParseSentimentScore() (int, error) {
	c := colly.NewCollector()

	var htmlScore string

	c.OnHTML("a[href='/charts/#fear-and-greed-index']", func(el *colly.HTMLElement) {
		children := el.DOM.Children().Nodes

		if len(children) == 0 {
			slog.Warn("Did not find fear and greed index URL element")
			return
		}

		// Take first child of first link element:
		linkElement := children[0]
		if linkElement.FirstChild.Data == "" {
			slog.Warn("Did not find score text in fear and greed URL")
			return
		}

		htmlScore = linkElement.FirstChild.Data
	})

	c.OnError(func(_ *colly.Response, err error) {
		if err != nil {
			slog.Error("Collector error", slog.String("error", err.Error()))
		}
	})

	err := c.Visit(CoinMarketCapHomePageURL)

	if err != nil {
		return 0, err
	}

	if htmlScore == "" {
		return 0, ScoreElementNotFoundError{}
	}

	score, err := strconv.Atoi(htmlScore)

	if err != nil {
		return 0, ScoreParseError{Value: htmlScore}
	}

	slog.Info("Parsed Fear & Greed Sentiment Score", slog.Int("score", score))

	return score, nil
}
