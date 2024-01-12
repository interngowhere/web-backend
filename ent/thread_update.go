// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/interngowhere/web-backend/ent/comment"
	"github.com/interngowhere/web-backend/ent/predicate"
	"github.com/interngowhere/web-backend/ent/tag"
	"github.com/interngowhere/web-backend/ent/thread"
	"github.com/interngowhere/web-backend/ent/topic"
	"github.com/interngowhere/web-backend/ent/user"
)

// ThreadUpdate is the builder for updating Thread entities.
type ThreadUpdate struct {
	config
	hooks    []Hook
	mutation *ThreadMutation
}

// Where appends a list predicates to the ThreadUpdate builder.
func (tu *ThreadUpdate) Where(ps ...predicate.Thread) *ThreadUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetTitle sets the "title" field.
func (tu *ThreadUpdate) SetTitle(s string) *ThreadUpdate {
	tu.mutation.SetTitle(s)
	return tu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (tu *ThreadUpdate) SetNillableTitle(s *string) *ThreadUpdate {
	if s != nil {
		tu.SetTitle(*s)
	}
	return tu
}

// SetSlug sets the "slug" field.
func (tu *ThreadUpdate) SetSlug(s string) *ThreadUpdate {
	tu.mutation.SetSlug(s)
	return tu
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (tu *ThreadUpdate) SetNillableSlug(s *string) *ThreadUpdate {
	if s != nil {
		tu.SetSlug(*s)
	}
	return tu
}

// SetDescription sets the "description" field.
func (tu *ThreadUpdate) SetDescription(s string) *ThreadUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tu *ThreadUpdate) SetNillableDescription(s *string) *ThreadUpdate {
	if s != nil {
		tu.SetDescription(*s)
	}
	return tu
}

// ClearDescription clears the value of the "description" field.
func (tu *ThreadUpdate) ClearDescription() *ThreadUpdate {
	tu.mutation.ClearDescription()
	return tu
}

// SetModifiedAt sets the "modified_at" field.
func (tu *ThreadUpdate) SetModifiedAt(t time.Time) *ThreadUpdate {
	tu.mutation.SetModifiedAt(t)
	return tu
}

// SetNillableModifiedAt sets the "modified_at" field if the given value is not nil.
func (tu *ThreadUpdate) SetNillableModifiedAt(t *time.Time) *ThreadUpdate {
	if t != nil {
		tu.SetModifiedAt(*t)
	}
	return tu
}

// ClearModifiedAt clears the value of the "modified_at" field.
func (tu *ThreadUpdate) ClearModifiedAt() *ThreadUpdate {
	tu.mutation.ClearModifiedAt()
	return tu
}

// AddThreadCommentIDs adds the "thread_comments" edge to the Comment entity by IDs.
func (tu *ThreadUpdate) AddThreadCommentIDs(ids ...int) *ThreadUpdate {
	tu.mutation.AddThreadCommentIDs(ids...)
	return tu
}

// AddThreadComments adds the "thread_comments" edges to the Comment entity.
func (tu *ThreadUpdate) AddThreadComments(c ...*Comment) *ThreadUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tu.AddThreadCommentIDs(ids...)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (tu *ThreadUpdate) AddTagIDs(ids ...int) *ThreadUpdate {
	tu.mutation.AddTagIDs(ids...)
	return tu
}

// AddTags adds the "tags" edges to the Tag entity.
func (tu *ThreadUpdate) AddTags(t ...*Tag) *ThreadUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTagIDs(ids...)
}

// SetTopicsID sets the "topics" edge to the Topic entity by ID.
func (tu *ThreadUpdate) SetTopicsID(id int) *ThreadUpdate {
	tu.mutation.SetTopicsID(id)
	return tu
}

// SetTopics sets the "topics" edge to the Topic entity.
func (tu *ThreadUpdate) SetTopics(t *Topic) *ThreadUpdate {
	return tu.SetTopicsID(t.ID)
}

// AddKudoedUserIDs adds the "kudoed_users" edge to the User entity by IDs.
func (tu *ThreadUpdate) AddKudoedUserIDs(ids ...uuid.UUID) *ThreadUpdate {
	tu.mutation.AddKudoedUserIDs(ids...)
	return tu
}

// AddKudoedUsers adds the "kudoed_users" edges to the User entity.
func (tu *ThreadUpdate) AddKudoedUsers(u ...*User) *ThreadUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.AddKudoedUserIDs(ids...)
}

