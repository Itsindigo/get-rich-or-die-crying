package app_config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type CoinbaseConfig struct {
	ApiKeyName string `env:"CB_API_KEY,required"`
	Secret     string `env:"CB_API_PRIVACY_KEY,required"`
}

type AppConfig struct {
	Coinbase            CoinbaseConfig
	ShouldMakeMinTrades bool `env:"SHOULD_MAKE_MIN_TRADES"`
	ForceSell           bool `env:"FORCE_SELL"`
	ForceBuy            bool `env:"FORCE_BUY"`
}

func LoadConfig() AppConfig {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := AppConfig{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Error mounting config: %v", err)
	}

	return cfg
}

func obfuscateSecret(s string, reveal_n int) string {
	if reveal_n > len(s) {
		return "***"
	}
	return fmt.Sprintf("%s***", s[0:reveal_n])
}

func (c AppConfig) String() string {
	return fmt.Sprintf(
		"AppConfig(Coinbase(ApiKeyName: %s, Secret: %s), ShouldMakeMinTrades: %t, ForceSell: %t, ForceBuy: %t)",
		obfuscateSecret(c.Coinbase.ApiKeyName, 3),
		obfuscateSecret(c.Coinbase.Secret, 5),
		c.ShouldMakeMinTrades,
		c.ForceSell,
		c.ForceBuy,
	)
}
