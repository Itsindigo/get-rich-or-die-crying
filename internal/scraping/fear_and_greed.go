package scraping

import (
	"fmt"
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
			fmt.Println("Did not find fear and greed index URL element")
			return
		}

		// Take first child of first link element:
		linkElement := children[0]

		if linkElement.FirstChild.Data == "" {
			fmt.Println("Did not find score text in fear and greed URL")
			return
		}

		htmlScore = linkElement.FirstChild.Data
	})

	c.OnError(func(_ *colly.Response, err error) {
		if err != nil {
			fmt.Printf("Collector error %v", err)
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

	return score, nil
}
