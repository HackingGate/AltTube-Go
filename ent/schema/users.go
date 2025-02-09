package schema

import (
	"github.com/hackinggate/alttube-go/ent/mixins"

	"entgo.io/ent/schema/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User schema definition.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		mixins.UuidIdMixin{},
	}
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique().NotEmpty(),
		field.String("password").NotEmpty(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("access_tokens", AccessToken.Type),
		edge.To("like_videos", LikeVideo.Type),
		edge.To("refresh_tokens", RefreshToken.Type),
	}
}
