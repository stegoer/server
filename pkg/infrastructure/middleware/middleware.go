package middleware

import (
	"StegoLSB/ent"
	"StegoLSB/pkg/entity/model"
	"StegoLSB/pkg/util"
	"github.com/pkg/errors"

	"StegoLSB/user"
	"context"
	"fmt"
	"net/http"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// JwtMiddleware handler user authorization via JWT tokens
func JwtMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)

				return
			}

			// validate jwt token
			username, err := util.ParseToken(header)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				_, err := w.Write([]byte(
					fmt.Sprintf(
						`{"errors":[{"message": "%v", "path": ["Authorization"]}], data: null}`,
						err.Error(),
					)))
				if err != nil {
					return
				}

				return
			}

			entUser, err := user.GetUserByUsername(r.Context(), username)
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

// ForContext finds the user id from the context. Requires Middleware to have run.
func ForContext(ctx context.Context) (*ent.User, error) {
	entUser, ok := ctx.Value(userCtxKey).(*ent.User)
	if !ok {
		return nil, model.NewAuthorizationError(errors.Errorf("invalid token"))
	}

	return entUser, nil
}
