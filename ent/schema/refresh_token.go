package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RefreshToken schema definition.
type RefreshToken struct {
	ent.Schema
}

func (RefreshToken) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Time("created_at").Optional(),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional(),
		field.String("token").Optional(),
		field.Time("expiry").Optional(),
		field.String("user_agent").Optional(),
		field.String("ip_address").Optional(),
		field.String("user_id").Optional(),
	}
}

func (RefreshToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("refresh_tokens").Field("user_id").Unique(),
		edge.To("access_tokens", AccessToken.Type),
	}
}
