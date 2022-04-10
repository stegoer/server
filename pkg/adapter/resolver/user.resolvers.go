package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/stegoer/server/ent"
	"github.com/stegoer/server/graph/generated"
	"github.com/stegoer/server/pkg/cryptography"
	"github.com/stegoer/server/pkg/infrastructure/middleware"
	"github.com/stegoer/server/pkg/model"
	"github.com/stegoer/server/pkg/util"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input generated.NewUser) (*generated.CreateUserPayload, error) {
	entUser, err := r.controller.User.Create(ctx, input)
	if err != nil {
		r.logger.Errorw("create_user: failed to create",
			"error", err.Error(),
		)

		return nil, model.NewDBError(ctx, "failed to create user")
	}

	auth, err := util.GenerateAuth(ctx, *entUser)
	if err != nil {
		r.logger.Errorw("create_user: failed to generate auth",
			"error", err.Error(),
			"user_id", entUser.ID,
		)

		return nil, model.NewInternalServerError(
			ctx,
			"failed to generate authentication",
		)
	}

	r.logger.Debugw("create_user: success",
		"user_id", entUser.ID,
		"expires", auth.Expires,
	)

	return &generated.CreateUserPayload{
		User: entUser,
		Auth: auth,
	}, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input generated.UpdateUser) (*generated.UpdateUserPayload, error) {
	entUser, err := middleware.JwtForContext(ctx)
	if err != nil {
		r.logger.Infow("update_user: user not authenticated", "error", err.Error())

		return nil, model.NewAuthorizationError(ctx, "user is not authenticated")
	}

	entUser, err = r.controller.User.Update(ctx, *entUser, input)
	if err != nil {
		r.logger.Infow("update_user: failed to update",
			"error", err.Error(),
			"user_id", entUser.ID,
		)

		return nil, model.NewDBError(ctx, "failed to update user")
	}

	r.logger.Debugw("update_user: success",
		"user_id", entUser.ID,
	)

	return &generated.UpdateUserPayload{User: entUser}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input generated.Login) (*generated.LoginPayload, error) {
	entUser, err := r.controller.User.GetByEmail(ctx, input.Email)
	if err != nil {
		r.logger.Infow("login: user not found", "error", err.Error())
	}

	if entUser == nil || !cryptography.CheckPasswordHash(
		input.Password,
		entUser.Password,
	) {
		r.logger.Infow("login: invalid credentials")

		return nil, model.NewNotFoundError(ctx, "email or password is incorrect")
	}

	entUser, err = r.controller.User.SetLoggedIn(ctx, *entUser)
	if err != nil {
		r.logger.Errorw("login: failed to set last login date",
			"error", err.Error(),
			"user_id", entUser.ID,
		)

		return nil, model.NewDBError(ctx, "failed to set last login date")
	}

	auth, err := util.GenerateAuth(ctx, *entUser)
	if err != nil {
		r.logger.Errorw("login: failed to generate auth",
			"error", err.Error(),
			"user_id", entUser.ID,
		)

		return nil, model.NewInternalServerError(
			ctx,
			"failed to generate authentication",
		)
	}

	r.logger.Debugw("login: success",
		"user_id", entUser.ID,
		"expires", auth.Expires,
	)

	return &generated.LoginPayload{
		User: entUser,
		Auth: auth,
	}, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input generated.RefreshTokenInput) (*generated.RefreshTokenPayload, error) {
	userID, err := util.ParseToken(input.Token)
	if err != nil {
		r.logger.Errorw("refresh_token: failed to parse token",
			"error", err.Error(),
			"token", input.Token,
		)

		return nil, model.NewAuthorizationError(ctx, "invalid token to refresh")
	}

	entUser, err := r.controller.User.GetByID(ctx, userID)
	if err != nil {
		r.logger.Errorw("refresh_token: user not found",
			"error", err.Error(),
			"user_id", userID,
			"token", input.Token,
		)

		return nil, model.NewNotFoundError(ctx, "user not found")
	}

	auth, err := util.GenerateAuth(ctx, *entUser)
	if err != nil {
		r.logger.Errorw("refresh_token: failed to generate auth",
			"error", err.Error(),
			"user_id", entUser.ID,
		)

		return nil, model.NewInternalServerError(
			ctx,
			"failed to generate authentication",
		)
	}

	r.logger.Debugw("refresh_token: success",
		"user_id", entUser.ID,
		"expires", auth.Expires,
	)

	return &generated.RefreshTokenPayload{
		User: entUser,
		Auth: auth,
	}, nil
}

func (r *queryResolver) Overview(ctx context.Context) (*generated.OverviewPayload, error) {
	entUser, err := middleware.JwtForContext(ctx)
	if err != nil {
		r.logger.Infow("overview: user not authenticated", "error", err.Error())

		return nil, model.NewAuthorizationError(ctx, "user is not authenticated")
	}

	r.logger.Debugw("overview: success", "user", entUser)

	return &generated.OverviewPayload{User: entUser}, nil
}

func (r *userResolver) Username(ctx context.Context, obj *ent.User) (string, error) {
	return obj.Name, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
