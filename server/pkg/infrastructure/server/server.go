package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kucera-lukas/stegoer/ent"
	"github.com/kucera-lukas/stegoer/pkg/adapter/controller"
	"github.com/kucera-lukas/stegoer/pkg/adapter/repository"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/database"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/env"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/graphql"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/log"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/redis"
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
	entClient := database.MustNew(config, logger)
	redisClient := redis.MustNew(config, logger)
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

func newController(client *ent.Client) controller.Controller {
	return controller.Controller{
		User:  repository.NewUserRepository(client),
		Image: repository.NewImageRepository(client),
	}
}
