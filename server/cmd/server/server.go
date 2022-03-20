package main

import (
	"github.com/kucera-lukas/stegoer/server/pkg/infrastructure/env"
	"github.com/kucera-lukas/stegoer/server/pkg/infrastructure/log"
	"github.com/kucera-lukas/stegoer/server/pkg/infrastructure/server"
)

func main() {
	config := env.MustLoad()
	logger := log.MustNew(config)
	server.Run(config, logger)
}
