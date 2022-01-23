package client

import (
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq" // So we can use dialect.Postgres

	"github.com/kucera-lukas/stegoer/ent"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/env"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/log"
)

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
