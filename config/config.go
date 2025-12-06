package config

import (
	"log"
	"os"
)

type Config struct {
	AppName     string
	Env         string
	HTTPPort    string
	DatabaseDSN string
}

func LoadConfig() *Config {
	cfg := &Config{
		AppName:  getEnv("APP_NAME", "CRUD App API"),
		Env:      getEnv("APP_ENV", "development"),
		HTTPPort: ":" + getEnv("HTTP_PORT", "8181"),
		DatabaseDSN: getEnv("DATABASE_DSN",
			""),
	}

	if cfg.DatabaseDSN == "" {
		log.Fatal("DATABASE_DSN is required")
	}

	return cfg
}

func getEnv(key, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}
