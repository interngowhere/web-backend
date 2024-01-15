package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ThreadKudo holds the schema definition for the ThreadKudo entity.
type ThreadKudo struct {
	ent.Schema
}

// Annotations of the ThreadKudo to generate composite 
// primary key from user_id and thread_id.
func (ThreadKudo) Annotations() []schema.Annotation {
    return []schema.Annotation{
        field.ID("user_id", "thread_id"),
    }
}

// Fields of the ThreadKudo.
func (ThreadKudo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}).
			Immutable(),
		field.Int("thread_id").
			Immutable(),
	}
}

// Edges of the ThreadKudo.
func (ThreadKudo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
            Unique().
            Required().
			Immutable().
            Field("user_id").
			Annotations(entsql.OnDelete(entsql.Cascade)),
        edge.To("thread", Thread.Type).
            Unique().
            Required().
			Immutable().
            Field("thread_id").
			Annotations(entsql.OnDelete(entsql.Cascade)),
    }
}
