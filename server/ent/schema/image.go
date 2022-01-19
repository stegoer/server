package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"

	"stegoer/ent/mixin"
	"stegoer/pkg/const/globalid"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// ImageMixin defines Fields
type ImageMixin struct {
	entMixin.Schema
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("channel").
			NamedValues(
				"Red", "RED",
				"Green", "GREEN",
				"Blue", "BLUE",
				"RedGreen", "RED_GREEN",
				"RedBlue", "RED_BLUE",
				"GreenBlue", "GREEN_BLUE",
				"RedGreenBlue", "RED_GREEN_BLUE",
			).Annotations(
			entgql.OrderField("CHANNEL"),
		),
	}
}

// Edges of the Image.
func (Image) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("images").
			Unique(),
	}
}

// Mixin of the Image.
func (Image) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().Image.Prefix),
		ImageMixin{}, //nolint:exhaustivestruct
		mixin.NewDatetime(),
	}
}
