package repository

import (
	"StegoLSB/ent"
	"StegoLSB/ent/user"
	"StegoLSB/pkg/entity/model"
	usecaseRepository "StegoLSB/pkg/usecase/repository"
	"StegoLSB/pkg/util"
	"context"
)

type userRepository struct {
	client *ent.Client
}

// NewUserRepository is specific implementation of the interface
func NewUserRepository(client *ent.Client) usecaseRepository.User {
	return &userRepository{client: client}
}

func (r *userRepository) Get(ctx context.Context, name string) (*model.User, error) {
	entUser, err := r.client.User.Query().Where(user.NameEQ(name)).Only(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return entUser, nil
}

func (r *userRepository) Create(ctx context.Context, input model.NewUserInput) (*model.User, error) {
	hashedPassword, err := util.HashPassword(input.Password)

	if err != nil {
		return nil, model.NewValidationError(err)
	}

	entUser, err := r.client.User.Create().SetName(input.Username).SetPassword(hashedPassword).Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return entUser, nil
}

func (r *userRepository) Update(ctx context.Context, entUser model.User, input model.UpdateUserInput) (*model.User, error) {
	update := entUser.Update()

	if input.Name != nil {
		update = update.SetName(*input.Name)
	}

	if input.Password != nil {
		hashedPassword, err := util.HashPassword(*input.Password)
		if err != nil {
			return nil, model.NewValidationError(err)
		}

		update = update.SetPassword(hashedPassword)
	}

	updatedEntUser, err := update.Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return updatedEntUser, nil
}
