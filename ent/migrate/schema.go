// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "parent_id", Type: field.TypeString, Size: 255, Default: ""},
		{Name: "content", Type: field.TypeString, Size: 4096},
		{Name: "modified_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "thread_thread_comments", Type: field.TypeInt},
		{Name: "created_by", Type: field.TypeUUID},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comments_threads_thread_comments",
				Columns:    []*schema.Column{CommentsColumns[5]},
				RefColumns: []*schema.Column{ThreadsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "comments_users_user_comments",
				Columns:    []*schema.Column{CommentsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// CommentKudosColumns holds the columns for the "comment_kudos" table.
	CommentKudosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "comment_id", Type: field.TypeInt},
	}
	// CommentKudosTable holds the schema information for the "comment_kudos" table.
	CommentKudosTable = &schema.Table{
		Name:       "comment_kudos",
		Columns:    CommentKudosColumns,
		PrimaryKey: []*schema.Column{CommentKudosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comment_kudos_users_user",
				Columns:    []*schema.Column{CommentKudosColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "comment_kudos_comments_comment",
				Columns:    []*schema.Column{CommentKudosColumns[2]},
				RefColumns: []*schema.Column{CommentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "commentkudo_user_id_comment_id",
				Unique:  true,
				Columns: []*schema.Column{CommentKudosColumns[1], CommentKudosColumns[2]},
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "tag_name", Type: field.TypeString, Size: 255},
		{Name: "tag_description", Type: field.TypeString, Size: 4096},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// ThreadsColumns holds the columns for the "threads" table.
	ThreadsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString, Size: 255},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 4096},
		{Name: "modified_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "topic_topic_threads", Type: field.TypeInt},
		{Name: "created_by", Type: field.TypeUUID},
	}
	// ThreadsTable holds the schema information for the "threads" table.
	ThreadsTable = &schema.Table{
		Name:       "threads",
		Columns:    ThreadsColumns,
		PrimaryKey: []*schema.Column{ThreadsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "threads_topics_topic_threads",
				Columns:    []*schema.Column{ThreadsColumns[5]},
				RefColumns: []*schema.Column{TopicsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "threads_users_user_threads",
				Columns:    []*schema.Column{ThreadsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ThreadKudosColumns holds the columns for the "thread_kudos" table.
	ThreadKudosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "thread_id", Type: field.TypeInt},
	}
	// ThreadKudosTable holds the schema information for the "thread_kudos" table.
	ThreadKudosTable = &schema.Table{
		Name:       "thread_kudos",
		Columns:    ThreadKudosColumns,
		PrimaryKey: []*schema.Column{ThreadKudosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "thread_kudos_users_user",
				Columns:    []*schema.Column{ThreadKudosColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "thread_kudos_threads_thread",
				Columns:    []*schema.Column{ThreadKudosColumns[2]},
				RefColumns: []*schema.Column{ThreadsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "threadkudo_user_id_thread_id",
				Unique:  true,
				Columns: []*schema.Column{ThreadKudosColumns[1], ThreadKudosColumns[2]},
			},
		},
	}
	// TopicsColumns holds the columns for the "topics" table.
	TopicsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString, Size: 255},
		{Name: "short_title", Type: field.TypeString, Size: 255},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 4096},
		{Name: "moderators", Type: field.TypeJSON},
		{Name: "created_by", Type: field.TypeString, Size: 255},
		{Name: "created_at", Type: field.TypeTime},
	}
	// TopicsTable holds the schema information for the "topics" table.
	TopicsTable = &schema.Table{
		Name:       "topics",
		Columns:    TopicsColumns,
		PrimaryKey: []*schema.Column{TopicsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "email", Type: field.TypeString, Size: 255},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 255},
		{Name: "first_name", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "last_name", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "hash", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "salt", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "is_moderator", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// ThreadTagsColumns holds the columns for the "thread_tags" table.
	ThreadTagsColumns = []*schema.Column{
		{Name: "thread_id", Type: field.TypeInt},
		{Name: "tag_id", Type: field.TypeInt},
	}
	// ThreadTagsTable holds the schema information for the "thread_tags" table.
	ThreadTagsTable = &schema.Table{
		Name:       "thread_tags",
		Columns:    ThreadTagsColumns,
		PrimaryKey: []*schema.Column{ThreadTagsColumns[0], ThreadTagsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "thread_tags_thread_id",
				Columns:    []*schema.Column{ThreadTagsColumns[0]},
				RefColumns: []*schema.Column{ThreadsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "thread_tags_tag_id",
				Columns:    []*schema.Column{ThreadTagsColumns[1]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CommentsTable,
		CommentKudosTable,
		TagsTable,
		ThreadsTable,
		ThreadKudosTable,
		TopicsTable,
		UsersTable,
		ThreadTagsTable,
	}
)

func init() {
	CommentsTable.ForeignKeys[0].RefTable = ThreadsTable
	CommentsTable.ForeignKeys[1].RefTable = UsersTable
	CommentKudosTable.ForeignKeys[0].RefTable = UsersTable
	CommentKudosTable.ForeignKeys[1].RefTable = CommentsTable
	ThreadsTable.ForeignKeys[0].RefTable = TopicsTable
	ThreadsTable.ForeignKeys[1].RefTable = UsersTable
	ThreadKudosTable.ForeignKeys[0].RefTable = UsersTable
	ThreadKudosTable.ForeignKeys[1].RefTable = ThreadsTable
	ThreadTagsTable.ForeignKeys[0].RefTable = ThreadsTable
	ThreadTagsTable.ForeignKeys[1].RefTable = TagsTable
}
