package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"StegoLSB/ent"
	"StegoLSB/graph/generated"
	"StegoLSB/pkg/entity/model"
	"StegoLSB/pkg/infrastructure/middleware"
	"StegoLSB/pkg/util"
	"context"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input generated.NewUser) (*generated.AuthUser, error) {
	entUser, err := r.controller.User.Create(ctx, input)

	if err != nil {
		return nil, err
	}

	return util.GenerateAuthUser(ctx, *entUser)
}

func (r *mutationResolver) Login(ctx context.Context, input generated.Login) (*generated.AuthUser, error) {
	entUser, _ := r.controller.User.Get(ctx, input.Username)

	if entUser == nil || !util.CheckPasswordHash(input.Password, entUser.Password) {
		return nil, model.NewNotFoundError(ctx, "username or password is incorrect", "user")
	}

	return util.GenerateAuthUser(ctx, *entUser)
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input generated.RefreshTokenInput) (*generated.AuthUser, error) {
	username, err := util.ParseToken(ctx, input.Token)

	if err != nil {
		return nil, err
	}

	entUser, err := r.controller.User.Get(ctx, username)

	if err != nil {
		return nil, err
	}

	return util.GenerateAuthUser(ctx, *entUser)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input generated.UpdateUser) (*ent.User, error) {
	entUser, err := middleware.ForContext(ctx)
	if err != nil {
		return nil, model.NewAuthorizationError(ctx, err.Error())
	}

	return r.controller.User.Update(ctx, *entUser, input)
}

func (r *queryResolver) Overview(ctx context.Context) (*ent.User, error) {
	entUser, err := middleware.ForContext(ctx)
	if err != nil {
		return nil, model.NewAuthorizationError(ctx, err.Error())
	}

	return entUser, nil
}

func (r *userResolver) Images(ctx context.Context, obj *ent.User, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.ImageWhereInput, orderBy *ent.ImageOrder) (*ent.ImageConnection, error) {
	return r.controller.Image.List(ctx, *obj, after, first, before, last, where, orderBy)
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
