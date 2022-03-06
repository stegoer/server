package database

import (
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/lib/pq"

	"github.com/kucera-lukas/stegoer/ent"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/env"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/log"
)

// MustNew ensure that a new ent.Client is created and panics if not.
func MustNew(config *env.Config, logger *log.Logger) *ent.Client {
	client, err := New(config, logger)
	if err != nil {
		logger.Panic(err)
	}

	return client
}

// New returns a new instance of ent.Client.
func New(config *env.Config, logger *log.Logger) (*ent.Client, error) {
	var entOptions []ent.Option

	var drv dialect.Driver

	url, err := pq.ParseURL(config.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database url: %w", err)
	}

	drv, err = sql.Open(dialect.Postgres, url)
	if err != nil {
		return nil, fmt.Errorf("error opening database client: %w", err)
	}

	if config.Debug {
		drv = dialect.Debug(drv, debugOperation(logger))
		logger.Info("debug driver set up")
	}

	return ent.NewClient(append(entOptions, ent.Driver(drv))...), nil
}

func debugOperation(logger *log.Logger) func(...interface{}) {
	return func(op ...interface{}) {
		logger.Debugf("ent operation: %s", op)
	}
}
