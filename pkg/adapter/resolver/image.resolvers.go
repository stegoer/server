package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/stegoer/server/ent"
	"github.com/stegoer/server/ent/schema"
	"github.com/stegoer/server/graph/generated"
	"github.com/stegoer/server/pkg/entity/model"
	"github.com/stegoer/server/pkg/infrastructure/middleware"
	"github.com/stegoer/server/pkg/steganography"
	"github.com/stegoer/server/pkg/util"
)

func (r *mutationResolver) EncodeImage(ctx context.Context, input generated.EncodeImageInput) (*generated.EncodeImagePayload, error) {
	if !util.ValidLSBUsed(input.LsbUsed) {
		return nil, model.NewValidationError(
			ctx,
			fmt.Sprintf(
				"%d is out of the range [%d : %d] for least significant bit amount",
				input.LsbUsed,
				schema.LsbMin,
				schema.LsbMax,
			),
		)
	}

	imgBuffer, err := steganography.Encode(input)
	if err != nil {
		return nil, model.NewValidationError(ctx, err.Error())
	}

	entUser, err := middleware.JwtForContext(ctx)
	if err != nil {
		return nil, err
	}

	entImage, err := r.controller.Image.Create(ctx, *entUser, input)
	if err != nil {
		return nil, err
	}

	return &generated.EncodeImagePayload{
		Image: entImage,
		File: &generated.FileType{
			Name:    input.File.Filename,
			Content: base64.StdEncoding.EncodeToString(imgBuffer.Bytes()),
		},
	}, nil
}

func (r *mutationResolver) DecodeImage(ctx context.Context, input generated.DecodeImageInput) (*generated.DecodeImagePayload, error) {
	message, err := steganography.Decode(input)
	if err != nil {
		return nil, model.NewValidationError(ctx, err.Error())
	}

	return &generated.DecodeImagePayload{Message: message}, nil
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
