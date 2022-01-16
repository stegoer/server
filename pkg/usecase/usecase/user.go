package usecase

import (
	"StegoLSB/pkg/entity/model"
	"StegoLSB/pkg/usecase/repository"
	"context"
)

type user struct {
	userRepository repository.User
}

// User of usecase.
type User interface {
	Get(ctx context.Context, name string) (*model.User, error)
	Create(ctx context.Context, input model.NewUserInput) (*model.User, error)
	Update(ctx context.Context, entUser model.User, input model.UpdateUserInput) (*model.User, error)
}

// NewUserUsecase returns User usecase.
func NewUserUsecase(r repository.User) User {
	return &user{userRepository: r}
}

func (u *user) Get(ctx context.Context, name string) (*model.User, error) {
	return u.userRepository.Get(ctx, name)
}

func (u *user) Create(ctx context.Context, input model.NewUserInput) (*model.User, error) {
	return u.userRepository.Create(ctx, input)
}

func (u *user) Update(ctx context.Context, entUser model.User, input model.UpdateUserInput) (*model.User, error) {
	return u.userRepository.Update(ctx, entUser, input)
}
