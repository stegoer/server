package log

import (
	"log"

	"go.uber.org/zap"

	"github.com/kucera-lukas/stegoer/pkg/infrastructure/env"
)

// Logger wraps the zap.SugaredLogger and zap.Logger types.
type Logger = zap.SugaredLogger

func New(config *env.Config) *Logger {
	cfg := updateConfig(getConfig(config))

	logger, err := cfg.Build()
	if err != nil {
		log.Panicf("can't initialize logger: %v", err)
	}

	return logger.Sugar()
}

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
