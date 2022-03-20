package repository

import (
	"context"

	"github.com/kucera-lukas/stegoer/ent"
	"github.com/kucera-lukas/stegoer/ent/image"
	"github.com/kucera-lukas/stegoer/graph/generated"
	"github.com/kucera-lukas/stegoer/pkg/adapter/controller"
	"github.com/kucera-lukas/stegoer/pkg/entity/model"
)

// NewImageRepository returns implementation of the controller.Image interface.
func NewImageRepository(client *ent.Client) controller.Image { //nolint:ireturn
	return &imageRepository{client: client}
}

type imageRepository struct {
	client *ent.Client
}

func (r *imageRepository) Get(
	ctx context.Context,
	entUser model.User,
	id *model.ID,
) (*model.Image, error) {
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
	if first == nil && last == nil {
		return nil, model.NewBadRequestError(
			ctx,
			"query must specify first or last",
		)
	}

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
	input generated.EncodeImageInput,
) (*model.Image, error) {
	entImage, err := r.client.
		Image.
		Create().
		SetMessage(input.Message).
		SetLsbUsed(input.LsbUsed).
		SetChannel(input.Channel).
		SetUser(&entUser).
		Save(ctx)
	if err != nil {
		return nil, model.NewDBError(ctx, err.Error())
	}

	return entImage, nil
}

func (r *imageRepository) Count(
	ctx context.Context,
	entUser model.User,
) (int, error) {
	count, err := entUser.QueryImages().Count(ctx)
	if err != nil {
		return 0, model.NewDBError(ctx, err.Error())
	}

	return count, nil
}
