// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/interngowhere/web-backend/ent/comment"
	"github.com/interngowhere/web-backend/ent/schema"
	"github.com/interngowhere/web-backend/ent/tag"
	"github.com/interngowhere/web-backend/ent/thread"
	"github.com/interngowhere/web-backend/ent/topic"
	"github.com/interngowhere/web-backend/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	commentFields := schema.Comment{}.Fields()
	_ = commentFields
	// commentDescParentID is the schema descriptor for parent_id field.
	commentDescParentID := commentFields[0].Descriptor()
	// comment.DefaultParentID holds the default value on creation for the parent_id field.
	comment.DefaultParentID = commentDescParentID.Default.(string)
	// comment.ParentIDValidator is a validator for the "parent_id" field. It is called by the builders before save.
	comment.ParentIDValidator = commentDescParentID.Validators[0].(func(string) error)
	// commentDescContent is the schema descriptor for content field.
	commentDescContent := commentFields[1].Descriptor()
	// comment.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	comment.ContentValidator = commentDescContent.Validators[0].(func(string) error)
	// commentDescCreatedAt is the schema descriptor for created_at field.
	commentDescCreatedAt := commentFields[4].Descriptor()
	// comment.DefaultCreatedAt holds the default value on creation for the created_at field.
	comment.DefaultCreatedAt = commentDescCreatedAt.Default.(time.Time)
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescTagName is the schema descriptor for tag_name field.
	tagDescTagName := tagFields[0].Descriptor()
	// tag.TagNameValidator is a validator for the "tag_name" field. It is called by the builders before save.
	tag.TagNameValidator = tagDescTagName.Validators[0].(func(string) error)
	// tagDescTagDescription is the schema descriptor for tag_description field.
	tagDescTagDescription := tagFields[1].Descriptor()
	// tag.TagDescriptionValidator is a validator for the "tag_description" field. It is called by the builders before save.
	tag.TagDescriptionValidator = tagDescTagDescription.Validators[0].(func(string) error)
	threadFields := schema.Thread{}.Fields()
	_ = threadFields
	// threadDescTitle is the schema descriptor for title field.
	threadDescTitle := threadFields[0].Descriptor()
	// thread.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	thread.TitleValidator = threadDescTitle.Validators[0].(func(string) error)
	// threadDescDescription is the schema descriptor for description field.
	threadDescDescription := threadFields[1].Descriptor()
	// thread.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	thread.DescriptionValidator = threadDescDescription.Validators[0].(func(string) error)
	// threadDescCreatedAt is the schema descriptor for created_at field.
	threadDescCreatedAt := threadFields[4].Descriptor()
	// thread.DefaultCreatedAt holds the default value on creation for the created_at field.
	thread.DefaultCreatedAt = threadDescCreatedAt.Default.(time.Time)
	topicFields := schema.Topic{}.Fields()
	_ = topicFields
	// topicDescTitle is the schema descriptor for title field.
	topicDescTitle := topicFields[0].Descriptor()
	// topic.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	topic.TitleValidator = topicDescTitle.Validators[0].(func(string) error)
	// topicDescShortTitle is the schema descriptor for short_title field.
	topicDescShortTitle := topicFields[1].Descriptor()
	// topic.ShortTitleValidator is a validator for the "short_title" field. It is called by the builders before save.
	topic.ShortTitleValidator = topicDescShortTitle.Validators[0].(func(string) error)
	// topicDescDescription is the schema descriptor for description field.
	topicDescDescription := topicFields[2].Descriptor()
	// topic.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	topic.DescriptionValidator = topicDescDescription.Validators[0].(func(string) error)
	// topicDescCreatedBy is the schema descriptor for created_by field.
	topicDescCreatedBy := topicFields[4].Descriptor()
	// topic.CreatedByValidator is a validator for the "created_by" field. It is called by the builders before save.
	topic.CreatedByValidator = topicDescCreatedBy.Validators[0].(func(string) error)
	// topicDescCreatedAt is the schema descriptor for created_at field.
	topicDescCreatedAt := topicFields[5].Descriptor()
	// topic.DefaultCreatedAt holds the default value on creation for the created_at field.
	topic.DefaultCreatedAt = topicDescCreatedAt.Default.(time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[2].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescFirstName is the schema descriptor for first_name field.
	userDescFirstName := userFields[3].Descriptor()
	// user.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	user.FirstNameValidator = userDescFirstName.Validators[0].(func(string) error)
	// userDescLastName is the schema descriptor for last_name field.
	userDescLastName := userFields[4].Descriptor()
	// user.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	user.LastNameValidator = userDescLastName.Validators[0].(func(string) error)
	// userDescSalt is the schema descriptor for salt field.
	userDescSalt := userFields[6].Descriptor()
	// user.SaltValidator is a validator for the "salt" field. It is called by the builders before save.
	user.SaltValidator = userDescSalt.Validators[0].(func(string) error)
	// userDescIsModerator is the schema descriptor for is_moderator field.
	userDescIsModerator := userFields[7].Descriptor()
	// user.DefaultIsModerator holds the default value on creation for the is_moderator field.
	user.DefaultIsModerator = userDescIsModerator.Default.(bool)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[8].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}