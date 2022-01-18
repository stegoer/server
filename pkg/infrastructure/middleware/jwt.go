package middleware

import (
	"StegoLSB/ent"
	"StegoLSB/pkg/adapter/repository"
	"StegoLSB/pkg/entity/model"
	"StegoLSB/pkg/util"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

const (
	errorFormatString = `{"errors":[{"message": "%v", "path": ["Authorization"], "extensions": %v}], data: null}`
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// Jwt handles user authorization via JSON Web Tokens.
func Jwt(client *ent.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)

				return
			}

			// validate jwt token
			username, err := util.ParseToken(r.Context(), header)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				_, err := w.Write([]byte(getError(err)))
				if err != nil {
					return
				}

				return
			}

			entUser, err := repository.NewUserRepository(client).Get(r.Context(), username)
			if err != nil {
				next.ServeHTTP(w, r)

				return
			}

			// put user into context
			ctx := context.WithValue(r.Context(), userCtxKey, entUser)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// JwtForContext finds the user id from the context. Requires Jwt to have run.
func JwtForContext(ctx context.Context) (*ent.User, error) {
	entUser, ok := ctx.Value(userCtxKey).(*ent.User)
	if !ok {
		return nil, errors.Errorf("invalid token")
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
