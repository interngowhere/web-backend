package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CommentKudo holds the schema definition for the CommentKudo entity.
type CommentKudo struct {
	ent.Schema
}

// Annotations of the CommentKudo to generate composite
// primary key from user_id and comment_id.
func (CommentKudo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "comment_id"),
	}
}

// Fields of the CommentKudo.
func (CommentKudo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}).
			Immutable(),
		field.Int("comment_id").
			Immutable(),
	}
}

// Edges of the CommentKudo.
func (CommentKudo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required().
			Immutable().
			Field("user_id").
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("comment", Comment.Type).
			Unique().
			Required().
			Immutable().
			Field("comment_id").
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
