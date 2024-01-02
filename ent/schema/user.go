package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
            Default(uuid.New).
			Immutable(),
		field.String("email").
			MaxLen(255),
		field.String("username").
			MaxLen(255).
			Unique(),
		field.String("first_name").
			MaxLen(255).
			Optional(),
		field.String("last_name").
			MaxLen(255).
			Optional(),
		field.String("hash").
			Optional().
			Sensitive().
			Unique(),
		field.String("salt").
			MaxLen(255).
			Sensitive().
			Optional(),
		field.Time("created_at").
			Default(time.Now()).
			Immutable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_threads", Thread.Type),
        edge.To("kudoed_threads", Thread.Type).
			Through("thread_kudoes", ThreadKudo.Type),
		edge.To("user_comments", Comment.Type),
		edge.To("kudoed_comments", Comment.Type).
			Through("comment_kudoes", CommentKudo.Type),
		edge.To("moderated_topics", Topic.Type).
			Through("moderators", Moderator.Type),
    }
}
