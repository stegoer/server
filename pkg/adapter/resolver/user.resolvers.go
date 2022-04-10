package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/stegoer/server/ent"
	"github.com/stegoer/server/gqlgen"
	"github.com/stegoer/server/pkg/cryptography"
	"github.com/stegoer/server/pkg/infrastructure/middleware"
	"github.com/stegoer/server/pkg/util"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input gqlgen.NewUser) (*gqlgen.CreateUserPayload, error) {
	entUser, err := r.controller.User.Create(ctx, input)
	if err != nil {
		r.logger.Errorw("create_user: failed to create",
			"error", err.Error(),
		)

		return nil, util.NewDBError(ctx, "failed to create user")
	}

	token, exp, err := util.GenerateToken(ctx, *entUser)
	if err != nil {
		r.logger.Errorw("create_user: failed to generate token",
			"error", err.Error(),
			"user_id", entUser.ID,
		)

		return nil, util.NewInternalServerError(
			ctx,
			"failed to generate token",
		)
	}

	r.logger.Debugw("create_user: success",
		"user_id", entUser.ID,
		"expires", exp,
	)

	return &gqlgen.CreateUserPayload{
		User: entUser,
		Auth: &gqlgen.Auth{
			Token:   token,
			Expires: *exp,
		},
	}, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input gqlgen.UpdateUser) (*gqlgen.UpdateUserPayload, error) {
	entUser, err := middleware.JwtForContext(ctx)
	if err != nil {
		r.logger.Infow("update_user: user not authenticated", "error", err.Error())

		return nil, util.NewAuthorizationError(ctx, "user is not authenticated")
	}

	entUser, err = r.controller.User.Update(ctx, *entUser, input)
	if err != nil {
		r.logger.Infow("update_user: failed to update",
			"error", err.Error(),
			"user_id", entUser.ID,
		)

		return nil, util.NewDBError(ctx, "failed to update user")
	}

	r.logger.Debugw("update_user: success",
		"user_id", entUser.ID,
	)

	return &gqlgen.UpdateUserPayload{User: entUser}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input gqlgen.Login) (*gqlgen.LoginPayload, error) {
	entUser, err := r.controller.User.GetByEmail(ctx, input.Email)
	if err != nil {
		r.logger.Infow("login: user not found", "error", err.Error())
	}

	if entUser == nil || !cryptography.CheckPasswordHash(
		input.Password,
		entUser.Password,
	) {
		r.logger.Infow("login: invalid credentials")

		return nil, util.NewNotFoundError(ctx, "email or password is incorrect")
	}

	entUser, err = r.controller.User.SetLoggedIn(ctx, *entUser)
	if err != nil {
		r.logger.Errorw("login: failed to set last login date",
			"error", err.Error(),
			"user_id", entUser.ID,
		)

		return nil, util.NewDBError(ctx, "failed to set last login date")
	}

	token, exp, err := util.GenerateToken(ctx, *entUser)
	if err != nil {
		r.logger.Errorw("login: failed to generate token",
			"error", err.Error(),
			"user_id", entUser.ID,
		)

		return nil, util.NewInternalServerError(
			ctx,
			"failed to generate token",
		)
	}

	r.logger.Debugw("login: success",
		"user_id", entUser.ID,
		"expires", exp,
	)

	return &gqlgen.LoginPayload{
		User: entUser,
		Auth: &gqlgen.Auth{
			Token:   token,
			Expires: *exp,
		},
	}, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input gqlgen.RefreshTokenInput) (*gqlgen.RefreshTokenPayload, error) {
	userID, err := util.ParseToken(input.Token)
	if err != nil {
		r.logger.Errorw("refresh_token: failed to parse token",
			"error", err.Error(),
			"token", input.Token,
		)

		return nil, util.NewAuthorizationError(ctx, "invalid token to refresh")
	}

	entUser, err := r.controller.User.GetByID(ctx, userID)
	if err != nil {
		r.logger.Errorw("refresh_token: user not found",
			"error", err.Error(),
			"user_id", userID,
			"token", input.Token,
		)

		return nil, util.NewNotFoundError(ctx, "user not found")
	}

	token, exp, err := util.GenerateToken(ctx, *entUser)
	if err != nil {
		r.logger.Errorw("refresh_token: failed to generate token",
			"error", err.Error(),
			"user_id", entUser.ID,
		)

		return nil, util.NewInternalServerError(
			ctx,
			"failed to generate token",
		)
	}

	r.logger.Debugw("refresh_token: success",
		"user_id", entUser.ID,
		"expires", exp,
	)

	return &gqlgen.RefreshTokenPayload{
		User: entUser,
		Auth: &gqlgen.Auth{
			Token:   token,
			Expires: *exp,
		},
	}, nil
}

func (r *queryResolver) Overview(ctx context.Context) (*gqlgen.OverviewPayload, error) {
	entUser, err := middleware.JwtForContext(ctx)
	if err != nil {
		r.logger.Infow("overview: user not authenticated", "error", err.Error())

		return nil, util.NewAuthorizationError(ctx, "user is not authenticated")
	}

	r.logger.Debugw("overview: success", "user", entUser)

	return &gqlgen.OverviewPayload{User: entUser}, nil
}

func (r *userResolver) Username(ctx context.Context, obj *ent.User) (string, error) {
	return obj.Name, nil
}

// User returns gqlgen.UserResolver implementation.
func (r *Resolver) User() gqlgen.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
