package schema

import (
	"AltTube-Go/ent/mixins"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RefreshToken schema definition.
type RefreshToken struct {
	ent.Schema
}

func (RefreshToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.UintIDMixin{},
	}
}

func (RefreshToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").Optional(),
		field.Time("expiry").Optional(),
		field.String("user_agent").Optional(),
		field.String("ip_address").Optional(),
	}
}

func (RefreshToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("refresh_tokens").Unique().Required(),
		edge.To("access_tokens", AccessToken.Type).Annotations(entsql.OnDelete("cascade")),
	}
}
