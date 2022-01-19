package graphql

import (
	"context"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"

	"stegoer/ent"
	"stegoer/pkg/adapter/controller"
	"stegoer/pkg/adapter/resolver"
	"stegoer/pkg/entity/model"
)

const complexityLimit = 1000

// NewServer generates a new handler.Server.
func NewServer(client *ent.Client, controller controller.Controller) *handler.Server {
	srv := handler.NewDefaultServer(resolver.NewSchema(client, controller))
	srv.Use(entgql.Transactioner{TxOpener: client})
	srv.Use(extension.Introspection{})
	srv.Use(extension.FixedComplexityLimit(complexityLimit))
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		return model.NewInternalServerError(ctx, fmt.Sprintf(`%v`, err))
	})

	return srv
}
