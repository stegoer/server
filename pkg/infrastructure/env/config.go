package env

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload" // load .env variables
	"github.com/spf13/viper"
)

type environment string

// Config represents the environment configuration.
type Config struct {
	Env           environment `mapstructure:"ENV"`
	Debug         bool        `mapstructure:"DEBUG"`
	Port          int         `mapstructure:"PORT"`
	SecretKey     string      `mapstructure:"SECRET_KEY"`
	EncryptionKey string      `mapstructure:"ENCRYPTION_KEY"`
	DatabaseURL   string      `mapstructure:"DATABASE_URL"`
	RedisURL      string      `mapstructure:"REDIS_URL"`
}

const (
	configPath = "."

	development environment = "DEVELOPMENT"
	production  environment = "PRODUCTION"
)

// IsDevelopment returns whether the Config represents a development environment.
func (c *Config) IsDevelopment() bool {
	return c.Env == development
}

// IsProduction returns whether the Config represents a production environment.
func (c *Config) IsProduction() bool {
	return c.Env == production
}

// MustLoad ensures that a new env.Config struct is loaded and panics if not.
func MustLoad() *Config {
	config, err := Load()
	if err != nil {
		log.Panic(err)
	}

	return config
}

// Load loads and returns the env.Config struct.
func Load() (*Config, error) {
	config, err := load(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return config, nil
}

func load(path string) (*Config, error) {
	if os.Getenv("ENV") != "PRODUCTION" {
		if err := setConfig(path); err != nil {
			return nil, err
		}
	}

	setDefault()

	config := Config{} //nolint:exhaustivestruct

	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf(`error unmarshaling config: %w`, err)
	}

	return &config, nil
}

func setConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf(`error reading configuration: %w`, err)
	}

	return nil
}

func setDefault() {
	viper.SetDefault("ENV", development)
	viper.SetDefault("DEBUG", false)

	for _, key := range getEnvDefaultKeys() {
		setEnvDefault(key)
	}
}

func setEnvDefault(key string) {
	viper.SetDefault(key, os.Getenv(key))
}

func getEnvDefaultKeys() []string {
	return []string{
		"PORT",
		"SECRET_KEY",
		"ENCRYPTION_KEY",
		"DATABASE_URL",
		"REDIS_URL",
	}
}
