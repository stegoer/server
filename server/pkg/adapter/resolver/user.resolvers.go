package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kucera-lukas/stegoer/ent"
	"github.com/kucera-lukas/stegoer/graph/generated"
	"github.com/kucera-lukas/stegoer/pkg/entity/model"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/middleware"
	"github.com/kucera-lukas/stegoer/pkg/util"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input generated.NewUser) (*generated.AuthUser, error) {
	entUser, err := r.controller.User.Create(ctx, input)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return util.GenerateAuthUser(ctx, *entUser) //nolint:wrapcheck
}

func (r *mutationResolver) Login(ctx context.Context, input generated.Login) (*generated.AuthUser, error) {
	entUser, _ := r.controller.User.Get(ctx, input.Username)

	if entUser == nil || !util.CheckPasswordHash(
		input.Password,
		entUser.Password,
	) {
		return nil, model.NewNotFoundError(
			ctx,
			"username or password is incorrect", "user",
		)
	}

	return util.GenerateAuthUser(ctx, *entUser) //nolint:wrapcheck
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input generated.RefreshTokenInput) (*generated.AuthUser, error) {
	username, err := util.ParseToken(ctx, input.Token)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	entUser, err := r.controller.User.Get(ctx, username)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return util.GenerateAuthUser(ctx, *entUser) //nolint:wrapcheck
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input generated.UpdateUser) (*ent.User, error) {
	entUser, err := middleware.JwtForContext(ctx)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return r.controller.User.Update(ctx, *entUser, input) //nolint:wrapcheck
}

func (r *queryResolver) Overview(ctx context.Context) (*ent.User, error) {
	entUser, err := middleware.JwtForContext(ctx)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return entUser, nil
}

func (r *userResolver) Images(ctx context.Context, obj *ent.User, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.ImageWhereInput, orderBy *ent.ImageOrder) (*ent.ImageConnection, error) {
	return r.controller.Image.List( //nolint:wrapcheck
		ctx,
		*obj,
		after,
		first,
		before,
		last,
		where,
		orderBy,
	)
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
