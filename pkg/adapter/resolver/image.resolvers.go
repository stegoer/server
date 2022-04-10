package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/stegoer/server/ent"
	"github.com/stegoer/server/ent/schema/ulid"
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
	entUser, err := middleware.JwtForContext(ctx)
	if err != nil {
		r.logger.Infow("encode: user not authenticated", "error", err.Error())
	}

	if err := steganography.ValidateEncodeInput(
		ctx,
		entUser,
		input,
	); err != nil {
		r.logger.Errorw("encode: invalid input",
			"error", err.Error(),
			"input", input,
			"user", entUser,
		)

		return nil, err
	}

	content, err := steganography.Encode(input)
	if err != nil {
		r.logger.Errorw(
			"encode: failure",
			"error", err.Error(),
			"input", input,
			"user", entUser,
		)

		return nil, model.NewValidationError(ctx, err.Error())
	}

	if entUser != nil {
		if _, err := r.controller.Image.Create(
			ctx,
			*entUser,
			input.Upload.Filename,
			content,
		); err != nil {
			r.logger.Errorw("encode: failed to create image",
				"error", err.Error(),
				"filename", input.Upload.Filename,
				"user", entUser,
			)

			return nil, model.NewDBError(ctx, "failed to create image")
		}
	}

	r.logger.Debugw("encode: success", "user", entUser)

	return &generated.EncodeImagePayload{
		File: &generated.FileType{
			Name:    input.Upload.Filename,
			Content: content,
		},
	}, nil
}

func (r *mutationResolver) DecodeImage(ctx context.Context, input generated.DecodeImageInput) (*generated.DecodeImagePayload, error) {
	entUser, err := middleware.JwtForContext(ctx)
	if err != nil {
		r.logger.Infow("decode: user not authenticated", "error", err.Error())
	}

	if err := steganography.ValidateDecodeInput(
		ctx,
		entUser,
		input,
	); err != nil {
		r.logger.Errorw("decode: invalid input",
			"error", err.Error(),
			"input", input,
			"user", entUser,
		)

		return nil, err
	}

	data, err := steganography.Decode(input)
	if err != nil {
		r.logger.Errorw(
			"decode: failure",
			"error", err.Error(),
			"input", input,
			"user", entUser,
		)

		return nil, model.NewValidationError(ctx, err.Error())
	}

	r.logger.Debugw("decode: success", "user", entUser)

	return &generated.DecodeImagePayload{Data: data}, nil
}

func (r *queryResolver) Image(ctx context.Context, id ulid.ID) (*ent.Image, error) {
	entUser, err := middleware.JwtForContext(ctx)
	if err != nil {
		r.logger.Errorw("image: unauthenticated user",
			"error", err.Error(),
			"image_id", id,
		)

		return nil, model.NewAuthorizationError(ctx, "user is not authenticated")
	}

	entImage, err := r.controller.Image.Get(
		ctx,
		*entUser,
		&id,
	)
	if err != nil {
		r.logger.Errorw(
			"image: not found",
			"error", err.Error(),
			"image_id", id,
			"user_id", entUser.ID,
		)

		return nil, model.NewNotFoundError(ctx, "image not found")
	}

	r.logger.Debugw("image: found", "image_id", entImage.ID, "user_id", entUser.ID)

	return entImage, nil
}

func (r *queryResolver) Images(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.ImageWhereInput, orderBy *ent.ImageOrder) (*generated.ImagesConnection, error) {
	entUser, err := middleware.JwtForContext(ctx)
	if err != nil {
		r.logger.Errorw("images: unauthenticated user", "error", err.Error())

		return nil, model.NewAuthorizationError(ctx, "user is not authenticated")
	}

	if first == nil && last == nil {
		return nil, model.NewValidationError(
			ctx,
			"query must specify first or last",
		)
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
		r.logger.Errorw("images: failed to list",
			"error", err.Error(),
			"user_id", entUser.ID,
			"after", after,
			"before", before,
			"last", last,
			"where", where,
			"orderBy", orderBy,
		)

		return nil, model.NewDBError(ctx, "failed to list images")
	}

	r.logger.Debugw(
		"images: successfully listed",
		"page_info", imageList.PageInfo,
	)

	return &generated.ImagesConnection{
		TotalCount: imageList.TotalCount,
		PageInfo:   &imageList.PageInfo,
		Edges:      imageList.Edges,
	}, nil
}

// Image returns generated.ImageResolver implementation.
func (r *Resolver) Image() generated.ImageResolver { return &imageResolver{r} }

type imageResolver struct{ *Resolver }
