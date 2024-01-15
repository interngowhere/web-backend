package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Tag holds the schema definition for the Tag entity.
type Moderator struct {
	ent.Schema
}

// Annotations of the Moderator to generate composite 
// primary key from moderator_id and topic_id.
func (Moderator) Annotations() []schema.Annotation {
    return []schema.Annotation{
        field.ID("moderator_id", "topic_id"),
    }
}

// Fields of the Tag.
func (Moderator) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("moderator_id", uuid.UUID{}).
			Immutable(),
		field.Int("topic_id").
			Immutable(),
	}
}

// Edges of the Moderator.
func (Moderator) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("moderator", User.Type).
            Unique().
            Required().
			Immutable().
            Field("moderator_id").
			Annotations(entsql.OnDelete(entsql.Cascade)),
        edge.To("topic", Topic.Type).
            Unique().
            Required().
			Immutable().
            Field("topic_id").
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
