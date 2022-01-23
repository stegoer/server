package env

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

const configPath = "."

// Config represents the env configuration.
type Config struct {
	Debug       bool   `mapstructure:"DEBUG"`
	ServerPort  int    `mapstructure:"SERVER_PORT"`
	SecretKey   string `mapstructure:"SECRET_KEY"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
}

// Load loads and returns the env.Config struct.
func Load() *Config {
	config, err := load(configPath)
	if err != nil {
		log.Panicf("failed to load config: %v", err)
	}

	return config
}

func load(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf(`error reading configuration: %w`, err)
	}

	config := Config{} //nolint:exhaustivestruct

	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf(`error unmarshaling config: %w`, err)
	}

	return &config, nil
}
