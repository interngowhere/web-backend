package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
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
		field.Bytes("hash").
			MaxLen(255).
			Sensitive().
			Unique(),
		field.Bool("email_verified").
			Default(false),
		field.Time("created_at").
			Default(time.Now()).
			Immutable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_threads", Thread.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("kudoed_threads", Thread.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Through("thread_kudoes", ThreadKudo.Type),
		edge.To("user_comments", Comment.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("kudoed_comments", Comment.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Through("comment_kudoes", CommentKudo.Type),
		edge.To("moderated_topics", Topic.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Through("moderators", Moderator.Type),
	}
}
