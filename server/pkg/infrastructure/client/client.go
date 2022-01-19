package client

import (
	"entgo.io/ent/dialect"
	_ "github.com/lib/pq" // So we can use dialect.Postgres
	"stegoer/ent"
	"stegoer/pkg/infrastructure/env"
)

// New returns a new instance of ent.Client.
func New(config env.Config) (*ent.Client, error) {
	var entOptions []ent.Option

	if config.Debug {
		_ = append(entOptions, ent.Debug())
	}

	return ent.Open(dialect.Postgres, config.DatabaseURL, entOptions...)
}
