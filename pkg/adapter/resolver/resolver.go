package resolver

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"StegoLSB/ent"
	"StegoLSB/graph/generated"
	"StegoLSB/pkg/adapter/controller"
	"StegoLSB/pkg/infrastructure/middleware"
	"context"
	"github.com/99designs/gqlgen/graphql"
)

// Resolver is a context struct
type Resolver struct {
	client     *ent.Client
	controller controller.Controller
}

// NewSchema creates NewExecutableSchema
func NewSchema(client *ent.Client, controller controller.Controller) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers:  getResolver(client, controller),
		Directives: getDirective(),
		Complexity: getComplexity(),
	})
}

func getResolver(client *ent.Client, controller controller.Controller) *Resolver {
	return &Resolver{
		client:     client,
		controller: controller,
	}
}

func getDirective() generated.DirectiveRoot {
	return generated.DirectiveRoot{
		IsAuthenticated: func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
			if _, err := middleware.ForContext(ctx); err != nil {
				// block calling the next resolver
				return nil, err
			}

			// or let it pass through
			return next(ctx)
		},
	}
}

func getComplexity() generated.ComplexityRoot {
	return generated.ComplexityRoot{}
}
