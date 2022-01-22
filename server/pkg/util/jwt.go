package util

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/kucera-lukas/stegoer/graph/generated"
	"github.com/kucera-lukas/stegoer/pkg/entity/model"
)

const (
	tokenExpiration       = time.Minute * 15
	mapClaimsErrorMessage = "failed to convert token claims to standard claims"
)

// SecretKey being used to sign tokens.
var (
	SecretKey = []byte(os.Getenv("SECRET_KEY")) //nolint:gochecknoglobals
)

// GenerateAuthUser generates a jwt token and assigns a username to its claims.
func GenerateAuthUser(
	ctx context.Context,
	entUser model.User,
) (*generated.AuthUser, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, model.NewInternalServerError(ctx, mapClaimsErrorMessage)
	}

	claims["username"] = entUser.Name
	exp := time.Now().Add(tokenExpiration)
	claims["exp"] = exp.Unix()

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return nil, model.NewInternalServerError(ctx, err.Error())
	}

	return &generated.AuthUser{
		Auth: &generated.Auth{
			Ok:      true,
			Token:   tokenString,
			Expires: exp,
		},
		User: &entUser,
	}, nil
}

// ParseToken parses a jwt token and returns the username in its claims.
func ParseToken(ctx context.Context, tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return "", model.NewAuthorizationError(ctx, err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return fmt.Sprintf("%v", claims["username"]), nil
	}

	return "", model.NewAuthorizationError(ctx, mapClaimsErrorMessage)
}
