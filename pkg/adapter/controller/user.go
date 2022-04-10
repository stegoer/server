package controller

import (
	"context"

	"github.com/stegoer/server/ent"
	"github.com/stegoer/server/ent/schema/ulid"
	"github.com/stegoer/server/gqlgen"
)

// User Controller interface.
type User interface {
	GetByID(
		ctx context.Context,
		id ulid.ID,
	) (*ent.User, error)
	GetByEmail(
		ctx context.Context,
		email string,
	) (*ent.User, error)
	Create(
		ctx context.Context,
		input gqlgen.NewUser,
	) (*ent.User, error)
	Update(
		ctx context.Context,
		entUser ent.User,
		input gqlgen.UpdateUser,
	) (*ent.User, error)
	SetLoggedIn(
		ctx context.Context,
		entUser ent.User,
	) (*ent.User, error)
}
