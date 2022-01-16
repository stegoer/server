package util

import (
	"StegoLSB/graph/generated"
	"StegoLSB/pkg/entity/model"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"os"
	"time"
)

const tokenExpiration = time.Minute * 15

// SecretKey being used to sign tokens
var (
	SecretKey = []byte(os.Getenv("SECRET_KEY")) //nolint:gochecknoglobals
)

type mapClaimsError struct{}

func (m *mapClaimsError) Error() string {
	return "failed to convert token claims to standard claims"
}

// GenerateAuthUser generates a jwt token and assigns a username to it's claims
func GenerateAuthUser(entUser model.User) (*generated.AuthUser, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, model.NewAuthorizationError(&mapClaimsError{})
	}

	claims["username"] = entUser.Name
	exp := time.Now().Add(tokenExpiration)
	claims["exp"] = exp.Unix()

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Printf(`error generating key for user %s`, entUser.ID)

		return nil, model.NewInternalServerError(err)
	}

	return &generated.AuthUser{
		Auth: &generated.Auth{
			Ok:      true,
			Token:   tokenString,
			Expires: FormatDate(exp),
		},
		User: &entUser,
	}, nil
}

// ParseToken parses a jwt token and returns the username in its claims
func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return "", model.NewAuthorizationError(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return fmt.Sprintf("%v", claims["username"]), nil
	}
	return "", model.NewAuthorizationError(&mapClaimsError{})
}
