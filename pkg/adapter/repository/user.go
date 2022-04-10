package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/stegoer/server/ent"
	"github.com/stegoer/server/ent/schema/ulid"
	"github.com/stegoer/server/ent/user"
	"github.com/stegoer/server/gqlgen"
	"github.com/stegoer/server/pkg/adapter/controller"
	"github.com/stegoer/server/pkg/cryptography"
)

// NewUserRepository returns implementation of the controller.User interface.
func NewUserRepository(client *ent.Client) controller.User { //nolint:ireturn
	return &userRepository{client: client}
}

type userRepository struct {
	client *ent.Client
}

func (r *userRepository) GetByID(
	ctx context.Context,
	id ulid.ID,
) (*ent.User, error) {
	entUser, err := r.client.User.Query().Where(user.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("getByID: %w", err)
	}

	return entUser, nil
}

func (r *userRepository) GetByEmail(
	ctx context.Context,
	email string,
) (*ent.User, error) {
	entUser, err := r.client.User.Query().Where(user.EmailEQ(email)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("getByEmail: %w", err)
	}

	return entUser, nil
}

func (r *userRepository) Create(
	ctx context.Context,
	input gqlgen.NewUser,
) (*ent.User, error) {
	hashedPassword, err := cryptography.HashPassword(input.Password)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	entUser, err := r.client.User.
		Create().
		SetName(input.Username).
		SetEmail(input.Email).
		SetPassword(hashedPassword).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	return entUser, nil
}

func (r *userRepository) Update(
	ctx context.Context,
	entUser ent.User,
	input gqlgen.UpdateUser,
) (*ent.User, error) {
	update := entUser.Update()

	if input.Username != nil {
		update = update.SetName(*input.Username)
	}

	if input.Email != nil {
		update = update.SetEmail(*input.Email)
	}

	if input.Password != nil {
		hashedPassword, err := cryptography.HashPassword(*input.Password)
		if err != nil {
			return nil, fmt.Errorf("update: %w", err)
		}

		update = update.SetPassword(hashedPassword)
	}

	updatedEntUser, err := update.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("update: %w", err)
	}

	return updatedEntUser, nil
}

func (r *userRepository) SetLoggedIn(
	ctx context.Context,
	entUser ent.User,
) (*ent.User, error) {
	updatedEntUser, err := entUser.Update().SetLastLogin(time.Now()).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("setLoggedIn: %w", err)
	}

	return updatedEntUser, nil
}
