package controller

import (
	"StegoLSB/pkg/entity/model"
	"StegoLSB/pkg/usecase/usecase"
	"context"
)

// Image is an interface of controller
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

type image struct {
	imageUsecase usecase.Image
}

// NewImageController generates test user controller
func NewImageController(iu usecase.Image) Image {
	return &image{
		imageUsecase: iu,
	}
}

func (t *image) Get(ctx context.Context, entUser model.User, id *model.ID) (*model.Image, error) {
	return t.imageUsecase.Get(ctx, entUser, id)
}

func (t *image) List(ctx context.Context,
	entUser model.User,
	after *model.Cursor,
	first *int,
	before *model.Cursor,
	last *int,
	where *model.ImageWhereInput,
	orderBy *model.ImageOrderInput,
) (*model.ImageConnection, error) {
	return t.imageUsecase.List(ctx, entUser, after, first, before, last, where, orderBy)
}

func (t *image) Create(ctx context.Context, entUser model.User, input model.NewImageInput) (*model.Image, error) {
	return t.imageUsecase.Create(ctx, entUser, input)
}

func (t *image) Count(ctx context.Context, entUser model.User) (int, error) {
	return t.imageUsecase.Count(ctx, entUser)
}
