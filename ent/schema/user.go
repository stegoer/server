package schema

import (
	"net/mail"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"

	"github.com/stegoer/server/ent/mixin"
	"github.com/stegoer/server/pkg/const/globalid"
)

const nameMinLen = 5

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// UserMixin defines Fields.
type UserMixin struct {
	entMixin.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			MinLen(nameMinLen),
		field.String("email").
			Unique().
			Validate(func(email string) error {
				_, err := mail.ParseAddress(email)

				return err //nolint:wrapcheck
			}),
		field.String("password").
			Sensitive(),
		field.Time("last_login").
			Default(time.Now).
			Annotations(
				entgql.OrderField("LAST_LOGIN"),
			),
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
