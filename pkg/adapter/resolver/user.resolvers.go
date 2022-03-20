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

func (r *mutationResolver) CreateUser(ctx context.Context, input generated.NewUser) (*generated.CreateUserPayload, error) {
	entUser, err := r.controller.User.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	auth, err := util.GenerateAuth(ctx, *entUser)
	if err != nil {
		return nil, err
	}

	return &generated.CreateUserPayload{
		User: entUser,
		Auth: auth,
	}, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input generated.UpdateUser) (*generated.UpdateUserPayload, error) {
	entUser, err := middleware.JwtForContext(ctx)
	if err != nil {
		return nil, err
	}

	entUser, err = r.controller.User.Update(ctx, *entUser, input)
	if err != nil {
		return nil, err
	}

	return &generated.UpdateUserPayload{User: entUser}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input generated.Login) (*generated.LoginPayload, error) {
	entUser, _ := r.controller.User.GetByEmail(ctx, input.Email)

	if entUser == nil || !util.CheckPasswordHash(
		input.Password,
		entUser.Password,
	) {
		return nil, model.NewNotFoundError(ctx, "email or password is incorrect")
	}

	entUser, err := r.controller.User.SetLoggedIn(ctx, *entUser)
	if err != nil {
		return nil, err
	}

	auth, err := util.GenerateAuth(ctx, *entUser)
	if err != nil {
		return nil, err
	}

	return &generated.LoginPayload{
		User: entUser,
		Auth: auth,
	}, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input generated.RefreshTokenInput) (*generated.RefreshTokenPayload, error) {
	userID, err := util.ParseToken(ctx, input.Token)
	if err != nil {
		return nil, err
	}

	entUser, err := r.controller.User.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	auth, err := util.GenerateAuth(ctx, *entUser)
	if err != nil {
		return nil, err
	}

	return &generated.RefreshTokenPayload{
		User: entUser,
		Auth: auth,
	}, nil
}

func (r *queryResolver) Overview(ctx context.Context) (*generated.OverviewPayload, error) {
	entUser, err := middleware.JwtForContext(ctx)
	if err != nil {
		return &generated.OverviewPayload{User: nil}, err
	}

	return &generated.OverviewPayload{User: entUser}, nil
}

func (r *userResolver) Username(ctx context.Context, obj *ent.User) (string, error) {
	return obj.Name, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
