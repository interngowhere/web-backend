// Code generated by ent, DO NOT EDIT.

package comment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the comment type in the database.
	Label = "comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldModifiedAt holds the string denoting the modified_at field in the database.
	FieldModifiedAt = "modified_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeThreads holds the string denoting the threads edge name in mutations.
	EdgeThreads = "threads"
	// EdgeCommentAuthors holds the string denoting the comment_authors edge name in mutations.
	EdgeCommentAuthors = "comment_authors"
	// EdgeKudoedUsers holds the string denoting the kudoed_users edge name in mutations.
	EdgeKudoedUsers = "kudoed_users"
	// EdgeCommentKudoes holds the string denoting the comment_kudoes edge name in mutations.
	EdgeCommentKudoes = "comment_kudoes"
	// Table holds the table name of the comment in the database.
	Table = "comments"
	// ThreadsTable is the table that holds the threads relation/edge.
	ThreadsTable = "comments"
	// ThreadsInverseTable is the table name for the Thread entity.
	// It exists in this package in order to avoid circular dependency with the "thread" package.
	ThreadsInverseTable = "threads"
	// ThreadsColumn is the table column denoting the threads relation/edge.
	ThreadsColumn = "thread_thread_comments"
	// CommentAuthorsTable is the table that holds the comment_authors relation/edge.
	CommentAuthorsTable = "comments"
	// CommentAuthorsInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CommentAuthorsInverseTable = "users"
	// CommentAuthorsColumn is the table column denoting the comment_authors relation/edge.
	CommentAuthorsColumn = "created_by"
	// KudoedUsersTable is the table that holds the kudoed_users relation/edge. The primary key declared below.
	KudoedUsersTable = "comment_kudos"
	// KudoedUsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	KudoedUsersInverseTable = "users"
	// CommentKudoesTable is the table that holds the comment_kudoes relation/edge.
	CommentKudoesTable = "comment_kudos"
	// CommentKudoesInverseTable is the table name for the CommentKudo entity.
	// It exists in this package in order to avoid circular dependency with the "commentkudo" package.
	CommentKudoesInverseTable = "comment_kudos"
	// CommentKudoesColumn is the table column denoting the comment_kudoes relation/edge.
	CommentKudoesColumn = "comment_id"
)

// Columns holds all SQL columns for comment fields.
var Columns = []string{
	FieldID,
	FieldParentID,
	FieldContent,
	FieldModifiedAt,
	FieldCreatedBy,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "comments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"thread_thread_comments",
}

var (
	// KudoedUsersPrimaryKey and KudoedUsersColumn2 are the table columns denoting the
	// primary key for the kudoed_users relation (M2M).
	KudoedUsersPrimaryKey = []string{"user_id", "comment_id"}
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
	// DefaultParentID holds the default value on creation for the "parent_id" field.
	DefaultParentID string
	// ParentIDValidator is a validator for the "parent_id" field. It is called by the builders before save.
	ParentIDValidator func(string) error
	// ContentValidator is a validator for the "content" field. It is called by the builders before save.
	ContentValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
)

// OrderOption defines the ordering options for the Comment queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByParentID orders the results by the parent_id field.
func ByParentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentID, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
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

// ByThreadsField orders the results by threads field.
func ByThreadsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newThreadsStep(), sql.OrderByField(field, opts...))
	}
}

// ByCommentAuthorsField orders the results by comment_authors field.
func ByCommentAuthorsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCommentAuthorsStep(), sql.OrderByField(field, opts...))
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

// ByCommentKudoesCount orders the results by comment_kudoes count.
func ByCommentKudoesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCommentKudoesStep(), opts...)
	}
}

// ByCommentKudoes orders the results by comment_kudoes terms.
func ByCommentKudoes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCommentKudoesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newThreadsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ThreadsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ThreadsTable, ThreadsColumn),
	)
}
func newCommentAuthorsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CommentAuthorsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CommentAuthorsTable, CommentAuthorsColumn),
	)
}
func newKudoedUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(KudoedUsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, KudoedUsersTable, KudoedUsersPrimaryKey...),
	)
}
func newCommentKudoesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CommentKudoesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, CommentKudoesTable, CommentKudoesColumn),
	)
}
