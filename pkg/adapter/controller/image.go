package controller

import (
	"context"

	model2 "github.com/stegoer/server/pkg/model"
)

// Image controller interface.
type Image interface {
	Get(
		ctx context.Context,
		entUser model2.User,
		id *model2.ID,
	) (*model2.Image, error)
	List(
		ctx context.Context,
		entUser model2.User,
		after *model2.Cursor,
		first *int,
		before *model2.Cursor,
		last *int,
		where *model2.ImageWhereInput,
		orderBy *model2.ImageOrderInput,
	) (*model2.ImageConnection, error)
	Create(
		ctx context.Context,
		entUser model2.User,
		filename string,
		content string,
	) (*model2.Image, error)
	Count(
		ctx context.Context,
		entUser model2.User,
	) (int, error)
}
