package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/stegoer/server/ent"
	"github.com/stegoer/server/graph/generated"
	"github.com/stegoer/server/pkg/infrastructure/middleware"
	"github.com/stegoer/server/pkg/model"
	"github.com/stegoer/server/pkg/steganography"
)

func (r *imageResolver) File(ctx context.Context, obj *ent.Image) (*generated.FileType, error) {
	return &generated.FileType{
		Name:    obj.FileName,
		Content: obj.Content,
	}, nil
}

func (r *mutationResolver) EncodeImage(ctx context.Context, input generated.EncodeImageInput) (*generated.EncodeImagePayload, error) {
	entUser, _ := middleware.JwtForContext(ctx)

	if err := steganography.ValidateEncodeInput(
		ctx,
		entUser,
		input,
	); err != nil {
		return nil, err
	}

	content, err := steganography.Encode(input)
	if err != nil {
		return nil, model.NewValidationError(ctx, err.Error())
	}

	if entUser != nil {
		if _, err := r.controller.Image.Create(
			ctx,
			*entUser,
			input.Upload.Filename,
			content,
		); err != nil {
			return nil, err
		}
	}

	return &generated.EncodeImagePayload{
		File: &generated.FileType{
			Name:    input.Upload.Filename,
			Content: content,
		},
	}, nil
}

func (r *mutationResolver) DecodeImage(ctx context.Context, input generated.DecodeImageInput) (*generated.DecodeImagePayload, error) {
	entUser, _ := middleware.JwtForContext(ctx)

	if err := steganography.ValidateDecodeInput(
		ctx,
		entUser,
		input,
	); err != nil {
		return nil, err
	}

	data, err := steganography.Decode(input)
	if err != nil {
		return nil, model.NewValidationError(ctx, err.Error())
	}

	return &generated.DecodeImagePayload{Data: data}, nil
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

// Image returns generated.ImageResolver implementation.
func (r *Resolver) Image() generated.ImageResolver { return &imageResolver{r} }

type imageResolver struct{ *Resolver }
