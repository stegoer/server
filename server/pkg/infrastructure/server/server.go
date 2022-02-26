package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/kucera-lukas/stegoer/ent"
	"github.com/kucera-lukas/stegoer/ent/migrate"
	"github.com/kucera-lukas/stegoer/pkg/adapter/controller"
	"github.com/kucera-lukas/stegoer/pkg/adapter/repository"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/client"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/env"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/graphql"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/log"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/router"
)

const (
	timeOutDeadline = time.Second * 30
	shutdownSignal  = 1
)

// Run runs the server with the given env.Config configuration.
func Run(config *env.Config, logger *log.Logger) {
	srv := create(config, logger)
	run(logger, srv)
}

func create(config *env.Config, logger *log.Logger) *http.Server {
	entClient := newDBClient(config, logger)
	redisClient := newRedisClient(config, logger)
	ctrl := newController(entClient)

	gqlSrv := graphql.NewServer(config, logger, entClient, redisClient, ctrl)
	muxRouter := router.New(config, logger, gqlSrv, entClient)

	return &http.Server{ //nolint:exhaustivestruct
		Addr:         fmt.Sprintf(`0.0.0.0:%d`, config.Port),
		WriteTimeout: timeOutDeadline,
		ReadTimeout:  timeOutDeadline,
		IdleTimeout:  timeOutDeadline,
		Handler:      muxRouter,
	}
}

func run(logger *log.Logger, srv *http.Server) {
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		logger.Infof("listening on %s", srv.Addr)

		if err := srv.ListenAndServe(); err != nil {
			logger.Infof("http server terminated: %v", err)
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
		logger.Panicf("error shutting down the server: %v", err)
	}

	logger.Info("server shutdown")
}

func newDBClient(config *env.Config, logger *log.Logger) *ent.Client {
	entClient, err := client.New(config, logger)
	if err != nil {
		logger.Panicf("failed to open postgres client: %v", err)
	}

	if err := entClient.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(true),
	); err != nil {
		logger.Panicf("failed to create schema resources: %v", err)
	}

	return entClient
}

func newRedisClient(config *env.Config, logger *log.Logger) *redis.Client {
	redisOptions, err := redis.ParseURL(config.RedisURL)
	if err != nil {
		logger.Panicf("failed to parse %s as a redis url: %v", config.RedisURL, err)
	}

	redisClient := redis.NewClient(redisOptions)

	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		logger.Panicf("failed to open redis client: %v", err)
	}

	return redisClient
}

func newController(client *ent.Client) controller.Controller {
	return controller.Controller{
		User:  repository.NewUserRepository(client),
		Image: repository.NewImageRepository(client),
	}
}
