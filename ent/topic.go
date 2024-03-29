// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/interngowhere/web-backend/ent/topic"
)

// Topic is the model entity for the Topic schema.
type Topic struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// ShortDescription holds the value of the "short_description" field.
	ShortDescription string `json:"short_description,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// ProfilePicURL holds the value of the "profile_pic_url" field.
	ProfilePicURL string `json:"profile_pic_url,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TopicQuery when eager-loading is set.
	Edges        TopicEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TopicEdges holds the relations/edges for other nodes in the graph.
type TopicEdges struct {
	// TopicThreads holds the value of the topic_threads edge.
	TopicThreads []*Thread `json:"topic_threads,omitempty"`
	// TopicModerators holds the value of the topic_moderators edge.
	TopicModerators []*User `json:"topic_moderators,omitempty"`
	// Moderators holds the value of the moderators edge.
	Moderators []*Moderator `json:"moderators,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// TopicThreadsOrErr returns the TopicThreads value or an error if the edge
// was not loaded in eager-loading.
func (e TopicEdges) TopicThreadsOrErr() ([]*Thread, error) {
	if e.loadedTypes[0] {
		return e.TopicThreads, nil
	}
	return nil, &NotLoadedError{edge: "topic_threads"}
}

// TopicModeratorsOrErr returns the TopicModerators value or an error if the edge
// was not loaded in eager-loading.
func (e TopicEdges) TopicModeratorsOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.TopicModerators, nil
	}
	return nil, &NotLoadedError{edge: "topic_moderators"}
}

// ModeratorsOrErr returns the Moderators value or an error if the edge
// was not loaded in eager-loading.
func (e TopicEdges) ModeratorsOrErr() ([]*Moderator, error) {
	if e.loadedTypes[2] {
		return e.Moderators, nil
	}
	return nil, &NotLoadedError{edge: "moderators"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Topic) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case topic.FieldID:
			values[i] = new(sql.NullInt64)
		case topic.FieldTitle, topic.FieldSlug, topic.FieldShortDescription, topic.FieldDescription, topic.FieldProfilePicURL:
			values[i] = new(sql.NullString)
		case topic.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Topic fields.
func (t *Topic) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case topic.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case topic.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				t.Title = value.String
			}
		case topic.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				t.Slug = value.String
			}
		case topic.FieldShortDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field short_description", values[i])
			} else if value.Valid {
				t.ShortDescription = value.String
			}
		case topic.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = value.String
			}
		case topic.FieldProfilePicURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field profile_pic_url", values[i])
			} else if value.Valid {
				t.ProfilePicURL = value.String
			}
		case topic.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Topic.
// This includes values selected through modifiers, order, etc.
func (t *Topic) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryTopicThreads queries the "topic_threads" edge of the Topic entity.
func (t *Topic) QueryTopicThreads() *ThreadQuery {
	return NewTopicClient(t.config).QueryTopicThreads(t)
}

// QueryTopicModerators queries the "topic_moderators" edge of the Topic entity.
func (t *Topic) QueryTopicModerators() *UserQuery {
	return NewTopicClient(t.config).QueryTopicModerators(t)
}

// QueryModerators queries the "moderators" edge of the Topic entity.
func (t *Topic) QueryModerators() *ModeratorQuery {
	return NewTopicClient(t.config).QueryModerators(t)
}

// Update returns a builder for updating this Topic.
// Note that you need to call Topic.Unwrap() before calling this method if this Topic
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Topic) Update() *TopicUpdateOne {
	return NewTopicClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Topic entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Topic) Unwrap() *Topic {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Topic is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Topic) String() string {
	var builder strings.Builder
	builder.WriteString("Topic(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("title=")
	builder.WriteString(t.Title)
	builder.WriteString(", ")
	builder.WriteString("slug=")
	builder.WriteString(t.Slug)
	builder.WriteString(", ")
	builder.WriteString("short_description=")
	builder.WriteString(t.ShortDescription)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(t.Description)
	builder.WriteString(", ")
	builder.WriteString("profile_pic_url=")
	builder.WriteString(t.ProfilePicURL)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Topics is a parsable slice of Topic.
type Topics []*Topic
