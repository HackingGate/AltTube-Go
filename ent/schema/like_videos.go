package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// LikeVideo schema definition.
type LikeVideo struct {
	ent.Schema
}

func (LikeVideo) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Time("created_at").Optional(),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional(),
		field.String("user_id").Optional(),
		field.String("video_id").Optional(),
	}
}

func (LikeVideo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("like_videos").Field("user_id").Unique(),
		edge.From("video", Video.Type).Ref("like_videos").Field("video_id").Unique(),
	}
}
