package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kucera-lukas/stegoer/ent"
	"github.com/kucera-lukas/stegoer/pkg/adapter/repository"
	"github.com/kucera-lukas/stegoer/pkg/entity/model"
	"github.com/kucera-lukas/stegoer/pkg/util"
)

const (
	errorFormatString = `{
"errors":[{"message": "%v", 
"path": ["Authorization"], 
"extensions": %v}],
data: null
}`
)

var userCtxKey = &contextKey{"user"} //nolint:gochecknoglobals

type contextKey struct {
	name string
}

// Jwt handles user authorization via JSON Web Tokens.
func Jwt(client *ent.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(writer http.ResponseWriter,
				request *http.Request,
			) {
				header := request.Header.Get("Authorization")

				// Allow unauthenticated users in
				if header == "" {
					next.ServeHTTP(writer, request)

					return
				}

				// validate jwt token
				username, err := util.ParseToken(request.Context(), header)
				if err != nil {
					writer.Header().Set("Content-Type", "application/json")
					writer.WriteHeader(http.StatusUnauthorized)
					_, err := writer.Write([]byte(getError(err)))
					if err != nil {
						return
					}

					return
				}

				entUser, err := repository.
					NewUserRepository(client).
					Get(request.Context(), username)
				if err != nil {
					next.ServeHTTP(writer, request)

					return
				}

				// put user into context
				ctx := context.WithValue(request.Context(), userCtxKey, entUser)

				next.ServeHTTP(writer, request.WithContext(ctx))
			})
	}
}

// JwtForContext finds the user id from the context. Requires Jwt to have run.
func JwtForContext(ctx context.Context) (*ent.User, error) {
	entUser, ok := ctx.Value(userCtxKey).(*ent.User)
	if !ok {
		return nil, model.NewAuthorizationError(ctx, "invalid token")
	}

	return entUser, nil
}

func getError(e error) string {
	return fmt.Sprintf(
		errorFormatString,
		e.Error(),
		model.Extensions{"code": model.AuthorizationError},
	)
}
