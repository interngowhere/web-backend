package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Topic holds the schema definition for the Topic entity.
type Topic struct {
	ent.Schema
}

// Fields of the Topic.
func (Topic) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			MaxLen(255).
			Unique(),
		field.String("slug").
			MaxLen(255).
			Unique(),
		field.String("short_description").
			MaxLen(255),
		field.String("description").
			MaxLen(4096).
			Optional(),
		field.String("profile_pic_url").
			Optional(),
		field.Time("created_at").
			Default(time.Now()).
			Immutable(),
	}
}

// Edges of the Topic.
func (Topic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("topic_threads", Thread.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From("topic_moderators", User.Type).
			Ref("moderated_topics").
			Through("moderators", Moderator.Type),
	}
}
