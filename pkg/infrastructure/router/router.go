package router

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/stegoer/server/ent"
	"github.com/stegoer/server/pkg/infrastructure/env"
	"github.com/stegoer/server/pkg/infrastructure/log"
	"github.com/stegoer/server/pkg/infrastructure/middleware"
)

// Routes of our mux.Router.
const (
	queryPath      = "/graphql"
	playgroundPath = "/playground"
)

// New creates new mux router.
func New(
	config *env.Config,
	logger *log.Logger,
	srv http.Handler,
	client *ent.Client,
) http.Handler {
	router := mux.NewRouter()
	router.Use(middleware.Logging, middleware.Jwt(logger, client))

	router.Handle(queryPath, srv)

	var crossOrigin *cors.Cors

	switch config.IsDevelopment() {
	case true:
		router.HandleFunc(
			playgroundPath,
			playground.Handler("GQL Playground", queryPath),
		)

		crossOrigin = cors.AllowAll()
	case false:
		crossOrigin = cors.New(cors.Options{ //nolint:exhaustivestruct
			AllowedOrigins: []string{"*"},
			AllowedHeaders: []string{"Authorization"},
			Debug:          config.Debug,
		})
	}

	return crossOrigin.Handler(router)
}
