package repository

import (
	"StegoLSB/pkg/entity/model"
	"context"
)

// Image is interface of repository
type Image interface {
	Get(ctx context.Context, entUser model.User, id *model.ID) (*model.Image, error)
	List(
		ctx context.Context,
		entUser model.User,
		after *model.Cursor,
		first *int,
		before *model.Cursor,
		last *int,
		where *model.ImageWhereInput,
		orderBy *model.ImageOrderInput,
	) (*model.ImageConnection, error)
	Create(ctx context.Context, entUser model.User, input model.NewImageInput) (*model.Image, error)
	Count(ctx context.Context, entUser model.User) (int, error)
}
