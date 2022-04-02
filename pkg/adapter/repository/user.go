package repository

import (
	"context"
	"time"

	"github.com/stegoer/server/ent"
	"github.com/stegoer/server/ent/schema/ulid"
	"github.com/stegoer/server/ent/user"
	"github.com/stegoer/server/graph/generated"
	"github.com/stegoer/server/pkg/adapter/controller"
	model2 "github.com/stegoer/server/pkg/model"
	"github.com/stegoer/server/pkg/util"
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
) (*model2.User, error) {
	entUser, err := r.client.User.Query().Where(user.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, model2.NewDBError(ctx, err.Error())
	}

	return entUser, nil
}

func (r *userRepository) GetByEmail(
	ctx context.Context,
	email string,
) (*model2.User, error) {
	entUser, err := r.client.User.Query().Where(user.EmailEQ(email)).Only(ctx)
	if err != nil {
		return nil, model2.NewDBError(ctx, err.Error())
	}

	return entUser, nil
}

func (r *userRepository) Create(
	ctx context.Context,
	input generated.NewUser,
) (*model2.User, error) {
	hashedPassword, err := util.HashPassword(input.Password)
	if err != nil {
		return nil, model2.NewValidationError(ctx, err.Error())
	}

	entUser, err := r.client.User.
		Create().
		SetName(input.Username).
		SetEmail(input.Email).
		SetPassword(hashedPassword).
		Save(ctx)
	if err != nil {
		return nil, model2.NewDBError(ctx, err.Error())
	}

	return entUser, nil
}

func (r *userRepository) Update(
	ctx context.Context,
	entUser model2.User,
	input generated.UpdateUser,
) (*model2.User, error) {
	update := entUser.Update()

	if input.Username != nil {
		update = update.SetName(*input.Username)
	}

	if input.Email != nil {
		update = update.SetEmail(*input.Email)
	}

	if input.Password != nil {
		hashedPassword, err := util.HashPassword(*input.Password)
		if err != nil {
			return nil, model2.NewValidationError(ctx, err.Error())
		}

		update = update.SetPassword(hashedPassword)
	}

	updatedEntUser, err := update.Save(ctx)
	if err != nil {
		return nil, model2.NewDBError(ctx, err.Error())
	}

	return updatedEntUser, nil
}

func (r *userRepository) SetLoggedIn(
	ctx context.Context,
	entUser model2.User,
) (*model2.User, error) {
	updatedEntUser, err := entUser.Update().SetLastLogin(time.Now()).Save(ctx)
	if err != nil {
		return nil, model2.NewDBError(ctx, err.Error())
	}

	return updatedEntUser, nil
}
