package app_config

import (
	"encoding/base64"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type CoinbaseSecret []byte

func (p *CoinbaseSecret) UnmarshalText(text []byte) error {
	out, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		if err != nil {
			log.Fatalf("Could not decode CoinbaseSecret: %s", err.Error())
		}
	}
	*p = out
	return nil
}

type CoinbaseConfig struct {
	ApiKeyName string         `env:"CB_API_KEY,required"`
	Secret     CoinbaseSecret `env:"CB_API_PRIVACY_KEY_B64,required"`
}

type SlackConfig struct {
	WebhookID string `env:"SLACK_WEBHOOK_ID,required"`
}

type AppConfig struct {
	EnableDebugLogs     bool `env:"ENABLE_DEBUG_LOGS"`
	Coinbase            CoinbaseConfig
	Slack               SlackConfig
	ShouldMakeMinTrades bool `env:"SHOULD_MAKE_MIN_TRADES"`
	ForceSell           bool `env:"FORCE_SELL"`
	ForceBuy            bool `env:"FORCE_BUY"`
}

func b64DecodeConfigVar(str string, fieldName string) string {
	decoded, err := base64.StdEncoding.DecodeString(str)

	if err != nil {
		log.Fatalf("Could not decode %q: %s", fieldName, err.Error())
	}

	return string(decoded)
}

func ConfigureApp() AppConfig {
	if os.Getenv("IS_REMOTE_ENVIRONMENT") == "" {
		err := godotenv.Load()

		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	cfg := AppConfig{}
	err := env.Parse(&cfg)

	if err != nil {
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
		"AppConfig(EnableDebugLogs: %t, Coinbase(ApiKeyName: %s, Secret: %s), Slack(WebhookID: %s) ShouldMakeMinTrades: %t, ForceSell: %t, ForceBuy: %t)",
		c.EnableDebugLogs,
		obfuscateSecret(c.Coinbase.ApiKeyName, 3),
		obfuscateSecret(string(c.Coinbase.Secret), 5),
		obfuscateSecret(c.Slack.WebhookID, 12),
		c.ShouldMakeMinTrades,
		c.ForceSell,
		c.ForceBuy,
	)
}
