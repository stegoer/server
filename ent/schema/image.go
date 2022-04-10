package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"

	"github.com/stegoer/server/ent/mixin"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// ImageMixin defines Fields.
type ImageMixin struct {
	entMixin.Schema
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.String("file_name").
			Sensitive().
			NotEmpty().
			Annotations(entgql.OrderField("FILE_NAME")),
		field.String("content").
			Sensitive().
			NotEmpty().
			Annotations(entgql.OrderField("CONTENT")),
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
		mixin.NewUlid("IMG"),
		ImageMixin{}, //nolint:exhaustivestruct
		mixin.NewDatetime(),
	}
}
