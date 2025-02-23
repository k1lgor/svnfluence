package config

import (
	"fmt"
	"os"
)

type Config struct {
	OpenAIAPIKey string
}

func LoadConfig() (*Config, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable not set")
	}

	return &Config{
		OpenAIAPIKey: apiKey,
	}, nil
}
