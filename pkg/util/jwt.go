package util

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/stegoer/server/graph/generated"
	"github.com/stegoer/server/pkg/model"
)

const (
	tokenExpiration       = time.Minute * 120
	mapClaimsErrorMessage = "failed to convert token claims to standard claims"
)

// secretKey used to sign tokens.
var secretKey = []byte(os.Getenv("SECRET_KEY")) //nolint:gochecknoglobals

// GenerateAuth generates a jwt token and assigns a username to its claims.
func GenerateAuth(
	ctx context.Context,
	entUser model.User,
) (*generated.Auth, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New(mapClaimsErrorMessage)
	}

	claims["id"] = entUser.ID
	exp := time.Now().Add(tokenExpiration)
	claims["exp"] = exp.Unix()

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, model.NewInternalServerError(ctx, err.Error())
	}

	return &generated.Auth{
		Token:   tokenString,
		Expires: exp,
	}, nil
}

// ParseToken parses a jwt token and returns the model.ID in its claims.
func ParseToken(tokenStr string) (model.ID, error) {
	token, err := jwt.Parse(
		Trim(tokenStr, '"'),
		func(_ *jwt.Token) (interface{}, error) {
			return secretKey, nil
		},
	)
	if err != nil {
		return "", fmt.Errorf("parse_token: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return model.ID(fmt.Sprintf("%v", claims["id"])), nil
	}

	return "", errors.New(mapClaimsErrorMessage)
}
