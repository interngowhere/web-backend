// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/interngowhere/web-backend/ent/comment"
	"github.com/interngowhere/web-backend/ent/predicate"
	"github.com/interngowhere/web-backend/ent/thread"
	"github.com/interngowhere/web-backend/ent/topic"
	"github.com/interngowhere/web-backend/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUsername(s *string) *UserUpdate {
	if s != nil {
		uu.SetUsername(*s)
	}
	return uu
}

// SetFirstName sets the "first_name" field.
func (uu *UserUpdate) SetFirstName(s string) *UserUpdate {
	uu.mutation.SetFirstName(s)
	return uu
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableFirstName(s *string) *UserUpdate {
	if s != nil {
		uu.SetFirstName(*s)
	}
	return uu
}

// ClearFirstName clears the value of the "first_name" field.
func (uu *UserUpdate) ClearFirstName() *UserUpdate {
	uu.mutation.ClearFirstName()
	return uu
}

// SetLastName sets the "last_name" field.
func (uu *UserUpdate) SetLastName(s string) *UserUpdate {
	uu.mutation.SetLastName(s)
	return uu
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLastName(s *string) *UserUpdate {
	if s != nil {
		uu.SetLastName(*s)
	}
	return uu
}

// ClearLastName clears the value of the "last_name" field.
func (uu *UserUpdate) ClearLastName() *UserUpdate {
	uu.mutation.ClearLastName()
	return uu
}

// SetHash sets the "hash" field.
func (uu *UserUpdate) SetHash(b []byte) *UserUpdate {
	uu.mutation.SetHash(b)
	return uu
}

// SetEmailVerified sets the "email_verified" field.
func (uu *UserUpdate) SetEmailVerified(b bool) *UserUpdate {
	uu.mutation.SetEmailVerified(b)
	return uu
}

// SetNillableEmailVerified sets the "email_verified" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmailVerified(b *bool) *UserUpdate {
	if b != nil {
		uu.SetEmailVerified(*b)
	}
	return uu
}

// AddUserThreadIDs adds the "user_threads" edge to the Thread entity by IDs.
func (uu *UserUpdate) AddUserThreadIDs(ids ...int) *UserUpdate {
	uu.mutation.AddUserThreadIDs(ids...)
	return uu
}

// AddUserThreads adds the "user_threads" edges to the Thread entity.
func (uu *UserUpdate) AddUserThreads(t ...*Thread) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.AddUserThreadIDs(ids...)
}

// AddKudoedThreadIDs adds the "kudoed_threads" edge to the Thread entity by IDs.
func (uu *UserUpdate) AddKudoedThreadIDs(ids ...int) *UserUpdate {
	uu.mutation.AddKudoedThreadIDs(ids...)
	return uu
}

// AddKudoedThreads adds the "kudoed_threads" edges to the Thread entity.
func (uu *UserUpdate) AddKudoedThreads(t ...*Thread) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.AddKudoedThreadIDs(ids...)
}

// AddUserCommentIDs adds the "user_comments" edge to the Comment entity by IDs.
func (uu *UserUpdate) AddUserCommentIDs(ids ...int) *UserUpdate {
	uu.mutation.AddUserCommentIDs(ids...)
	return uu
}

// AddUserComments adds the "user_comments" edges to the Comment entity.
func (uu *UserUpdate) AddUserComments(c ...*Comment) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddUserCommentIDs(ids...)
}

// AddKudoedCommentIDs adds the "kudoed_comments" edge to the Comment entity by IDs.
func (uu *UserUpdate) AddKudoedCommentIDs(ids ...int) *UserUpdate {
	uu.mutation.AddKudoedCommentIDs(ids...)
	return uu
}

// AddKudoedComments adds the "kudoed_comments" edges to the Comment entity.
func (uu *UserUpdate) AddKudoedComments(c ...*Comment) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddKudoedCommentIDs(ids...)
}

