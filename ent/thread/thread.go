// Code generated by ent, DO NOT EDIT.

package thread

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the thread type in the database.
	Label = "thread"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldModifiedAt holds the string denoting the modified_at field in the database.
	FieldModifiedAt = "modified_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeThreadComments holds the string denoting the thread_comments edge name in mutations.
	EdgeThreadComments = "thread_comments"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeTopics holds the string denoting the topics edge name in mutations.
	EdgeTopics = "topics"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeKudoedUsers holds the string denoting the kudoed_users edge name in mutations.
	EdgeKudoedUsers = "kudoed_users"
	// EdgeThreadKudoes holds the string denoting the thread_kudoes edge name in mutations.
	EdgeThreadKudoes = "thread_kudoes"
	// Table holds the table name of the thread in the database.
	Table = "threads"
	// ThreadCommentsTable is the table that holds the thread_comments relation/edge.
	ThreadCommentsTable = "comments"
	// ThreadCommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	ThreadCommentsInverseTable = "comments"
	// ThreadCommentsColumn is the table column denoting the thread_comments relation/edge.
	ThreadCommentsColumn = "thread_thread_comments"
	// TagsTable is the table that holds the tags relation/edge. The primary key declared below.
	TagsTable = "thread_tags"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// TopicsTable is the table that holds the topics relation/edge.
	TopicsTable = "threads"
	// TopicsInverseTable is the table name for the Topic entity.
	// It exists in this package in order to avoid circular dependency with the "topic" package.
	TopicsInverseTable = "topics"
	// TopicsColumn is the table column denoting the topics relation/edge.
	TopicsColumn = "topic_topic_threads"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "threads"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "created_by"
	// KudoedUsersTable is the table that holds the kudoed_users relation/edge. The primary key declared below.
	KudoedUsersTable = "thread_kudos"
	// KudoedUsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	KudoedUsersInverseTable = "users"
	// ThreadKudoesTable is the table that holds the thread_kudoes relation/edge.
	ThreadKudoesTable = "thread_kudos"
	// ThreadKudoesInverseTable is the table name for the ThreadKudo entity.
	// It exists in this package in order to avoid circular dependency with the "threadkudo" package.
	ThreadKudoesInverseTable = "thread_kudos"
	// ThreadKudoesColumn is the table column denoting the thread_kudoes relation/edge.
	ThreadKudoesColumn = "thread_id"
)

// Columns holds all SQL columns for thread fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldSlug,
	FieldDescription,
	FieldModifiedAt,
	FieldCreatedBy,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "threads"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"topic_topic_threads",
}

var (
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"thread_id", "tag_id"}
	// KudoedUsersPrimaryKey and KudoedUsersColumn2 are the table columns denoting the
	// primary key for the kudoed_users relation (M2M).
	KudoedUsersPrimaryKey = []string{"user_id", "thread_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	SlugValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
)

// OrderOption defines the ordering options for the Thread queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// BySlug orders the results by the slug field.
func BySlug(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSlug, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByModifiedAt orders the results by the modified_at field.
func ByModifiedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModifiedAt, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByThreadCommentsCount orders the results by thread_comments count.
func ByThreadCommentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newThreadCommentsStep(), opts...)
	}
}

// ByThreadComments orders the results by thread_comments terms.
func ByThreadComments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newThreadCommentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTagsCount orders the results by tags count.
func ByTagsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTagsStep(), opts...)
	}
}

// ByTags orders the results by tags terms.
func ByTags(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTagsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTopicsField orders the results by topics field.
func ByTopicsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTopicsStep(), sql.OrderByField(field, opts...))
	}
}

// ByUsersField orders the results by users field.
func ByUsersField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), sql.OrderByField(field, opts...))
	}
}

// ByKudoedUsersCount orders the results by kudoed_users count.
func ByKudoedUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newKudoedUsersStep(), opts...)
	}
}

// ByKudoedUsers orders the results by kudoed_users terms.
func ByKudoedUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newKudoedUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByThreadKudoesCount orders the results by thread_kudoes count.
func ByThreadKudoesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newThreadKudoesStep(), opts...)
	}
}

// ByThreadKudoes orders the results by thread_kudoes terms.
func ByThreadKudoes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newThreadKudoesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newThreadCommentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ThreadCommentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ThreadCommentsTable, ThreadCommentsColumn),
	)
}
func newTagsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TagsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, TagsTable, TagsPrimaryKey...),
	)
}
func newTopicsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TopicsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TopicsTable, TopicsColumn),
	)
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
	)
}
func newKudoedUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(KudoedUsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, KudoedUsersTable, KudoedUsersPrimaryKey...),
	)
}
func newThreadKudoesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ThreadKudoesInverseTable, ThreadKudoesColumn),
		sqlgraph.Edge(sqlgraph.O2M, true, ThreadKudoesTable, ThreadKudoesColumn),
	)
}
