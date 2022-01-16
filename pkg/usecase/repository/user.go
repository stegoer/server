package repository

import (
	"StegoLSB/graph/generated"
	"StegoLSB/pkg/entity/model"
	"context"
)

// User is interface of repository
type User interface {
	Get(ctx context.Context, name string) (*model.User, error)
	Create(ctx context.Context, input generated.NewUser) (*model.User, error)
	Update(ctx context.Context, entUser model.User, input model.UpdateUserInput) (*model.User, error)
}