// AddModeratedTopicIDs adds the "moderated_topics" edge to the Topic entity by IDs.
func (uu *UserUpdate) AddModeratedTopicIDs(ids ...int) *UserUpdate {
	uu.mutation.AddModeratedTopicIDs(ids...)
	return uu
}

// AddModeratedTopics adds the "moderated_topics" edges to the Topic entity.
func (uu *UserUpdate) AddModeratedTopics(t ...*Topic) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.AddModeratedTopicIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearUserThreads clears all "user_threads" edges to the Thread entity.
func (uu *UserUpdate) ClearUserThreads() *UserUpdate {
	uu.mutation.ClearUserThreads()
	return uu
}

// RemoveUserThreadIDs removes the "user_threads" edge to Thread entities by IDs.
func (uu *UserUpdate) RemoveUserThreadIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveUserThreadIDs(ids...)
	return uu
}

// RemoveUserThreads removes "user_threads" edges to Thread entities.
func (uu *UserUpdate) RemoveUserThreads(t ...*Thread) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.RemoveUserThreadIDs(ids...)
}

// ClearKudoedThreads clears all "kudoed_threads" edges to the Thread entity.
func (uu *UserUpdate) ClearKudoedThreads() *UserUpdate {
	uu.mutation.ClearKudoedThreads()
	return uu
}

// RemoveKudoedThreadIDs removes the "kudoed_threads" edge to Thread entities by IDs.
func (uu *UserUpdate) RemoveKudoedThreadIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveKudoedThreadIDs(ids...)
	return uu
}

// RemoveKudoedThreads removes "kudoed_threads" edges to Thread entities.
func (uu *UserUpdate) RemoveKudoedThreads(t ...*Thread) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.RemoveKudoedThreadIDs(ids...)
}

// ClearUserComments clears all "user_comments" edges to the Comment entity.
func (uu *UserUpdate) ClearUserComments() *UserUpdate {
	uu.mutation.ClearUserComments()
	return uu
}

// RemoveUserCommentIDs removes the "user_comments" edge to Comment entities by IDs.
func (uu *UserUpdate) RemoveUserCommentIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveUserCommentIDs(ids...)
	return uu
}

// RemoveUserComments removes "user_comments" edges to Comment entities.
func (uu *UserUpdate) RemoveUserComments(c ...*Comment) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveUserCommentIDs(ids...)
}

// ClearKudoedComments clears all "kudoed_comments" edges to the Comment entity.
func (uu *UserUpdate) ClearKudoedComments() *UserUpdate {
	uu.mutation.ClearKudoedComments()
	return uu
}

// RemoveKudoedCommentIDs removes the "kudoed_comments" edge to Comment entities by IDs.
func (uu *UserUpdate) RemoveKudoedCommentIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveKudoedCommentIDs(ids...)
	return uu
}

// RemoveKudoedComments removes "kudoed_comments" edges to Comment entities.
func (uu *UserUpdate) RemoveKudoedComments(c ...*Comment) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveKudoedCommentIDs(ids...)
}

// ClearModeratedTopics clears all "moderated_topics" edges to the Topic entity.
func (uu *UserUpdate) ClearModeratedTopics() *UserUpdate {
	uu.mutation.ClearModeratedTopics()
	return uu
}

// RemoveModeratedTopicIDs removes the "moderated_topics" edge to Topic entities by IDs.
func (uu *UserUpdate) RemoveModeratedTopicIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveModeratedTopicIDs(ids...)
	return uu
}

