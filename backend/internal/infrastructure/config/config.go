package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	DatabaseURL    string
	ExternalAPIURL string
	AuthToken      string
	AllowedOrigin  string
}

func Load() (*Config, error) {
	godotenv.Load() // Load .env variables if exists

	return &Config{
		Port:           getEnvOrDefault("PORT", "8080"),
		DatabaseURL:    os.Getenv("DATABASE_URL"),
		ExternalAPIURL: os.Getenv("EXTERNAL_API_URL"),
		AuthToken:      os.Getenv("AUTH_TOKEN"),
		AllowedOrigin:  getEnvOrDefault("ALLOWED_ORIGIN", "*"),
	}, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
