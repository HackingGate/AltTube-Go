package schema

import (
	"entgo.io/ent/schema/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Video schema definition.
type Video struct {
	ent.Schema
}

func (Video) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Video) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().NotEmpty(),
		field.String("title").NotEmpty(),
		field.String("description").NotEmpty(),
		field.Time("upload_date"),
		field.String("uploader").NotEmpty(),
		field.String("uploader_url").NotEmpty(),
		field.String("thumbnail_url").NotEmpty(),
	}
}

func (Video) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("like_videos", LikeVideo.Type),
	}
}
