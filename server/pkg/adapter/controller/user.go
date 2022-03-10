package controller

import (
	"context"

	"github.com/kucera-lukas/stegoer/ent/schema/ulid"
	"github.com/kucera-lukas/stegoer/graph/generated"
	"github.com/kucera-lukas/stegoer/pkg/entity/model"
)

// User controller interface.
type User interface {
	GetByID(
		ctx context.Context,
		id ulid.ID,
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
