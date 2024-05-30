package app_config

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type CoinbaseConfig struct {
	ApiKeyName string `env:"CB_API_KEY,required"`
	Secret     string `env:"CB_API_PRIVACY_KEY,required"`
}

type KnockConfig struct {
	Secret string `env:"KNOCK_SECRET_KEY,required"`
}

type AppConfig struct {
	EnableDebugLogs     bool `env:"ENABLE_DEBUG_LOGS"`
	Coinbase            CoinbaseConfig
	Knock               KnockConfig
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

	if cfg.EnableDebugLogs {
		slog.SetLogLoggerLevel(slog.LevelDebug)
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
		"AppConfig(EnableDebugLogs: %t, Coinbase(ApiKeyName: %s, Secret: %s), Knock(Secret: %s) ShouldMakeMinTrades: %t, ForceSell: %t, ForceBuy: %t)",
		c.EnableDebugLogs,
		obfuscateSecret(c.Coinbase.ApiKeyName, 3),
		obfuscateSecret(c.Coinbase.Secret, 5),
		obfuscateSecret(c.Knock.Secret, 12),
		c.ShouldMakeMinTrades,
		c.ForceSell,
		c.ForceBuy,
	)
}
