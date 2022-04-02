package graphql

import (
	"context"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/go-redis/redis/v8"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/stegoer/server/ent"
	"github.com/stegoer/server/pkg/adapter/controller"
	"github.com/stegoer/server/pkg/adapter/resolver"
	"github.com/stegoer/server/pkg/infrastructure/apq"
	"github.com/stegoer/server/pkg/infrastructure/env"
	"github.com/stegoer/server/pkg/infrastructure/log"
	"github.com/stegoer/server/pkg/model"
)

const complexityLimit = 1000

// NewServer generates a new handler.Server.
func NewServer(
	config *env.Config,
	logger *log.Logger,
	client *ent.Client,
	redisClient *redis.Client,
	controller controller.Controller,
) *handler.Server {
	srv := handler.NewDefaultServer(
		resolver.NewSchema(config, logger, client, controller),
	)
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.Use(entgql.Transactioner{TxOpener: client})
	srv.Use(extension.Introspection{})
	srv.Use(extension.FixedComplexityLimit(complexityLimit))
	srv.Use(extension.AutomaticPersistedQuery{Cache: apq.NewCache(*redisClient)})
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		logger.Debugw("graphql request failed",
			"message", err.Message,
			"path", err.Path,
			"locations", err.Locations,
			"extensions", err.Extensions,
			"rule", err.Rule,
		)

		return err
	})
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		return model.NewInternalServerError(ctx, fmt.Sprintf(`%v`, err))
	})

	return srv
}
