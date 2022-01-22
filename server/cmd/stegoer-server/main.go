package main

import (
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/env"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/server"
)

func main() {
	config := env.Load()
	server.Run(config)
}
