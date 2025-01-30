package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User schema definition.
type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("email").Unique().Optional(),
		field.String("password").NotEmpty(),
		field.Time("created_at").Optional(),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("access_tokens", AccessToken.Type),
		edge.To("like_videos", LikeVideo.Type),
		edge.To("refresh_tokens", RefreshToken.Type),
	}
}
