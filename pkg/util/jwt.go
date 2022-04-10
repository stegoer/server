package util

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/stegoer/server/ent"
	"github.com/stegoer/server/ent/schema/ulid"
)

const (
	tokenExpiration       = time.Minute * 120
	mapClaimsErrorMessage = "failed to convert token claims to standard claims"
)

// secretKey used to sign tokens.
var secretKey = []byte(os.Getenv("SECRET_KEY")) //nolint:gochecknoglobals

// GenerateToken created a jwt token and puts an ulid.ID into its claims.
func GenerateToken(
	ctx context.Context,
	entUser ent.User,
) (string, *time.Time, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", nil, errors.New(mapClaimsErrorMessage)
	}

	claims["id"] = entUser.ID
	exp := time.Now().Add(tokenExpiration)
	claims["exp"] = exp.Unix()

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", nil, NewInternalServerError(ctx, err.Error())
	}

	return tokenString, &exp, nil
}

// ParseToken parses a jwt token and returns the ulid.ID in its claims.
func ParseToken(tokenStr string) (ulid.ID, error) {
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
		return ulid.ID(fmt.Sprintf("%v", claims["id"])), nil
	}

	return "", errors.New(mapClaimsErrorMessage)
}
