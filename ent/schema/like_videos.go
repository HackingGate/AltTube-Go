package schema

import (
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// LikeVideo schema definition.
type LikeVideo struct {
	ent.Schema
}

func (LikeVideo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (LikeVideo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}),
		field.String("video_id"),
	}
}

func (LikeVideo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("like_videos").Field("user_id").Unique().Required(),
		edge.From("video", Video.Type).Ref("like_videos").Field("video_id").Unique().Required(),
	}
}
