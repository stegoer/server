//go:server go run ./app.go

package main

import (
	"StegoLSB/ent"
	"StegoLSB/ent/migrate"
	"StegoLSB/pkg/adapter/controller"
	"StegoLSB/pkg/infrastructure/client"
	"StegoLSB/pkg/infrastructure/graphql"
	"StegoLSB/pkg/infrastructure/router"
	"StegoLSB/pkg/registry"
	"context"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	defaultPort     = "8080"
	timeOutDeadline = time.Second * 15
)

func main() {
	entClient := newDBClient()

	ctrl := newController(entClient)

	gqlSrv := graphql.NewServer(entClient, ctrl)
	muxRouter := router.New(gqlSrv, entClient)

	port, ok := os.LookupEnv("PORT")
	if !ok || port == "" {
		port = defaultPort
	}

	srv := &http.Server{ //nolint:exhaustivestruct
		Addr:         fmt.Sprintf(`:%s`, port),
		WriteTimeout: timeOutDeadline,
		ReadTimeout:  timeOutDeadline,
		IdleTimeout:  timeOutDeadline,
		Handler:      muxRouter,
	}

	runServer(srv)
}

func runServer(srv *http.Server) {
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

func newDBClient() *ent.Client {
	entClient, err := client.New()
	if err != nil {
		log.Fatalf("failed to open postgres client: %v", err)
	}

	if err := entClient.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return entClient
}

func newController(client *ent.Client) controller.Controller {
	return registry.New(client).NewController()
}
