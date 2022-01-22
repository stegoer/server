package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kucera-lukas/stegoer/ent"
	"github.com/kucera-lukas/stegoer/graph/generated"
	"github.com/kucera-lukas/stegoer/pkg/entity/model"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/middleware"
)

func (r *mutationResolver) CreateImage(ctx context.Context, input generated.NewImage) (*ent.Image, error) { //nolint:lll
	entUser, err := middleware.JwtForContext(ctx)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return r.controller.Image.Create(ctx, *entUser, input) //nolint:wrapcheck
}

func (r *queryResolver) Images(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.ImageWhereInput, orderBy *ent.ImageOrder) (*ent.ImageConnection, error) { //nolint:lll
	entUser, err := middleware.JwtForContext(ctx)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	if first == nil || last == nil {
		return nil, model.NewBadRequestError(ctx, "query must specify first or last")
	}

	return r.controller.Image.List( //nolint:wrapcheck
		ctx,
		*entUser,
		after, first,
		before,
		last,
		where,
		orderBy,
	)
}
