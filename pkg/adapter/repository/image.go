package repository

import (
	"StegoLSB/ent"
	"StegoLSB/ent/image"
	"StegoLSB/pkg/entity/model"
	usecaseRepository "StegoLSB/pkg/usecase/repository"
	"context"
)

type imageRepository struct {
	client *ent.Client
}

// NewImageRepository generates new repository
func NewImageRepository(client *ent.Client) usecaseRepository.Image {
	return &imageRepository{client: client}
}

func (r *imageRepository) Get(ctx context.Context, entUser model.User, id *model.ID) (*model.Image, error) {
	entImage, err := entUser.QueryImages().Where(image.ID(*id)).Only(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return entImage, nil
}

func (r *imageRepository) List(ctx context.Context,
	entUser model.User,
	after *model.Cursor,
	first *int,
	before *model.Cursor,
	last *int,
	where *model.ImageWhereInput,
	orderBy *model.ImageOrderInput,
) (*model.ImageConnection, error) {
	connection, err := entUser.QueryImages().
		Paginate(ctx, after, first, before, last,
			ent.WithImageFilter(where.Filter),
			ent.WithImageOrder(orderBy),
		)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return connection, nil
}

func (r *imageRepository) Create(
	ctx context.Context,
	entUser model.User,
	input model.NewImageInput,
) (*model.Image, error) {
	entImage, err := r.client.
		Image.
		Create().
		SetChannel(input.Channel).
		SetUser(&entUser).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return entImage, nil
}

func (r *imageRepository) Count(ctx context.Context, entUser model.User) (int, error) {
	return entUser.QueryImages().Count(ctx)
}
