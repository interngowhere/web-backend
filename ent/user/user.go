// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldSalt holds the string denoting the salt field in the database.
	FieldSalt = "salt"
	// FieldIsModerator holds the string denoting the is_moderator field in the database.
	FieldIsModerator = "is_moderator"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeUserThreads holds the string denoting the user_threads edge name in mutations.
	EdgeUserThreads = "user_threads"
	// EdgeKudoedThreads holds the string denoting the kudoed_threads edge name in mutations.
	EdgeKudoedThreads = "kudoed_threads"
	// EdgeUserComments holds the string denoting the user_comments edge name in mutations.
	EdgeUserComments = "user_comments"
	// EdgeKudoedComments holds the string denoting the kudoed_comments edge name in mutations.
	EdgeKudoedComments = "kudoed_comments"
	// EdgeThreadKudoes holds the string denoting the thread_kudoes edge name in mutations.
	EdgeThreadKudoes = "thread_kudoes"
	// EdgeCommentKudoes holds the string denoting the comment_kudoes edge name in mutations.
	EdgeCommentKudoes = "comment_kudoes"
	// Table holds the table name of the user in the database.
	Table = "users"
	// UserThreadsTable is the table that holds the user_threads relation/edge.
	UserThreadsTable = "threads"
	// UserThreadsInverseTable is the table name for the Thread entity.
	// It exists in this package in order to avoid circular dependency with the "thread" package.
	UserThreadsInverseTable = "threads"
	// UserThreadsColumn is the table column denoting the user_threads relation/edge.
	UserThreadsColumn = "created_by"
	// KudoedThreadsTable is the table that holds the kudoed_threads relation/edge. The primary key declared below.
	KudoedThreadsTable = "thread_kudos"
	// KudoedThreadsInverseTable is the table name for the Thread entity.
	// It exists in this package in order to avoid circular dependency with the "thread" package.
	KudoedThreadsInverseTable = "threads"
	// UserCommentsTable is the table that holds the user_comments relation/edge.
	UserCommentsTable = "comments"
	// UserCommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	UserCommentsInverseTable = "comments"
	// UserCommentsColumn is the table column denoting the user_comments relation/edge.
	UserCommentsColumn = "created_by"
	// KudoedCommentsTable is the table that holds the kudoed_comments relation/edge. The primary key declared below.
	KudoedCommentsTable = "comment_kudos"
	// KudoedCommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	KudoedCommentsInverseTable = "comments"
	// ThreadKudoesTable is the table that holds the thread_kudoes relation/edge.
	ThreadKudoesTable = "thread_kudos"
	// ThreadKudoesInverseTable is the table name for the ThreadKudo entity.
	// It exists in this package in order to avoid circular dependency with the "threadkudo" package.
	ThreadKudoesInverseTable = "thread_kudos"
	// ThreadKudoesColumn is the table column denoting the thread_kudoes relation/edge.
	ThreadKudoesColumn = "user_id"
	// CommentKudoesTable is the table that holds the comment_kudoes relation/edge.
	CommentKudoesTable = "comment_kudos"
	// CommentKudoesInverseTable is the table name for the CommentKudo entity.
	// It exists in this package in order to avoid circular dependency with the "commentkudo" package.
	CommentKudoesInverseTable = "comment_kudos"
	// CommentKudoesColumn is the table column denoting the comment_kudoes relation/edge.
	CommentKudoesColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldUsername,
	FieldFirstName,
	FieldLastName,
	FieldHash,
	FieldSalt,
	FieldIsModerator,
	FieldCreatedAt,
}

var (
	// KudoedThreadsPrimaryKey and KudoedThreadsColumn2 are the table columns denoting the
	// primary key for the kudoed_threads relation (M2M).
	KudoedThreadsPrimaryKey = []string{"user_id", "thread_id"}
	// KudoedCommentsPrimaryKey and KudoedCommentsColumn2 are the table columns denoting the
	// primary key for the kudoed_comments relation (M2M).
	KudoedCommentsPrimaryKey = []string{"user_id", "comment_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	FirstNameValidator func(string) error
	// LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	LastNameValidator func(string) error
	// SaltValidator is a validator for the "salt" field. It is called by the builders before save.
	SaltValidator func(string) error
	// DefaultIsModerator holds the default value on creation for the "is_moderator" field.
	DefaultIsModerator bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByFirstName orders the results by the first_name field.
func ByFirstName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstName, opts...).ToFunc()
}

// ByLastName orders the results by the last_name field.
func ByLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastName, opts...).ToFunc()
}

// ByHash orders the results by the hash field.
func ByHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHash, opts...).ToFunc()
}

// BySalt orders the results by the salt field.
func BySalt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSalt, opts...).ToFunc()
}

// ByIsModerator orders the results by the is_moderator field.
func ByIsModerator(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsModerator, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUserThreadsCount orders the results by user_threads count.
func ByUserThreadsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserThreadsStep(), opts...)
	}
}

// ByUserThreads orders the results by user_threads terms.
func ByUserThreads(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserThreadsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByKudoedThreadsCount orders the results by kudoed_threads count.
func ByKudoedThreadsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newKudoedThreadsStep(), opts...)
	}
}

// ByKudoedThreads orders the results by kudoed_threads terms.
func ByKudoedThreads(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newKudoedThreadsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserCommentsCount orders the results by user_comments count.
func ByUserCommentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserCommentsStep(), opts...)
	}
}

// ByUserComments orders the results by user_comments terms.
func ByUserComments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserCommentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByKudoedCommentsCount orders the results by kudoed_comments count.
func ByKudoedCommentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newKudoedCommentsStep(), opts...)
	}
}

// ByKudoedComments orders the results by kudoed_comments terms.
func ByKudoedComments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newKudoedCommentsStep(), append([]sql.OrderTerm{term}, terms...)...)
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
func newUserThreadsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserThreadsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserThreadsTable, UserThreadsColumn),
	)
}
func newKudoedThreadsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(KudoedThreadsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, KudoedThreadsTable, KudoedThreadsPrimaryKey...),
	)
}
func newUserCommentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserCommentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserCommentsTable, UserCommentsColumn),
	)
}
func newKudoedCommentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(KudoedCommentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, KudoedCommentsTable, KudoedCommentsPrimaryKey...),
	)
}
func newThreadKudoesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ThreadKudoesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, ThreadKudoesTable, ThreadKudoesColumn),
	)
}
func newCommentKudoesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CommentKudoesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, CommentKudoesTable, CommentKudoesColumn),
	)
}