// Mutation returns the ThreadMutation object of the builder.
func (tu *ThreadUpdate) Mutation() *ThreadMutation {
	return tu.mutation
}

// ClearThreadComments clears all "thread_comments" edges to the Comment entity.
func (tu *ThreadUpdate) ClearThreadComments() *ThreadUpdate {
	tu.mutation.ClearThreadComments()
	return tu
}

// RemoveThreadCommentIDs removes the "thread_comments" edge to Comment entities by IDs.
func (tu *ThreadUpdate) RemoveThreadCommentIDs(ids ...int) *ThreadUpdate {
	tu.mutation.RemoveThreadCommentIDs(ids...)
	return tu
}

// RemoveThreadComments removes "thread_comments" edges to Comment entities.
func (tu *ThreadUpdate) RemoveThreadComments(c ...*Comment) *ThreadUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tu.RemoveThreadCommentIDs(ids...)
}

// ClearTags clears all "tags" edges to the Tag entity.
func (tu *ThreadUpdate) ClearTags() *ThreadUpdate {
	tu.mutation.ClearTags()
	return tu
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (tu *ThreadUpdate) RemoveTagIDs(ids ...int) *ThreadUpdate {
	tu.mutation.RemoveTagIDs(ids...)
	return tu
}

// RemoveTags removes "tags" edges to Tag entities.
func (tu *ThreadUpdate) RemoveTags(t ...*Tag) *ThreadUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTagIDs(ids...)
}

// ClearTopics clears the "topics" edge to the Topic entity.
func (tu *ThreadUpdate) ClearTopics() *ThreadUpdate {
	tu.mutation.ClearTopics()
	return tu
}

// ClearKudoedUsers clears all "kudoed_users" edges to the User entity.
func (tu *ThreadUpdate) ClearKudoedUsers() *ThreadUpdate {
	tu.mutation.ClearKudoedUsers()
	return tu
}

// RemoveKudoedUserIDs removes the "kudoed_users" edge to User entities by IDs.
func (tu *ThreadUpdate) RemoveKudoedUserIDs(ids ...uuid.UUID) *ThreadUpdate {
	tu.mutation.RemoveKudoedUserIDs(ids...)
	return tu
}

// RemoveKudoedUsers removes "kudoed_users" edges to User entities.
func (tu *ThreadUpdate) RemoveKudoedUsers(u ...*User) *ThreadUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.RemoveKudoedUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *ThreadUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *ThreadUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *ThreadUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *ThreadUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *ThreadUpdate) check() error {
	if v, ok := tu.mutation.Title(); ok {
		if err := thread.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Thread.title": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Slug(); ok {
		if err := thread.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "Thread.slug": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Description(); ok {
		if err := thread.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Thread.description": %w`, err)}
		}
	}
	if _, ok := tu.mutation.TopicsID(); tu.mutation.TopicsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Thread.topics"`)
	}
	if _, ok := tu.mutation.UsersID(); tu.mutation.UsersCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Thread.users"`)
	}
	return nil
}

