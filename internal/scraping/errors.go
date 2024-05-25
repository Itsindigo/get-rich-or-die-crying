package scraping

import "fmt"

type ScoreElementNotFoundError struct{}

func (e ScoreElementNotFoundError) Error() string {
	return fmt.Sprintf("Did not find a score value while parsing %s", CoinMarketCapHomePageURL)
}

type ScoreParseError struct {
	Value string
}

func (e ScoreParseError) Error() string {
	return fmt.Sprintf("Could not parse score value %q to int", e.Value)
}
