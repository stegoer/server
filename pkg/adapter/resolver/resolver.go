package resolver

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your cmd,
// add any dependencies you require here.

import (
	"github.com/99designs/gqlgen/graphql"

	"github.com/stegoer/server/ent"
	"github.com/stegoer/server/gqlgen"
	"github.com/stegoer/server/pkg/adapter/controller"
	"github.com/stegoer/server/pkg/infrastructure/env"
	"github.com/stegoer/server/pkg/infrastructure/log"
)

// Resolver is a context struct.
type Resolver struct {
	config     *env.Config
	logger     *log.Logger
	client     *ent.Client
	controller controller.Controller
}

// NewSchema creates a new graphql.ExecutableSchema.
func NewSchema( //nolint:ireturn
	config *env.Config,
	logger *log.Logger,
	client *ent.Client,
	controller controller.Controller,
) graphql.ExecutableSchema {
	return gqlgen.NewExecutableSchema(gqlgen.Config{
		Resolvers:  getResolver(config, logger, client, controller),
		Directives: getDirective(),
		Complexity: getComplexity(),
	})
}

func getResolver(
	config *env.Config,
	logger *log.Logger,
	client *ent.Client,
	controller controller.Controller,
) *Resolver {
	return &Resolver{
		config:     config,
		logger:     logger,
		client:     client,
		controller: controller,
	}
}

func getDirective() gqlgen.DirectiveRoot {
	return gqlgen.DirectiveRoot{}
}

func getComplexity() gqlgen.ComplexityRoot {
	return gqlgen.ComplexityRoot{} //nolint:exhaustivestruct
}
