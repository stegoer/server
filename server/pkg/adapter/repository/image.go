package repository

import (
	"context"
	"stegoer/ent"
	"stegoer/ent/image"
	"stegoer/pkg/adapter/controller"
	"stegoer/pkg/entity/model"
)

// NewImageRepository returns a specific implementation of the controller.Image interface
func NewImageRepository(client *ent.Client) controller.Image {
	return &imageRepository{client: client}
}

type imageRepository struct {
	client *ent.Client
}

func (r *imageRepository) Get(ctx context.Context, entUser model.User, id *model.ID) (*model.Image, error) {
	entImage, err := entUser.QueryImages().Where(image.ID(*id)).Only(ctx)

	if err != nil {
		return nil, model.NewDBError(ctx, err.Error())
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
		return nil, model.NewDBError(ctx, err.Error())
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
		return nil, model.NewDBError(ctx, err.Error())
	}

	return entImage, nil
}

func (r *imageRepository) Count(ctx context.Context, entUser model.User) (int, error) {
	count, err := entUser.QueryImages().Count(ctx)

	if err != nil {
		return 0, model.NewDBError(ctx, err.Error())
	}

	return count, nil
}
