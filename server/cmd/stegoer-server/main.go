package main

import (
	"stegoer/pkg/infrastructure/env"
	"stegoer/pkg/infrastructure/server"
)

func main() {
	config := env.LoadConfig()
	server.Run(config)
}
