package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ThreadKudo holds the schema definition for the ThreadKudo entity.
type ThreadKudo struct {
	ent.Schema
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
            Field("user_id"),
        edge.To("thread", Thread.Type).
            Unique().
            Required().
			Immutable().
            Field("thread_id"),
    }
}