func (tu *ThreadUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(thread.Table, thread.Columns, sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Title(); ok {
		_spec.SetField(thread.FieldTitle, field.TypeString, value)
	}
	if value, ok := tu.mutation.Slug(); ok {
		_spec.SetField(thread.FieldSlug, field.TypeString, value)
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.SetField(thread.FieldDescription, field.TypeString, value)
	}
	if tu.mutation.DescriptionCleared() {
		_spec.ClearField(thread.FieldDescription, field.TypeString)
	}
	if value, ok := tu.mutation.ModifiedAt(); ok {
		_spec.SetField(thread.FieldModifiedAt, field.TypeTime, value)
	}
	if tu.mutation.ModifiedAtCleared() {
		_spec.ClearField(thread.FieldModifiedAt, field.TypeTime)
	}
	if tu.mutation.ThreadCommentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedThreadCommentsIDs(); len(nodes) > 0 && !tu.mutation.ThreadCommentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ThreadCommentsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.TagsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedTagsIDs(); len(nodes) > 0 && !tu.mutation.TagsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TagsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.TopicsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TopicsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.KudoedUsersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedKudoedUsersIDs(); len(nodes) > 0 && !tu.mutation.KudoedUsersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.KudoedUsersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{thread.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// ThreadUpdateOne is the builder for updating a single Thread entity.
type ThreadUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ThreadMutation
}

// SetTitle sets the "title" field.
func (tuo *ThreadUpdateOne) SetTitle(s string) *ThreadUpdateOne {
	tuo.mutation.SetTitle(s)
	return tuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (tuo *ThreadUpdateOne) SetNillableTitle(s *string) *ThreadUpdateOne {
	if s != nil {
		tuo.SetTitle(*s)
	}
	return tuo
}

// SetSlug sets the "slug" field.
func (tuo *ThreadUpdateOne) SetSlug(s string) *ThreadUpdateOne {
	tuo.mutation.SetSlug(s)
	return tuo
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (tuo *ThreadUpdateOne) SetNillableSlug(s *string) *ThreadUpdateOne {
	if s != nil {
		tuo.SetSlug(*s)
	}
	return tuo
}

// SetDescription sets the "description" field.
func (tuo *ThreadUpdateOne) SetDescription(s string) *ThreadUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tuo *ThreadUpdateOne) SetNillableDescription(s *string) *ThreadUpdateOne {
	if s != nil {
		tuo.SetDescription(*s)
	}
	return tuo
}

// ClearDescription clears the value of the "description" field.
func (tuo *ThreadUpdateOne) ClearDescription() *ThreadUpdateOne {
	tuo.mutation.ClearDescription()
	return tuo
}

// SetModifiedAt sets the "modified_at" field.
func (tuo *ThreadUpdateOne) SetModifiedAt(t time.Time) *ThreadUpdateOne {
	tuo.mutation.SetModifiedAt(t)
	return tuo
}

// SetNillableModifiedAt sets the "modified_at" field if the given value is not nil.
func (tuo *ThreadUpdateOne) SetNillableModifiedAt(t *time.Time) *ThreadUpdateOne {
	if t != nil {
		tuo.SetModifiedAt(*t)
	}
	return tuo
}

// ClearModifiedAt clears the value of the "modified_at" field.
func (tuo *ThreadUpdateOne) ClearModifiedAt() *ThreadUpdateOne {
	tuo.mutation.ClearModifiedAt()
	return tuo
}

// AddThreadCommentIDs adds the "thread_comments" edge to the Comment entity by IDs.
func (tuo *ThreadUpdateOne) AddThreadCommentIDs(ids ...int) *ThreadUpdateOne {
	tuo.mutation.AddThreadCommentIDs(ids...)
	return tuo
}

// AddThreadComments adds the "thread_comments" edges to the Comment entity.
func (tuo *ThreadUpdateOne) AddThreadComments(c ...*Comment) *ThreadUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tuo.AddThreadCommentIDs(ids...)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (tuo *ThreadUpdateOne) AddTagIDs(ids ...int) *ThreadUpdateOne {
	tuo.mutation.AddTagIDs(ids...)
	return tuo
}

// AddTags adds the "tags" edges to the Tag entity.
func (tuo *ThreadUpdateOne) AddTags(t ...*Tag) *ThreadUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTagIDs(ids...)
}

// SetTopicsID sets the "topics" edge to the Topic entity by ID.
func (tuo *ThreadUpdateOne) SetTopicsID(id int) *ThreadUpdateOne {
	tuo.mutation.SetTopicsID(id)
	return tuo
}

// SetTopics sets the "topics" edge to the Topic entity.
func (tuo *ThreadUpdateOne) SetTopics(t *Topic) *ThreadUpdateOne {
	return tuo.SetTopicsID(t.ID)
}

// AddKudoedUserIDs adds the "kudoed_users" edge to the User entity by IDs.
func (tuo *ThreadUpdateOne) AddKudoedUserIDs(ids ...uuid.UUID) *ThreadUpdateOne {
	tuo.mutation.AddKudoedUserIDs(ids...)
	return tuo
}

// AddKudoedUsers adds the "kudoed_users" edges to the User entity.
func (tuo *ThreadUpdateOne) AddKudoedUsers(u ...*User) *ThreadUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.AddKudoedUserIDs(ids...)
}

// Mutation returns the ThreadMutation object of the builder.
func (tuo *ThreadUpdateOne) Mutation() *ThreadMutation {
	return tuo.mutation
}

// ClearThreadComments clears all "thread_comments" edges to the Comment entity.
func (tuo *ThreadUpdateOne) ClearThreadComments() *ThreadUpdateOne {
	tuo.mutation.ClearThreadComments()
	return tuo
}

// RemoveThreadCommentIDs removes the "thread_comments" edge to Comment entities by IDs.
func (tuo *ThreadUpdateOne) RemoveThreadCommentIDs(ids ...int) *ThreadUpdateOne {
	tuo.mutation.RemoveThreadCommentIDs(ids...)
	return tuo
}

// RemoveThreadComments removes "thread_comments" edges to Comment entities.
func (tuo *ThreadUpdateOne) RemoveThreadComments(c ...*Comment) *ThreadUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tuo.RemoveThreadCommentIDs(ids...)
}

