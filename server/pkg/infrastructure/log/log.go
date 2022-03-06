package log

import (
	"fmt"
	"log"

	"go.uber.org/zap"

	"github.com/kucera-lukas/stegoer/pkg/infrastructure/env"
)

// Logger wraps the zap.SugaredLogger and zap.Logger types.
type Logger = zap.SugaredLogger

// MustNew ensure that a new Logger is created and panics if not.
func MustNew(config *env.Config) *Logger {
	logger, err := New(config)
	if err != nil {
		log.Panic(err)
	}

	return logger
}

// New returns a new instance of Logger.
func New(config *env.Config) (*Logger, error) {
	cfg := updateConfig(getConfig(config))

	logger, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("can't initialize logger: %w", err)
	}

	return logger.Sugar(), nil
}

// Sync syncs a Logger.
func Sync(logger Logger) {
	if err := logger.Sync(); err != nil {
		log.Panicf("failed to sync logger: %v", err)
	}
}

func getConfig(config *env.Config) (cfg zap.Config) {
	switch config.IsDevelopment() {
	case true:
		cfg = zap.NewDevelopmentConfig()
	case false:
		cfg = zap.NewProductionConfig()
	}

	return
}

func updateConfig(cfg zap.Config) zap.Config {
	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	cfg.Encoding = "console"

	return cfg
}
