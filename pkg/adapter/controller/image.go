package controller

import (
	"context"

	"github.com/stegoer/server/pkg/model"
)

// Image controller interface.
type Image interface {
	Get(
		ctx context.Context,
		entUser model.User,
		id *model.ID,
	) (*model.Image, error)
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
	Create(
		ctx context.Context,
		entUser model.User,
		filename string,
		content string,
	) (*model.Image, error)
	Count(
		ctx context.Context,
		entUser model.User,
	) (int, error)
}
