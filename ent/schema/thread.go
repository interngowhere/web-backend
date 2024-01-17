package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Thread holds the schema definition for the Thread entity.
type Thread struct {
	ent.Schema
}

// Fields of the Thread.
func (Thread) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			MaxLen(255),
		field.String("slug").
			MaxLen(255),
		field.String("description").
			MaxLen(4096).
			Optional(),
		field.Time("modified_at").
			Optional(),
		field.UUID("created_by", uuid.UUID{}).
			Immutable(),
		field.Time("created_at").
			Default(time.Now()).
			Immutable(),
	}
}

// Edges of the Thread.
func (Thread) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("thread_comments", Comment.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("tags", Tag.Type),
		edge.From("topics", Topic.Type).
			Ref("topic_threads").
			Required().
			Unique(),
		edge.From("users", User.Type).
			Ref("user_threads").
			Field("created_by").
			Required().
			Unique().
			Immutable(),
		edge.From("kudoed_users", User.Type).
			Ref("kudoed_threads").
			Through("thread_kudoes", ThreadKudo.Type),
	}
}
