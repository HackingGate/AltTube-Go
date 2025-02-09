package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// UintIDMixin defines a mixin for using an unsigned integer as the primary key.
type UintIDMixin struct {
	mixin.Schema
}

// Fields of the UintIDMixin. These fields will be added to any schema that embeds this mixin.
func (UintIDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").
			Immutable().
			Unique(),
	}
}