// ClearTags clears all "tags" edges to the Tag entity.
func (tuo *ThreadUpdateOne) ClearTags() *ThreadUpdateOne {
	tuo.mutation.ClearTags()
	return tuo
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (tuo *ThreadUpdateOne) RemoveTagIDs(ids ...int) *ThreadUpdateOne {
	tuo.mutation.RemoveTagIDs(ids...)
	return tuo
}

// RemoveTags removes "tags" edges to Tag entities.
func (tuo *ThreadUpdateOne) RemoveTags(t ...*Tag) *ThreadUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTagIDs(ids...)
}

// ClearTopics clears the "topics" edge to the Topic entity.
func (tuo *ThreadUpdateOne) ClearTopics() *ThreadUpdateOne {
	tuo.mutation.ClearTopics()
	return tuo
}

// ClearKudoedUsers clears all "kudoed_users" edges to the User entity.
func (tuo *ThreadUpdateOne) ClearKudoedUsers() *ThreadUpdateOne {
	tuo.mutation.ClearKudoedUsers()
	return tuo
}

// RemoveKudoedUserIDs removes the "kudoed_users" edge to User entities by IDs.
func (tuo *ThreadUpdateOne) RemoveKudoedUserIDs(ids ...uuid.UUID) *ThreadUpdateOne {
	tuo.mutation.RemoveKudoedUserIDs(ids...)
	return tuo
}

// RemoveKudoedUsers removes "kudoed_users" edges to User entities.
func (tuo *ThreadUpdateOne) RemoveKudoedUsers(u ...*User) *ThreadUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.RemoveKudoedUserIDs(ids...)
}

// Where appends a list predicates to the ThreadUpdate builder.
func (tuo *ThreadUpdateOne) Where(ps ...predicate.Thread) *ThreadUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *ThreadUpdateOne) Select(field string, fields ...string) *ThreadUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Thread entity.
func (tuo *ThreadUpdateOne) Save(ctx context.Context) (*Thread, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *ThreadUpdateOne) SaveX(ctx context.Context) *Thread {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *ThreadUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *ThreadUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *ThreadUpdateOne) check() error {
	if v, ok := tuo.mutation.Title(); ok {
		if err := thread.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Thread.title": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Slug(); ok {
		if err := thread.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "Thread.slug": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Description(); ok {
		if err := thread.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Thread.description": %w`, err)}
		}
	}
	if _, ok := tuo.mutation.TopicsID(); tuo.mutation.TopicsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Thread.topics"`)
	}
	if _, ok := tuo.mutation.UsersID(); tuo.mutation.UsersCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Thread.users"`)
	}
	return nil
}

func (tuo *ThreadUpdateOne) sqlSave(ctx context.Context) (_node *Thread, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(thread.Table, thread.Columns, sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Thread.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, thread.FieldID)
		for _, f := range fields {
			if !thread.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != thread.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Title(); ok {
		_spec.SetField(thread.FieldTitle, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Slug(); ok {
		_spec.SetField(thread.FieldSlug, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.SetField(thread.FieldDescription, field.TypeString, value)
	}
	if tuo.mutation.DescriptionCleared() {
		_spec.ClearField(thread.FieldDescription, field.TypeString)
	}
	if value, ok := tuo.mutation.ModifiedAt(); ok {
		_spec.SetField(thread.FieldModifiedAt, field.TypeTime, value)
	}
	if tuo.mutation.ModifiedAtCleared() {
		_spec.ClearField(thread.FieldModifiedAt, field.TypeTime)
	}
	if tuo.mutation.ThreadCommentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedThreadCommentsIDs(); len(nodes) > 0 && !tuo.mutation.ThreadCommentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ThreadCommentsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.TagsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedTagsIDs(); len(nodes) > 0 && !tuo.mutation.TagsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TagsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.TopicsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TopicsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.KudoedUsersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedKudoedUsersIDs(); len(nodes) > 0 && !tuo.mutation.KudoedUsersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.KudoedUsersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Thread{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{thread.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
