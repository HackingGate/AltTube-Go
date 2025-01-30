package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AccessToken schema definition.
type AccessToken struct {
	ent.Schema
}

func (AccessToken) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Time("created_at").Optional(),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional(),
		field.String("token").Optional(),
		field.String("user_id").Optional(),
		field.Time("expiry").Optional(),
		field.Int64("refresh_token_id").Optional(),
	}
}

func (AccessToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("access_tokens").Field("user_id").Unique(),
		edge.From("refresh_token", RefreshToken.Type).Ref("access_tokens").Field("refresh_token_id").Unique(),
	}
}
