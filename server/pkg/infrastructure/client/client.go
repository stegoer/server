package client

import (
	"fmt"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq" // So we can use dialect.Postgres

	"github.com/kucera-lukas/stegoer/ent"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/env"
)

// New returns a new instance of ent.Client.
func New(config *env.Config) (*ent.Client, error) {
	var entOptions []ent.Option

	if config.Debug {
		_ = append(entOptions, ent.Debug())
	}

	client, err := ent.Open(dialect.Postgres, config.DatabaseURL, entOptions...)
	if err != nil {
		return nil, fmt.Errorf("error opening database client: %w", err)
	}

	return client, nil
}
