package controller

import (
	"context"

	"github.com/stegoer/server/graph/generated"
	"github.com/stegoer/server/pkg/model"
)

// User controller interface.
type User interface {
	GetByID(
		ctx context.Context,
		id model.ID,
	) (*model.User, error)
	GetByEmail(
		ctx context.Context,
		email string,
	) (*model.User, error)
	Create(
		ctx context.Context,
		input generated.NewUser,
	) (*model.User, error)
	Update(
		ctx context.Context,
		entUser model.User,
		input generated.UpdateUser,
	) (*model.User, error)
	SetLoggedIn(
		ctx context.Context,
		entUser model.User,
	) (*model.User, error)
}
