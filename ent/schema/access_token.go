package schema

import (
	"github.com/hackinggate/alttube-go/ent/mixins"

	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"

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
		mixin.Time{},
		mixins.UintIDMixin{},
	}
}

func (AccessToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").Optional(),
		field.UUID("user_id", uuid.UUID{}),
		field.Uint("refresh_token_id"),
		field.Time("expiry").Optional(),
	}
}

func (AccessToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("access_tokens").Field("user_id").Unique().Required(),
		edge.From("refresh_token", RefreshToken.Type).Ref("access_tokens").Field("refresh_token_id").Unique().Required(),
	}
}
