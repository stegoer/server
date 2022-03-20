package main

import (
	"context"

	"ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect/sql/schema"

	"github.com/kucera-lukas/stegoer/server/pkg/infrastructure/database"
	"github.com/kucera-lukas/stegoer/server/pkg/infrastructure/env"
	"github.com/kucera-lukas/stegoer/server/pkg/infrastructure/log"
)

func main() {
	config := env.MustLoad()
	logger := log.MustNew(config)
	client := database.MustNew(config, logger)

	defer func() {
		if err := client.Close(); err != nil {
			logger.Panic(err)
		}
	}()

	// Create a local migration directory.
	dir, err := migrate.NewLocalDir("migrations")
	if err != nil {
		logger.Warnf("failed creating atlas migration directory: %v", err)

		return
	}

	// Write migration diff.
	err = client.Schema.Diff(context.Background(), schema.WithDir(dir))
	if err != nil {
		logger.Warnf("failed writing migration diff: %v", err)

		return
	}
}
