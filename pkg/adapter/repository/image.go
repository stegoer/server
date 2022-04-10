package repository

import (
	"context"
	"fmt"

	"github.com/stegoer/server/ent"
	"github.com/stegoer/server/ent/image"
	"github.com/stegoer/server/ent/schema/ulid"
	"github.com/stegoer/server/pkg/adapter/controller"
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
	entUser ent.User,
	id *ulid.ID,
) (*ent.Image, error) {
	entImage, err := entUser.
		QueryImages().
		Where(image.ID(*id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return entImage, nil
}

func (r *imageRepository) List(ctx context.Context,
	entUser ent.User,
	after *ent.Cursor,
	first *int,
	before *ent.Cursor,
	last *int,
	where *ent.ImageWhereInput,
	orderBy *ent.ImageOrder,
) (*ent.ImageConnection, error) {
	connection, err := entUser.QueryImages().
		Paginate(ctx, after, first, before, last,
			ent.WithImageFilter(where.Filter),
			ent.WithImageOrder(orderBy),
		)
	if err != nil {
		return nil, fmt.Errorf("list: %w", err)
	}

	return connection, nil
}

func (r *imageRepository) Create(
	ctx context.Context,
	entUser ent.User,
	filename string,
	content string,
) (*ent.Image, error) {
	entImage, err := r.client.
		Image.
		Create().
		SetFileName(filename).
		SetContent(content).
		SetUser(&entUser).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	return entImage, nil
}

func (r *imageRepository) Count(
	ctx context.Context,
	entUser ent.User,
) (int, error) {
	count, err := entUser.QueryImages().Count(ctx)
	if err != nil {
		return 0, fmt.Errorf("count: %w", err)
	}

	return count, nil
}
