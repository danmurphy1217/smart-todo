package utils

import (
	"encoding/json"
	"os"
)

type GoogleCloudCfg struct {
	CredentialsSecretName string `json:"credentials_secret_name"`
}

type OpenAICfg struct {
	APIKeySecretName string `json:"api_key_secret_name"`
}

type APIConfig struct {
	Endpoint       string   `json:"endpoint"`
	AllowedOrigins []string `json:"allowed_origins"`
	URL            string   `json:"url"`
}

type Config struct {
	APIConfig      `json:"api"`
	GoogleCloudCfg `json:"google_cloud"`
	OpenAICfg      `json:"openai"`
}

func ParseConfig(cfgFile string) (*Config, error) {
	// Parse the config file
	var cfg *Config
	bytes, err := os.ReadFile(cfgFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
