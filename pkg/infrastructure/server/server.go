package server

import (
	"StegoLSB/ent"
	"StegoLSB/ent/migrate"
	"StegoLSB/pkg/adapter/controller"
	"StegoLSB/pkg/infrastructure/client"
	"StegoLSB/pkg/infrastructure/env"
	"StegoLSB/pkg/infrastructure/graphql"
	"StegoLSB/pkg/infrastructure/router"
	"StegoLSB/pkg/registry"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const timeOutDeadline = time.Second * 30

// Run runs the server with the given env.Config configuration.
func Run(config env.Config) {
	run(create(config))
}

func create(config env.Config) *http.Server {
	entClient := newDBClient(config)
	ctrl := newController(entClient)

	gqlSrv := graphql.NewServer(entClient, ctrl)
	muxRouter := router.New(config, gqlSrv, entClient)

	return &http.Server{ //nolint:exhaustivestruct
		Addr:         fmt.Sprintf(`:%d`, config.ServerPort),
		WriteTimeout: timeOutDeadline,
		ReadTimeout:  timeOutDeadline,
		IdleTimeout:  timeOutDeadline,
		Handler:      muxRouter,
	}
}

func run(srv *http.Server) {
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println("listening on", srv.Addr)

		if err := srv.ListenAndServe(); err != nil {
			log.Println(fmt.Sprintf("http server terminated: %v", err))
		}
	}()

	channel := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(channel, os.Interrupt)

	// Block until we receive our signal.
	<-channel

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), timeOutDeadline)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("error shutting down the server: %v", err)
	}
	log.Println("server shutdown")
	os.Exit(0)
}

func newDBClient(config env.Config) *ent.Client {
	entClient, err := client.New(config)
	if err != nil {
		log.Fatalf("failed to open postgres client: %v", err)
	}

	if err := entClient.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(true),
	); err != nil {
		log.Fatalf("failed to create schema resources: %v", err)
	}

	return entClient
}

func newController(client *ent.Client) controller.Controller {
	return registry.New(client).NewController()
}
