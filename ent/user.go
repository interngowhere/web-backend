// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/interngowhere/web-backend/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Hash holds the value of the "hash" field.
	Hash string `json:"-"`
	// EmailVerified holds the value of the "email_verified" field.
	EmailVerified bool `json:"email_verified,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// UserThreads holds the value of the user_threads edge.
	UserThreads []*Thread `json:"user_threads,omitempty"`
	// KudoedThreads holds the value of the kudoed_threads edge.
	KudoedThreads []*Thread `json:"kudoed_threads,omitempty"`
	// UserComments holds the value of the user_comments edge.
	UserComments []*Comment `json:"user_comments,omitempty"`
	// KudoedComments holds the value of the kudoed_comments edge.
	KudoedComments []*Comment `json:"kudoed_comments,omitempty"`
	// ModeratedTopics holds the value of the moderated_topics edge.
	ModeratedTopics []*Topic `json:"moderated_topics,omitempty"`
	// ThreadKudoes holds the value of the thread_kudoes edge.
	ThreadKudoes []*ThreadKudo `json:"thread_kudoes,omitempty"`
	// CommentKudoes holds the value of the comment_kudoes edge.
	CommentKudoes []*CommentKudo `json:"comment_kudoes,omitempty"`
	// Moderators holds the value of the moderators edge.
	Moderators []*Moderator `json:"moderators,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
}

// UserThreadsOrErr returns the UserThreads value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserThreadsOrErr() ([]*Thread, error) {
	if e.loadedTypes[0] {
		return e.UserThreads, nil
	}
	return nil, &NotLoadedError{edge: "user_threads"}
}

// KudoedThreadsOrErr returns the KudoedThreads value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) KudoedThreadsOrErr() ([]*Thread, error) {
	if e.loadedTypes[1] {
		return e.KudoedThreads, nil
	}
	return nil, &NotLoadedError{edge: "kudoed_threads"}
}

// UserCommentsOrErr returns the UserComments value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserCommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[2] {
		return e.UserComments, nil
	}
	return nil, &NotLoadedError{edge: "user_comments"}
}

// KudoedCommentsOrErr returns the KudoedComments value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) KudoedCommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[3] {
		return e.KudoedComments, nil
	}
	return nil, &NotLoadedError{edge: "kudoed_comments"}
}

// ModeratedTopicsOrErr returns the ModeratedTopics value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ModeratedTopicsOrErr() ([]*Topic, error) {
	if e.loadedTypes[4] {
		return e.ModeratedTopics, nil
	}
	return nil, &NotLoadedError{edge: "moderated_topics"}
}

// ThreadKudoesOrErr returns the ThreadKudoes value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ThreadKudoesOrErr() ([]*ThreadKudo, error) {
	if e.loadedTypes[5] {
		return e.ThreadKudoes, nil
	}
	return nil, &NotLoadedError{edge: "thread_kudoes"}
}

// CommentKudoesOrErr returns the CommentKudoes value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CommentKudoesOrErr() ([]*CommentKudo, error) {
	if e.loadedTypes[6] {
		return e.CommentKudoes, nil
	}
	return nil, &NotLoadedError{edge: "comment_kudoes"}
}

// ModeratorsOrErr returns the Moderators value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ModeratorsOrErr() ([]*Moderator, error) {
	if e.loadedTypes[7] {
		return e.Moderators, nil
	}
	return nil, &NotLoadedError{edge: "moderators"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldEmailVerified:
			values[i] = new(sql.NullBool)
		case user.FieldEmail, user.FieldUsername, user.FieldFirstName, user.FieldLastName, user.FieldHash:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case user.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				u.FirstName = value.String
			}
		case user.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				u.LastName = value.String
			}
		case user.FieldHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value.Valid {
				u.Hash = value.String
			}
		case user.FieldEmailVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field email_verified", values[i])
			} else if value.Valid {
				u.EmailVerified = value.Bool
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryUserThreads queries the "user_threads" edge of the User entity.
func (u *User) QueryUserThreads() *ThreadQuery {
	return NewUserClient(u.config).QueryUserThreads(u)
}

// QueryKudoedThreads queries the "kudoed_threads" edge of the User entity.
func (u *User) QueryKudoedThreads() *ThreadQuery {
	return NewUserClient(u.config).QueryKudoedThreads(u)
}

// QueryUserComments queries the "user_comments" edge of the User entity.
func (u *User) QueryUserComments() *CommentQuery {
	return NewUserClient(u.config).QueryUserComments(u)
}

// QueryKudoedComments queries the "kudoed_comments" edge of the User entity.
func (u *User) QueryKudoedComments() *CommentQuery {
	return NewUserClient(u.config).QueryKudoedComments(u)
}

// QueryModeratedTopics queries the "moderated_topics" edge of the User entity.
func (u *User) QueryModeratedTopics() *TopicQuery {
	return NewUserClient(u.config).QueryModeratedTopics(u)
}

// QueryThreadKudoes queries the "thread_kudoes" edge of the User entity.
func (u *User) QueryThreadKudoes() *ThreadKudoQuery {
	return NewUserClient(u.config).QueryThreadKudoes(u)
}

// QueryCommentKudoes queries the "comment_kudoes" edge of the User entity.
func (u *User) QueryCommentKudoes() *CommentKudoQuery {
	return NewUserClient(u.config).QueryCommentKudoes(u)
}

// QueryModerators queries the "moderators" edge of the User entity.
func (u *User) QueryModerators() *ModeratorQuery {
	return NewUserClient(u.config).QueryModerators(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(u.FirstName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(u.LastName)
	builder.WriteString(", ")
	builder.WriteString("hash=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("email_verified=")
	builder.WriteString(fmt.Sprintf("%v", u.EmailVerified))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
