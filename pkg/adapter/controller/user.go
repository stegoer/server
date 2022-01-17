package controller

import (
	"StegoLSB/pkg/entity/model"
	"context"
)

// User controller interface
type User interface {
	Get(ctx context.Context, name string) (*model.User, error)
	Create(ctx context.Context, input model.NewUserInput) (*model.User, error)
	Update(ctx context.Context, entUser model.User, input model.UpdateUserInput) (*model.User, error)
}
