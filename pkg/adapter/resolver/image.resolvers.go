package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"StegoLSB/ent"
	"StegoLSB/graph/generated"
	"StegoLSB/pkg/entity/model"
	"StegoLSB/pkg/infrastructure/middleware"
	"context"
)

func (r *mutationResolver) CreateImage(ctx context.Context, input generated.NewImage) (*ent.Image, error) {
	entUser, err := middleware.ForContext(ctx)
	if err != nil {
		return nil, model.NewAuthorizationError(ctx, err.Error())
	}

	return r.controller.Image.Create(ctx, *entUser, input)
}

func (r *queryResolver) Images(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.ImageWhereInput, orderBy *ent.ImageOrder) (*ent.ImageConnection, error) {
	entUser, err := middleware.ForContext(ctx)

	if err != nil {
		return nil, model.NewAuthorizationError(ctx, err.Error())
	}

	if first == nil || last == nil {
		return nil, model.NewBadRequestError(ctx, "query must specify first or last")
	}

	return r.controller.Image.List(ctx, *entUser, after, first, before, last, where, orderBy)
}
