package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// UuidIdMixin defines a mixin for using a UUID as the primary key.
type UuidIdMixin struct {
	mixin.Schema
}

// Fields of the UuidIdMixin. These fields will be added to any schema that embeds this mixin.
func (UuidIdMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			// For PostgreSQL, you can ensure the column is of type UUID.
			SchemaType(map[string]string{
				"postgres": "uuid",
			}).
			Immutable().
			Unique(),
	}
}
