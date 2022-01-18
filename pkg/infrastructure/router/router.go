package router

import (
	"StegoLSB/ent"
	"StegoLSB/pkg/infrastructure/env"
	"StegoLSB/pkg/infrastructure/middleware"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"net/http"
)

// Routes of Router
const (
	QueryPath      = "/graphql"
	PlaygroundPath = "/playground"
)

// New creates new mux router
func New(config env.Config, srv http.Handler, client *ent.Client) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.Jwt(client))

	router.Handle(QueryPath, srv)

	if config.Debug {
		router.HandleFunc(PlaygroundPath, playground.Handler("GQL Playground", QueryPath))
	}

	return router
}
