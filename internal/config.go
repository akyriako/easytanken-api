package internal

import (
	"errors"
	"os"
)

const (
	DefaultPortEnvName  string = "HTTP_PORT"
	DefaultPortEnvValue string = "8080"
	ApiKeyEnvName       string = "TK_API_KEY"
)

type Config struct {
	Port   string
	ApiKey string
}

func GetConfig() (*Config, error) {
	port, exists := os.LookupEnv(DefaultPortEnvName)
	if !exists {
		port = DefaultPortEnvValue
	}

	apiKey, exists := os.LookupEnv(ApiKeyEnvName)
	if !exists {
		apiKey = DefaultPortEnvValue

		return nil, errors.New("no TK_API_KEY found in environment variables")
	}

	config := &Config{
		Port:   port,
		ApiKey: apiKey,
	}

	return config, nil
}
