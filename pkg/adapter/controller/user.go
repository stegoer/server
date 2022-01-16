package controller

import (
	"StegoLSB/pkg/entity/model"
	"StegoLSB/pkg/usecase/usecase"
	"context"
)

type user struct {
	userUsecase usecase.User
}

// User of interface
type User interface {
	Get(ctx context.Context, name string) (*model.User, error)
	Create(ctx context.Context, input model.NewUserInput) (*model.User, error)
	Update(ctx context.Context, entUser model.User, input model.UpdateUserInput) (*model.User, error)
}

// NewUserController returns user controller
func NewUserController(uu usecase.User) User {
	return &user{userUsecase: uu}
}

func (u *user) Get(ctx context.Context, name string) (*model.User, error) {
	return u.userUsecase.Get(ctx, name)
}

func (u *user) Create(ctx context.Context, input model.NewUserInput) (*model.User, error) {
	return u.userUsecase.Create(ctx, input)
}

func (u *user) Update(ctx context.Context, entUser model.User, input model.UpdateUserInput) (*model.User, error) {
	return u.userUsecase.Update(ctx, entUser, input)
}
