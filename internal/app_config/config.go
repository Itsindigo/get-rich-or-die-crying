package app_config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required"`
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	Name     string `env:"DB_NAME,required"`
}

type CoinbaseConfig struct {
	ApiKeyName string `env:"CB_API_KEY,required"`
	Secret     string `env:"CB_API_PRIVAY_KEY,required"`
}

type AppConfig struct {
	DB                  DBConfig
	Coinbase            CoinbaseConfig
	ShouldMakeMinTrades bool `env:"SHOULD_MAKE_MIN_TRADES"`
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
		"AppConfig(DB(Host: %s, Port: %d, User: %s, Password: %s Name: %s), Coinbase(ApiKeyName: %s, Secret: %s), ShouldMakeMinTrades: %t)",
		c.DB.Host,
		c.DB.Port,
		c.DB.User,
		obfuscateSecret(c.DB.Password, 3),
		c.DB.Name,
		obfuscateSecret(c.Coinbase.ApiKeyName, 3),
		obfuscateSecret(c.Coinbase.Secret, 5),
		c.ShouldMakeMinTrades,
	)
}
