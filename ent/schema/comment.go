package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("parent_id").
			Default("").
			MaxLen(255),
		field.String("content").
			MaxLen(4096),
		field.Time("modified_at").
			Optional(),
		field.UUID("created_by", uuid.UUID{}).
			Immutable(),
		field.Time("created_at").
			Default(time.Now()).
			Immutable(),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("threads", Thread.Type).
			Ref("thread_comments").
			Required().
			Unique(),
		edge.From("comment_authors", User.Type).
			Ref("user_comments").
			Field("created_by").
			Required().
			Unique().
			Immutable(),
		edge.From("kudoed_users", User.Type).
            Ref("kudoed_comments").
            Through("comment_kudoes", CommentKudo.Type),
    }
}
