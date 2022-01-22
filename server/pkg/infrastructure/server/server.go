package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kucera-lukas/stegoer/ent"
	"github.com/kucera-lukas/stegoer/ent/migrate"
	"github.com/kucera-lukas/stegoer/pkg/adapter/controller"
	"github.com/kucera-lukas/stegoer/pkg/adapter/repository"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/client"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/env"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/graphql"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/router"
)

const (
	timeOutDeadline = time.Second * 30
	shutdownSignal  = 1
)

// Run runs the server with the given env.Config configuration.
func Run(config *env.Config) {
	run(create(config))
}

func create(config *env.Config) *http.Server {
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

	channel := make(chan os.Signal, shutdownSignal)
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
	if err := srv.Shutdown(ctx); err != nil {
		log.Panicf("error shutting down the server: %v", err)
	}

	log.Println("server shutdown")
}

func newDBClient(config *env.Config) *ent.Client {
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
	return controller.Controller{
		User:  repository.NewUserRepository(client),
		Image: repository.NewImageRepository(client),
	}
}
