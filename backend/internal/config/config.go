package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type Config struct {
	OpenAIKey string
	Port      string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.Wrap(err, "error loading .env file")
	}

	return &Config{
		OpenAIKey: os.Getenv("OPENAI_API_KEY"),
		Port:      getEnv("PORT", ":8080"),
	}, nil
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
