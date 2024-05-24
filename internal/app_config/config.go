package app_config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBHost     string `env:"DB_HOST,required"`
	DBPort     int    `env:"DB_PORT,required"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBName     string `env:"DB_NAME,required"`
}

type AppConfig struct {
	DB             DBConfig
	CoinbaseAPIKey string `env:"CB_ACCESS_KEY"`
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

func (c AppConfig) String() string {
	return fmt.Sprintf(
		"AppConfig(DB(host: %s, port: %d, user: %s, password: ***, name: %s), CoinbaseAPIKey: ***)",
		c.DB.DBHost,
		c.DB.DBPort,
		c.DB.DBUser,
		c.DB.DBName,
	)
}
