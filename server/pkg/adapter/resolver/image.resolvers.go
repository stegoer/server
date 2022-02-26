package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kucera-lukas/stegoer/ent"
	"github.com/kucera-lukas/stegoer/graph/generated"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/middleware"
)

func (r *mutationResolver) CreateImage(ctx context.Context, input generated.NewImage) (*generated.CreateImagePayload, error) {
	entUser, err := middleware.JwtForContext(ctx)
	if err != nil {
		return &generated.CreateImagePayload{Image: nil}, err
	}

	entImage, err := r.controller.Image.Create(ctx, *entUser, input)
	if err != nil {
		return &generated.CreateImagePayload{Image: nil}, err
	}

	return &generated.CreateImagePayload{Image: entImage}, nil
}

func (r *queryResolver) Images(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.ImageWhereInput, orderBy *ent.ImageOrder) (*generated.ImagesConnection, error) {
	entUser, err := middleware.JwtForContext(ctx)
	if err != nil {
		return &generated.ImagesConnection{
			TotalCount: 0,
			PageInfo: &ent.PageInfo{
				HasNextPage:     false,
				HasPreviousPage: false,
				StartCursor:     nil,
				EndCursor:       nil,
			},
			Edges: []*ent.ImageEdge{},
		}, err
	}

	imageList, err := r.controller.Image.List(
		ctx,
		*entUser,
		after,
		first,
		before,
		last,
		where,
		orderBy,
	)
	if err != nil {
		return &generated.ImagesConnection{
			TotalCount: 0,
			PageInfo: &ent.PageInfo{
				HasNextPage:     false,
				HasPreviousPage: false,
				StartCursor:     nil,
				EndCursor:       nil,
			},
			Edges: []*ent.ImageEdge{},
		}, err
	}

	return &generated.ImagesConnection{
		TotalCount: imageList.TotalCount,
		PageInfo:   &imageList.PageInfo,
		Edges:      imageList.Edges,
	}, nil
}
