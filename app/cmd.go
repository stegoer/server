package main

import (
	"StegoLSB/pkg/infrastructure/env"
	"StegoLSB/pkg/infrastructure/server"
)

func main() {
	config := env.LoadConfig()
	server.Run(config)
}
