package schema

import (
	"StegoLSB/ent/mixin"
	"StegoLSB/pkg/const/globalid"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
)

// A ColorChannel represents the chosen RGB Channels.
type ColorChannel int

// RGB Channel.
const (
	RED ColorChannel = iota
	GREEN
	BLUE
	RED_GREEN
	RED_BLUE
	GREEN_BLUE
	RED_GREEN_BLUE
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
			Values(
				RED.String(),
				GREEN.String(),
				BLUE.String(),
				RED_GREEN.String(),
				RED_BLUE.String(),
				GREEN_BLUE.String(),
				RED_GREEN_BLUE.String(),
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
