package graphql

import (
	"StegoLSB/ent"
	"StegoLSB/pkg/adapter/controller"
	"StegoLSB/pkg/adapter/resolver"
	"StegoLSB/pkg/entity/model"
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler/extension"

	"entgo.io/contrib/entgql"

	"github.com/99designs/gqlgen/graphql/handler"
)

const complexityLimit = 1000

// NewServer generates graphql server
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
