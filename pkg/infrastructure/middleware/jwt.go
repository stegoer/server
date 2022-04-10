package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/stegoer/server/ent"
	"github.com/stegoer/server/pkg/adapter/controller"
	"github.com/stegoer/server/pkg/infrastructure/log"
	"github.com/stegoer/server/pkg/util"
)

var userCtxKey = &contextKey{"user"} //nolint:gochecknoglobals

type contextKey struct {
	name string
}

// Jwt handles user authorization via JSON Web Tokens.
func Jwt(
	logger *log.Logger,
	controller controller.Controller,
) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(writer http.ResponseWriter,
				request *http.Request,
			) {
				header := request.Header.Get("Authorization")

				// allow unauthenticated users in
				if header == "" {
					logger.Debug("missing authorization header")

					next.ServeHTTP(writer, request)

					return
				}

				// validate jwt token
				userID, err := util.ParseToken(header)
				if err != nil {
					logger.Debugf("invalid token: %v", err)

					next.ServeHTTP(writer, request)

					return
				}

				ctx := request.Context()

				entUser, err := controller.User.GetByID(ctx, userID)
				if err != nil {
					logger.Debugf("user not found: %v", err)

					next.ServeHTTP(writer, request)

					return
				}

				logger.Debugf("user %s authorized", entUser.Name)

				// put user into context
				ctx = context.WithValue(ctx, userCtxKey, entUser)

				next.ServeHTTP(writer, request.WithContext(ctx))
			})
	}
}

// JwtForContext finds user from context. Requires Jwt to have run.
func JwtForContext(ctx context.Context) (*ent.User, error) {
	entUser, ok := ctx.Value(userCtxKey).(*ent.User)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return entUser, nil
}