// RemoveModeratedTopics removes "moderated_topics" edges to Topic entities.
func (uu *UserUpdate) RemoveModeratedTopics(t ...*Topic) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.RemoveModeratedTopicIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if v, ok := uu.mutation.FirstName(); ok {
		if err := user.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "User.first_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.LastName(); ok {
		if err := user.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "User.last_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Hash(); ok {
		if err := user.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf(`ent: validator failed for field "User.hash": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if uu.mutation.FirstNameCleared() {
		_spec.ClearField(user.FieldFirstName, field.TypeString)
	}
	if value, ok := uu.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if uu.mutation.LastNameCleared() {
		_spec.ClearField(user.FieldLastName, field.TypeString)
	}
	if value, ok := uu.mutation.Hash(); ok {
		_spec.SetField(user.FieldHash, field.TypeBytes, value)
	}
	if value, ok := uu.mutation.EmailVerified(); ok {
		_spec.SetField(user.FieldEmailVerified, field.TypeBool, value)
	}
	if uu.mutation.UserThreadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserThreadsTable,
			Columns: []string{user.UserThreadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedUserThreadsIDs(); len(nodes) > 0 && !uu.mutation.UserThreadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserThreadsTable,
			Columns: []string{user.UserThreadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UserThreadsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserThreadsTable,
			Columns: []string{user.UserThreadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.KudoedThreadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.KudoedThreadsTable,
			Columns: user.KudoedThreadsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedKudoedThreadsIDs(); len(nodes) > 0 && !uu.mutation.KudoedThreadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.KudoedThreadsTable,
			Columns: user.KudoedThreadsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.KudoedThreadsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.KudoedThreadsTable,
			Columns: user.KudoedThreadsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.UserCommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserCommentsTable,
			Columns: []string{user.UserCommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedUserCommentsIDs(); len(nodes) > 0 && !uu.mutation.UserCommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserCommentsTable,
			Columns: []string{user.UserCommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UserCommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserCommentsTable,
			Columns: []string{user.UserCommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.KudoedCommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.KudoedCommentsTable,
			Columns: user.KudoedCommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedKudoedCommentsIDs(); len(nodes) > 0 && !uu.mutation.KudoedCommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.KudoedCommentsTable,
			Columns: user.KudoedCommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.KudoedCommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.KudoedCommentsTable,
			Columns: user.KudoedCommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.ModeratedTopicsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.ModeratedTopicsTable,
			Columns: user.ModeratedTopicsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedModeratedTopicsIDs(); len(nodes) > 0 && !uu.mutation.ModeratedTopicsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.ModeratedTopicsTable,
			Columns: user.ModeratedTopicsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ModeratedTopicsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.ModeratedTopicsTable,
			Columns: user.ModeratedTopicsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUsername(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUsername(*s)
	}
	return uuo
}

// SetFirstName sets the "first_name" field.
func (uuo *UserUpdateOne) SetFirstName(s string) *UserUpdateOne {
	uuo.mutation.SetFirstName(s)
	return uuo
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableFirstName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetFirstName(*s)
	}
	return uuo
}

// ClearFirstName clears the value of the "first_name" field.
func (uuo *UserUpdateOne) ClearFirstName() *UserUpdateOne {
	uuo.mutation.ClearFirstName()
	return uuo
}

// SetLastName sets the "last_name" field.
func (uuo *UserUpdateOne) SetLastName(s string) *UserUpdateOne {
	uuo.mutation.SetLastName(s)
	return uuo
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLastName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetLastName(*s)
	}
	return uuo
}

// ClearLastName clears the value of the "last_name" field.
func (uuo *UserUpdateOne) ClearLastName() *UserUpdateOne {
	uuo.mutation.ClearLastName()
	return uuo
}

// SetHash sets the "hash" field.
func (uuo *UserUpdateOne) SetHash(b []byte) *UserUpdateOne {
	uuo.mutation.SetHash(b)
	return uuo
}

// SetEmailVerified sets the "email_verified" field.
func (uuo *UserUpdateOne) SetEmailVerified(b bool) *UserUpdateOne {
	uuo.mutation.SetEmailVerified(b)
	return uuo
}

// SetNillableEmailVerified sets the "email_verified" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmailVerified(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetEmailVerified(*b)
	}
	return uuo
}

// AddUserThreadIDs adds the "user_threads" edge to the Thread entity by IDs.
func (uuo *UserUpdateOne) AddUserThreadIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddUserThreadIDs(ids...)
	return uuo
}

// AddUserThreads adds the "user_threads" edges to the Thread entity.
func (uuo *UserUpdateOne) AddUserThreads(t ...*Thread) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.AddUserThreadIDs(ids...)
}

// AddKudoedThreadIDs adds the "kudoed_threads" edge to the Thread entity by IDs.
func (uuo *UserUpdateOne) AddKudoedThreadIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddKudoedThreadIDs(ids...)
	return uuo
}

// AddKudoedThreads adds the "kudoed_threads" edges to the Thread entity.
func (uuo *UserUpdateOne) AddKudoedThreads(t ...*Thread) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.AddKudoedThreadIDs(ids...)
}

// AddUserCommentIDs adds the "user_comments" edge to the Comment entity by IDs.
func (uuo *UserUpdateOne) AddUserCommentIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddUserCommentIDs(ids...)
	return uuo
}

// AddUserComments adds the "user_comments" edges to the Comment entity.
func (uuo *UserUpdateOne) AddUserComments(c ...*Comment) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddUserCommentIDs(ids...)
}

// AddKudoedCommentIDs adds the "kudoed_comments" edge to the Comment entity by IDs.
func (uuo *UserUpdateOne) AddKudoedCommentIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddKudoedCommentIDs(ids...)
	return uuo
}

// AddKudoedComments adds the "kudoed_comments" edges to the Comment entity.
func (uuo *UserUpdateOne) AddKudoedComments(c ...*Comment) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddKudoedCommentIDs(ids...)
}

// AddModeratedTopicIDs adds the "moderated_topics" edge to the Topic entity by IDs.
func (uuo *UserUpdateOne) AddModeratedTopicIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddModeratedTopicIDs(ids...)
	return uuo
}

// AddModeratedTopics adds the "moderated_topics" edges to the Topic entity.
func (uuo *UserUpdateOne) AddModeratedTopics(t ...*Topic) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.AddModeratedTopicIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearUserThreads clears all "user_threads" edges to the Thread entity.
func (uuo *UserUpdateOne) ClearUserThreads() *UserUpdateOne {
	uuo.mutation.ClearUserThreads()
	return uuo
}

// RemoveUserThreadIDs removes the "user_threads" edge to Thread entities by IDs.
func (uuo *UserUpdateOne) RemoveUserThreadIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveUserThreadIDs(ids...)
	return uuo
}

// RemoveUserThreads removes "user_threads" edges to Thread entities.
func (uuo *UserUpdateOne) RemoveUserThreads(t ...*Thread) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.RemoveUserThreadIDs(ids...)
}

// ClearKudoedThreads clears all "kudoed_threads" edges to the Thread entity.
func (uuo *UserUpdateOne) ClearKudoedThreads() *UserUpdateOne {
	uuo.mutation.ClearKudoedThreads()
	return uuo
}

// RemoveKudoedThreadIDs removes the "kudoed_threads" edge to Thread entities by IDs.
func (uuo *UserUpdateOne) RemoveKudoedThreadIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveKudoedThreadIDs(ids...)
	return uuo
}

// RemoveKudoedThreads removes "kudoed_threads" edges to Thread entities.
func (uuo *UserUpdateOne) RemoveKudoedThreads(t ...*Thread) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.RemoveKudoedThreadIDs(ids...)
}

// ClearUserComments clears all "user_comments" edges to the Comment entity.
func (uuo *UserUpdateOne) ClearUserComments() *UserUpdateOne {
	uuo.mutation.ClearUserComments()
	return uuo
}

// RemoveUserCommentIDs removes the "user_comments" edge to Comment entities by IDs.
func (uuo *UserUpdateOne) RemoveUserCommentIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveUserCommentIDs(ids...)
	return uuo
}

// RemoveUserComments removes "user_comments" edges to Comment entities.
func (uuo *UserUpdateOne) RemoveUserComments(c ...*Comment) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveUserCommentIDs(ids...)
}

// ClearKudoedComments clears all "kudoed_comments" edges to the Comment entity.
func (uuo *UserUpdateOne) ClearKudoedComments() *UserUpdateOne {
	uuo.mutation.ClearKudoedComments()
	return uuo
}

// RemoveKudoedCommentIDs removes the "kudoed_comments" edge to Comment entities by IDs.
func (uuo *UserUpdateOne) RemoveKudoedCommentIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveKudoedCommentIDs(ids...)
	return uuo
}

// RemoveKudoedComments removes "kudoed_comments" edges to Comment entities.
func (uuo *UserUpdateOne) RemoveKudoedComments(c ...*Comment) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveKudoedCommentIDs(ids...)
}

// ClearModeratedTopics clears all "moderated_topics" edges to the Topic entity.
func (uuo *UserUpdateOne) ClearModeratedTopics() *UserUpdateOne {
	uuo.mutation.ClearModeratedTopics()
	return uuo
}

// RemoveModeratedTopicIDs removes the "moderated_topics" edge to Topic entities by IDs.
func (uuo *UserUpdateOne) RemoveModeratedTopicIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveModeratedTopicIDs(ids...)
	return uuo
}

// RemoveModeratedTopics removes "moderated_topics" edges to Topic entities.
func (uuo *UserUpdateOne) RemoveModeratedTopics(t ...*Topic) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.RemoveModeratedTopicIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.FirstName(); ok {
		if err := user.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "User.first_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.LastName(); ok {
		if err := user.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "User.last_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Hash(); ok {
		if err := user.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf(`ent: validator failed for field "User.hash": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if uuo.mutation.FirstNameCleared() {
		_spec.ClearField(user.FieldFirstName, field.TypeString)
	}
	if value, ok := uuo.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if uuo.mutation.LastNameCleared() {
		_spec.ClearField(user.FieldLastName, field.TypeString)
	}
	if value, ok := uuo.mutation.Hash(); ok {
		_spec.SetField(user.FieldHash, field.TypeBytes, value)
	}
	if value, ok := uuo.mutation.EmailVerified(); ok {
		_spec.SetField(user.FieldEmailVerified, field.TypeBool, value)
	}
	if uuo.mutation.UserThreadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserThreadsTable,
			Columns: []string{user.UserThreadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedUserThreadsIDs(); len(nodes) > 0 && !uuo.mutation.UserThreadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserThreadsTable,
			Columns: []string{user.UserThreadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UserThreadsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserThreadsTable,
			Columns: []string{user.UserThreadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.KudoedThreadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.KudoedThreadsTable,
			Columns: user.KudoedThreadsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedKudoedThreadsIDs(); len(nodes) > 0 && !uuo.mutation.KudoedThreadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.KudoedThreadsTable,
			Columns: user.KudoedThreadsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.KudoedThreadsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.KudoedThreadsTable,
			Columns: user.KudoedThreadsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.UserCommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserCommentsTable,
			Columns: []string{user.UserCommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedUserCommentsIDs(); len(nodes) > 0 && !uuo.mutation.UserCommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserCommentsTable,
			Columns: []string{user.UserCommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UserCommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserCommentsTable,
			Columns: []string{user.UserCommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.KudoedCommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.KudoedCommentsTable,
			Columns: user.KudoedCommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedKudoedCommentsIDs(); len(nodes) > 0 && !uuo.mutation.KudoedCommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.KudoedCommentsTable,
			Columns: user.KudoedCommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.KudoedCommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.KudoedCommentsTable,
			Columns: user.KudoedCommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.ModeratedTopicsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.ModeratedTopicsTable,
			Columns: user.ModeratedTopicsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedModeratedTopicsIDs(); len(nodes) > 0 && !uuo.mutation.ModeratedTopicsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.ModeratedTopicsTable,
			Columns: user.ModeratedTopicsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ModeratedTopicsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.ModeratedTopicsTable,
			Columns: user.ModeratedTopicsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
