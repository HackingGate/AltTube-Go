package schema

import (
	"AltTube-Go/ent/mixins"

	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RefreshToken schema definition.
type RefreshToken struct {
	ent.Schema
}

func (RefreshToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixins.UintIDMixin{},
	}
}

func (RefreshToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").Optional(),
		field.Time("expiry").Optional(),
		field.String("user_agent").Optional(),
		field.String("ip_address").Optional(),
		field.UUID("user_id", uuid.UUID{}),
	}
}

func (RefreshToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("refresh_tokens").Field("user_id").Unique().Required(),
		edge.To("access_tokens", AccessToken.Type),
	}
}
