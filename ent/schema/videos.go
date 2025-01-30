package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Video schema definition.
type Video struct {
	ent.Schema
}

func (Video) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("title").NotEmpty(),
		field.String("description").NotEmpty(),
		field.Time("upload_date"),
		field.String("uploader").NotEmpty(),
		field.String("uploader_url").NotEmpty(),
		field.Time("created_at").Optional(),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional(),
		field.String("thumbnail_url").NotEmpty(),
	}
}

func (Video) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("like_videos", LikeVideo.Type),
	}
}
