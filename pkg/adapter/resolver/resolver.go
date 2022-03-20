package resolver

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your cmd,
// add any dependencies you require here.

import (
	"github.com/99designs/gqlgen/graphql"

	"github.com/kucera-lukas/stegoer/ent"
	"github.com/kucera-lukas/stegoer/graph/generated"
	"github.com/kucera-lukas/stegoer/pkg/adapter/controller"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/env"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/log"
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
	return generated.NewExecutableSchema(generated.Config{
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

func getDirective() generated.DirectiveRoot {
	return generated.DirectiveRoot{}
}

func getComplexity() generated.ComplexityRoot {
	return generated.ComplexityRoot{} //nolint:exhaustivestruct
}
