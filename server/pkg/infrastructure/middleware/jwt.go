package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kucera-lukas/stegoer/ent"
	"github.com/kucera-lukas/stegoer/pkg/adapter/repository"
	"github.com/kucera-lukas/stegoer/pkg/entity/model"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/log"
	"github.com/kucera-lukas/stegoer/pkg/util"
)

const (
	errorFormatString = `{
"errors":[{"message": "%s", 
"path": ["Authorization"], 
"extensions": {"code": "AUTHORIZATION_ERROR"}}],
data: null
}`
)

var userCtxKey = &contextKey{"user"} //nolint:gochecknoglobals

type contextKey struct {
	name string
}

// Jwt handles user authorization via JSON Web Tokens.
func Jwt(
	logger *log.Logger,
	client *ent.Client,
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

				ctx := request.Context()

				// validate jwt token
				username, err := util.ParseToken(ctx, header)
				if err != nil {
					logger.Debugf("invalid token: %v", err)

					writer.Header().Set("Content-Type", "application/json")
					writer.WriteHeader(http.StatusUnauthorized)
					_, _ = writer.Write([]byte(getError(err)))

					return
				}

				entUser, err := repository.
					NewUserRepository(client).
					Get(ctx, username)
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
		return nil, model.NewAuthorizationError(ctx, "invalid token")
	}

	return entUser, nil
}

func getError(e error) string {
	return fmt.Sprintf(errorFormatString, e.Error())
}
