package main

import (
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/env"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/log"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/server"
)

func main() {
	config := env.MustLoad()
	logger := log.MustNew(config)
	server.Run(config, logger)
}
