// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/interngowhere/web-backend/ent/comment"
	"github.com/interngowhere/web-backend/ent/tag"
	"github.com/interngowhere/web-backend/ent/thread"
	"github.com/interngowhere/web-backend/ent/topic"
	"github.com/interngowhere/web-backend/ent/user"
)

// ThreadCreate is the builder for creating a Thread entity.
type ThreadCreate struct {
	config
	mutation *ThreadMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (tc *ThreadCreate) SetTitle(s string) *ThreadCreate {
	tc.mutation.SetTitle(s)
	return tc
}

// SetSlug sets the "slug" field.
func (tc *ThreadCreate) SetSlug(s string) *ThreadCreate {
	tc.mutation.SetSlug(s)
	return tc
}

// SetDescription sets the "description" field.
func (tc *ThreadCreate) SetDescription(s string) *ThreadCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tc *ThreadCreate) SetNillableDescription(s *string) *ThreadCreate {
	if s != nil {
		tc.SetDescription(*s)
	}
	return tc
}

// SetModifiedAt sets the "modified_at" field.
func (tc *ThreadCreate) SetModifiedAt(t time.Time) *ThreadCreate {
	tc.mutation.SetModifiedAt(t)
	return tc
}

// SetNillableModifiedAt sets the "modified_at" field if the given value is not nil.
func (tc *ThreadCreate) SetNillableModifiedAt(t *time.Time) *ThreadCreate {
	if t != nil {
		tc.SetModifiedAt(*t)
	}
	return tc
}

// SetCreatedBy sets the "created_by" field.
func (tc *ThreadCreate) SetCreatedBy(u uuid.UUID) *ThreadCreate {
	tc.mutation.SetCreatedBy(u)
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *ThreadCreate) SetCreatedAt(t time.Time) *ThreadCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *ThreadCreate) SetNillableCreatedAt(t *time.Time) *ThreadCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// AddThreadCommentIDs adds the "thread_comments" edge to the Comment entity by IDs.
func (tc *ThreadCreate) AddThreadCommentIDs(ids ...int) *ThreadCreate {
	tc.mutation.AddThreadCommentIDs(ids...)
	return tc
}

// AddThreadComments adds the "thread_comments" edges to the Comment entity.
func (tc *ThreadCreate) AddThreadComments(c ...*Comment) *ThreadCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tc.AddThreadCommentIDs(ids...)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (tc *ThreadCreate) AddTagIDs(ids ...int) *ThreadCreate {
	tc.mutation.AddTagIDs(ids...)
	return tc
}

// AddTags adds the "tags" edges to the Tag entity.
func (tc *ThreadCreate) AddTags(t ...*Tag) *ThreadCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddTagIDs(ids...)
}

// SetTopicsID sets the "topics" edge to the Topic entity by ID.
func (tc *ThreadCreate) SetTopicsID(id int) *ThreadCreate {
	tc.mutation.SetTopicsID(id)
	return tc
}

// SetTopics sets the "topics" edge to the Topic entity.
func (tc *ThreadCreate) SetTopics(t *Topic) *ThreadCreate {
	return tc.SetTopicsID(t.ID)
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (tc *ThreadCreate) SetUsersID(id uuid.UUID) *ThreadCreate {
	tc.mutation.SetUsersID(id)
	return tc
}

// SetUsers sets the "users" edge to the User entity.
func (tc *ThreadCreate) SetUsers(u *User) *ThreadCreate {
	return tc.SetUsersID(u.ID)
}

// AddKudoedUserIDs adds the "kudoed_users" edge to the User entity by IDs.
func (tc *ThreadCreate) AddKudoedUserIDs(ids ...uuid.UUID) *ThreadCreate {
	tc.mutation.AddKudoedUserIDs(ids...)
	return tc
}

// AddKudoedUsers adds the "kudoed_users" edges to the User entity.
func (tc *ThreadCreate) AddKudoedUsers(u ...*User) *ThreadCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tc.AddKudoedUserIDs(ids...)
}

// Mutation returns the ThreadMutation object of the builder.
func (tc *ThreadCreate) Mutation() *ThreadMutation {
	return tc.mutation
}

// Save creates the Thread in the database.
func (tc *ThreadCreate) Save(ctx context.Context) (*Thread, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *ThreadCreate) SaveX(ctx context.Context) *Thread {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *ThreadCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *ThreadCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *ThreadCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := thread.DefaultCreatedAt
		tc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *ThreadCreate) check() error {
	if _, ok := tc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Thread.title"`)}
	}
	if v, ok := tc.mutation.Title(); ok {
		if err := thread.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Thread.title": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "Thread.slug"`)}
	}
	if v, ok := tc.mutation.Slug(); ok {
		if err := thread.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "Thread.slug": %w`, err)}
		}
	}
	if v, ok := tc.mutation.Description(); ok {
		if err := thread.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Thread.description": %w`, err)}
		}
	}
	if _, ok := tc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "Thread.created_by"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Thread.created_at"`)}
	}
	if _, ok := tc.mutation.TopicsID(); !ok {
		return &ValidationError{Name: "topics", err: errors.New(`ent: missing required edge "Thread.topics"`)}
	}
	if _, ok := tc.mutation.UsersID(); !ok {
		return &ValidationError{Name: "users", err: errors.New(`ent: missing required edge "Thread.users"`)}
	}
	return nil
}

func (tc *ThreadCreate) sqlSave(ctx context.Context) (*Thread, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *ThreadCreate) createSpec() (*Thread, *sqlgraph.CreateSpec) {
	var (
		_node = &Thread{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(thread.Table, sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt))
	)
	if value, ok := tc.mutation.Title(); ok {
		_spec.SetField(thread.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := tc.mutation.Slug(); ok {
		_spec.SetField(thread.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.SetField(thread.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := tc.mutation.ModifiedAt(); ok {
		_spec.SetField(thread.FieldModifiedAt, field.TypeTime, value)
		_node.ModifiedAt = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(thread.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := tc.mutation.ThreadCommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thread.ThreadCommentsTable,
			Columns: []string{thread.ThreadCommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   thread.TagsTable,
			Columns: thread.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TopicsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   thread.TopicsTable,
			Columns: []string{thread.TopicsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.topic_topic_threads = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   thread.UsersTable,
			Columns: []string{thread.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CreatedBy = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.KudoedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   thread.KudoedUsersTable,
			Columns: thread.KudoedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ThreadCreateBulk is the builder for creating many Thread entities in bulk.
type ThreadCreateBulk struct {
	config
	err      error
	builders []*ThreadCreate
}

// Save creates the Thread entities in the database.
func (tcb *ThreadCreateBulk) Save(ctx context.Context) ([]*Thread, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Thread, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ThreadMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *ThreadCreateBulk) SaveX(ctx context.Context) []*Thread {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *ThreadCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *ThreadCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
