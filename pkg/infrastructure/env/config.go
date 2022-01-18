package env

import (
	"github.com/spf13/viper"
	"log"
)

const configPath = "."

// Config represents the app configuration.
type Config struct {
	Debug       bool   `mapstructure:"DEBUG"`
	ServerPort  int    `mapstructure:"SERVER_PORT"`
	SecretKey   string `mapstructure:"SECRET_KEY"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
}

// LoadConfig loads and returns the env.Config struct.
func LoadConfig() Config {
	config, err := load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	return config
}

func load(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	config := Config{} //nolint:exhaustivestruct

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	return config, viper.Unmarshal(&config)
}
