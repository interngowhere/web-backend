package schema

import (
	"time"

	"entgo.io/ent"
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
			MaxLen(255),
		field.String("short_title").
			MaxLen(255),
		field.String("description").
			MaxLen(4096).
			Optional(),
		field.String("created_by").
			MaxLen(255),
		field.Time("created_at").
			Default(time.Now()).
			Immutable(),
	}
}

// Edges of the Topic.
func (Topic) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("topic_threads", Thread.Type),
		edge.From("topic_moderators", User.Type).
            Ref("moderated_topics").
            Through("moderators", Moderator.Type),
    }
}
