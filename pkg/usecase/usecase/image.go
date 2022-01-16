package usecase

import (
	"StegoLSB/pkg/entity/model"
	"StegoLSB/pkg/usecase/repository"
	"context"
)

type image struct {
	imageRepository repository.Image
}

// Image of usecase
type Image interface {
	Get(ctx context.Context, entUser model.User, id *model.ID) (*model.Image, error)
	List(ctx context.Context,
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

// NewImageUsecase generates test user repository
func NewImageUsecase(r repository.Image) Image {
	return &image{imageRepository: r}
}

func (t *image) Get(ctx context.Context, entUser model.User, id *model.ID) (*model.Image, error) {
	return t.imageRepository.Get(ctx, entUser, id)
}

func (t *image) List(
	ctx context.Context,
	entUser model.User,
	after *model.Cursor,
	first *int,
	before *model.Cursor,
	last *int,
	where *model.ImageWhereInput,
	orderBy *model.ImageOrderInput,
) (*model.ImageConnection, error) {
	return t.imageRepository.List(ctx, entUser, after, first, before, last, where, orderBy)
}

func (t *image) Create(ctx context.Context, entUser model.User, input model.NewImageInput) (*model.Image, error) {
	return t.imageRepository.Create(ctx, entUser, input)
}

func (t *image) Count(ctx context.Context, entUser model.User) (int, error) {
	return t.imageRepository.Count(ctx, entUser)
}
