package router

import (
	"StegoLSB/pkg/infrastructure/middleware"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
)

// Routes of Router
const (
	QueryPath      = "/graphql"
	PlaygroundPath = "/playground"
)

// New creates new mux router
func New(srv *handler.Server) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.JwtMiddleware())

	router.Handle(QueryPath, srv)
	router.HandleFunc(PlaygroundPath, playground.Handler("GQL Playground", QueryPath))

	return router
}
