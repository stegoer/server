package user

import (
	"StegoLSB/ent"
	"StegoLSB/ent/user"
	"StegoLSB/pkg/infrastructure/client"
	"context"
)

type InvalidUsernameOrPasswordError struct{}

func (m *InvalidUsernameOrPasswordError) Error() string {
	return "username or password is invalid"
}

// GetUserByUsername checks if a user exists in database by given name
func GetUserByUsername(ctx context.Context, name string) (*ent.User, error) {
	entClient, err := client.New()
	if err != nil {
		return nil, err
	}

	entUser, err := entClient.User.Query().Where(user.NameEQ(name)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return entUser, nil
}
