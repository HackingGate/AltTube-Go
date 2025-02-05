package mixins

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// BaseMixin defines common fields for all entities.
type BaseMixin struct {
	mixin.Schema
}

// Fields of the BaseMixin. These fields will be added to any schema that embeds this mixin.
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		// Automatically set when the record is created.
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		// Automatically set when the record is updated.
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		// Optional field for soft deletion.
		field.Time("deleted_at").
			Optional().
			Nillable(),
	}
}
