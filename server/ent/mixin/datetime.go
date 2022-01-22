package mixin

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// NewDatetime creates a Mixin that includes created_at and updated_at.
func NewDatetime() *DatetimeMixin {
	return &DatetimeMixin{} //nolint:exhaustivestruct
}

// DatetimeMixin defines an ent Mixin.
type DatetimeMixin struct {
	mixin.Schema
}

// Fields provides the created_at and updated_at field.
func (m DatetimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(
				entgql.OrderField("UPDATED_AT"),
			),
	}
}
