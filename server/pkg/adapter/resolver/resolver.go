package resolver

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your cmd,
// add any dependencies you require here.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"

	"github.com/kucera-lukas/stegoer/ent"
	"github.com/kucera-lukas/stegoer/graph/generated"
	"github.com/kucera-lukas/stegoer/pkg/adapter/controller"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/env"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/log"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/middleware"
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
		Directives: getDirective(logger),
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

func getDirective(logger *log.Logger) generated.DirectiveRoot {
	return generated.DirectiveRoot{
		IsAuthenticated: func(
			ctx context.Context,
			obj interface{},
			next graphql.Resolver,
		) (res interface{}, err error) {
			entUser, err := middleware.JwtForContext(ctx)
			if err != nil {
				logger.Debugf("@isAuthenticated invalid request: %v", err)

				return nil, err //nolint:wrapcheck
			}

			logger.Debugf("@isAuthenticated valid user: %s", entUser.Name)

			return next(ctx)
		},
	}
}

func getComplexity() generated.ComplexityRoot {
	return generated.ComplexityRoot{} //nolint:exhaustivestruct
}
