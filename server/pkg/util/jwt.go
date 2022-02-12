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

// SecretKey used to sign tokens.
var (
	SecretKey = []byte(os.Getenv("SECRET_KEY")) //nolint:gochecknoglobals
)

// GenerateAuth generates a jwt token and assigns a username to its claims.
func GenerateAuth(
	ctx context.Context,
	entUser model.User,
) (*generated.Auth, *model.Error) {
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

	return &generated.Auth{
		Token:   tokenString,
		Expires: exp,
	}, nil
}

// ParseToken parses a jwt token and returns the username in its claims.
func ParseToken(
	ctx context.Context,
	tokenStr string,
) (string, *model.Error) {
	token, err := jwt.Parse(tokenStr, func(_ *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return "", model.NewAuthorizationError(ctx, err.Error())
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return fmt.Sprintf("%v", claims["username"]), nil
	}

	return "", model.NewAuthorizationError(ctx, mapClaimsErrorMessage)
}
