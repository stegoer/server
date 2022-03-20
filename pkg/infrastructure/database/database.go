package database

import (
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq" // so we can use dialect.Postgres

	"github.com/stegoer/server/ent"
	"github.com/stegoer/server/pkg/infrastructure/env"
	"github.com/stegoer/server/pkg/infrastructure/log"
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

	drv, err := sql.Open(dialect.Postgres, config.DatabaseURL)
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
