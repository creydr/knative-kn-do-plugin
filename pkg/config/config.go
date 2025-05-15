package config

import "os"

type Config struct {
	BaseUrl string
	ApiKey  string
	Model   string
}

const (
	EnvKeyApiBaseUrl = "KN_DO_OPENAI_API_ENDPOINT"
	EnvKeyApiKey     = "KN_DO_OPENAI_API_ENDPOINT_KEY"
	EnvKeyModelName  = "KN_DO_MODEL_NAME"
)

func NewFromEnv() Config {
	return Config{
		BaseUrl: getEnvOrDefault(EnvKeyApiBaseUrl, "http://localhost:11434/v1"),
		ApiKey:  getEnvOrDefault(EnvKeyApiKey, "foobar"),
		Model:   getEnvOrDefault(EnvKeyModelName, "qwen3:1.7b"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}

	return value
}
