package controller

import (
	"context"

	"github.com/kucera-lukas/stegoer/graph/generated"
	"github.com/kucera-lukas/stegoer/pkg/entity/model"
)

// User controller interface.
type User interface {
	Get(
		ctx context.Context,
		name string,
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
}
