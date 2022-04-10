package controller

import (
	"context"

	"github.com/stegoer/server/ent"
	"github.com/stegoer/server/ent/schema/ulid"
)

// Image Controller interface.
type Image interface {
	Get(
		ctx context.Context,
		entUser ent.User,
		id *ulid.ID,
	) (*ent.Image, error)
	List(
		ctx context.Context,
		entUser ent.User,
		after *ent.Cursor,
		first *int,
		before *ent.Cursor,
		last *int,
		where *ent.ImageWhereInput,
		orderBy *ent.ImageOrder,
	) (*ent.ImageConnection, error)
	Create(
		ctx context.Context,
		entUser ent.User,
		filename string,
		content string,
	) (*ent.Image, error)
	Count(
		ctx context.Context,
		entUser ent.User,
	) (int, error)
}
