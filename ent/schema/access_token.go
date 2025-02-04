package schema

import (
	"AltTube-Go/ent/mixins"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AccessToken schema definition.
type AccessToken struct {
	ent.Schema
}

func (AccessToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (AccessToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").Optional(),
		field.String("user_id").Optional(),
		field.Time("expiry").Optional(),
	}
}

func (AccessToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("access_tokens").Field("user_id").Unique(),
		edge.From("refresh_token", RefreshToken.Type).Ref("access_tokens").Unique(),
	}
}
