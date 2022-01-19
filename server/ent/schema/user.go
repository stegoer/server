package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
	"stegoer/ent/mixin"
	"stegoer/pkg/const/globalid"
)

const nameMinLen = 5

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// UserMixin defines Fields
type UserMixin struct {
	entMixin.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			MinLen(nameMinLen),
		field.String("password").
			Sensitive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("images", Image.Type).
			StorageKey(edge.Column("user_id")),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().User.Prefix),
		UserMixin{}, //nolint:exhaustivestruct
		mixin.NewDatetime(),
	}
}
