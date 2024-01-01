// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/interngowhere/web-backend/ent/predicate"
	"github.com/interngowhere/web-backend/ent/thread"
	"github.com/interngowhere/web-backend/ent/topic"
)

// TopicUpdate is the builder for updating Topic entities.
type TopicUpdate struct {
	config
	hooks    []Hook
	mutation *TopicMutation
}

// Where appends a list predicates to the TopicUpdate builder.
func (tu *TopicUpdate) Where(ps ...predicate.Topic) *TopicUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetTitle sets the "title" field.
func (tu *TopicUpdate) SetTitle(s string) *TopicUpdate {
	tu.mutation.SetTitle(s)
	return tu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (tu *TopicUpdate) SetNillableTitle(s *string) *TopicUpdate {
	if s != nil {
		tu.SetTitle(*s)
	}
	return tu
}

// SetShortTitle sets the "short_title" field.
func (tu *TopicUpdate) SetShortTitle(s string) *TopicUpdate {
	tu.mutation.SetShortTitle(s)
	return tu
}

// SetNillableShortTitle sets the "short_title" field if the given value is not nil.
func (tu *TopicUpdate) SetNillableShortTitle(s *string) *TopicUpdate {
	if s != nil {
		tu.SetShortTitle(*s)
	}
	return tu
}

// SetDescription sets the "description" field.
func (tu *TopicUpdate) SetDescription(s string) *TopicUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tu *TopicUpdate) SetNillableDescription(s *string) *TopicUpdate {
	if s != nil {
		tu.SetDescription(*s)
	}
	return tu
}

// ClearDescription clears the value of the "description" field.
func (tu *TopicUpdate) ClearDescription() *TopicUpdate {
	tu.mutation.ClearDescription()
	return tu
}

// SetModerators sets the "moderators" field.
func (tu *TopicUpdate) SetModerators(s []string) *TopicUpdate {
	tu.mutation.SetModerators(s)
	return tu
}

// AppendModerators appends s to the "moderators" field.
func (tu *TopicUpdate) AppendModerators(s []string) *TopicUpdate {
	tu.mutation.AppendModerators(s)
	return tu
}

// SetCreatedBy sets the "created_by" field.
func (tu *TopicUpdate) SetCreatedBy(s string) *TopicUpdate {
	tu.mutation.SetCreatedBy(s)
	return tu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (tu *TopicUpdate) SetNillableCreatedBy(s *string) *TopicUpdate {
	if s != nil {
		tu.SetCreatedBy(*s)
	}
	return tu
}

// AddTopicThreadIDs adds the "topic_threads" edge to the Thread entity by IDs.
func (tu *TopicUpdate) AddTopicThreadIDs(ids ...int) *TopicUpdate {
	tu.mutation.AddTopicThreadIDs(ids...)
	return tu
}

// AddTopicThreads adds the "topic_threads" edges to the Thread entity.
func (tu *TopicUpdate) AddTopicThreads(t ...*Thread) *TopicUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTopicThreadIDs(ids...)
}

// Mutation returns the TopicMutation object of the builder.
func (tu *TopicUpdate) Mutation() *TopicMutation {
	return tu.mutation
}

// ClearTopicThreads clears all "topic_threads" edges to the Thread entity.
func (tu *TopicUpdate) ClearTopicThreads() *TopicUpdate {
	tu.mutation.ClearTopicThreads()
	return tu
}

// RemoveTopicThreadIDs removes the "topic_threads" edge to Thread entities by IDs.
func (tu *TopicUpdate) RemoveTopicThreadIDs(ids ...int) *TopicUpdate {
	tu.mutation.RemoveTopicThreadIDs(ids...)
	return tu
}

// RemoveTopicThreads removes "topic_threads" edges to Thread entities.
func (tu *TopicUpdate) RemoveTopicThreads(t ...*Thread) *TopicUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTopicThreadIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TopicUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TopicUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TopicUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TopicUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TopicUpdate) check() error {
	if v, ok := tu.mutation.Title(); ok {
		if err := topic.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Topic.title": %w`, err)}
		}
	}
	if v, ok := tu.mutation.ShortTitle(); ok {
		if err := topic.ShortTitleValidator(v); err != nil {
			return &ValidationError{Name: "short_title", err: fmt.Errorf(`ent: validator failed for field "Topic.short_title": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Description(); ok {
		if err := topic.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Topic.description": %w`, err)}
		}
	}
	if v, ok := tu.mutation.CreatedBy(); ok {
		if err := topic.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`ent: validator failed for field "Topic.created_by": %w`, err)}
		}
	}
	return nil
}

func (tu *TopicUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(topic.Table, topic.Columns, sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Title(); ok {
		_spec.SetField(topic.FieldTitle, field.TypeString, value)
	}
	if value, ok := tu.mutation.ShortTitle(); ok {
		_spec.SetField(topic.FieldShortTitle, field.TypeString, value)
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.SetField(topic.FieldDescription, field.TypeString, value)
	}
	if tu.mutation.DescriptionCleared() {
		_spec.ClearField(topic.FieldDescription, field.TypeString)
	}
	if value, ok := tu.mutation.Moderators(); ok {
		_spec.SetField(topic.FieldModerators, field.TypeJSON, value)
	}
	if value, ok := tu.mutation.AppendedModerators(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, topic.FieldModerators, value)
		})
	}
	if value, ok := tu.mutation.CreatedBy(); ok {
		_spec.SetField(topic.FieldCreatedBy, field.TypeString, value)
	}
	if tu.mutation.TopicThreadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.TopicThreadsTable,
			Columns: []string{topic.TopicThreadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedTopicThreadsIDs(); len(nodes) > 0 && !tu.mutation.TopicThreadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.TopicThreadsTable,
			Columns: []string{topic.TopicThreadsColumn},
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
	if nodes := tu.mutation.TopicThreadsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.TopicThreadsTable,
			Columns: []string{topic.TopicThreadsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{topic.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TopicUpdateOne is the builder for updating a single Topic entity.
type TopicUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TopicMutation
}

// SetTitle sets the "title" field.
func (tuo *TopicUpdateOne) SetTitle(s string) *TopicUpdateOne {
	tuo.mutation.SetTitle(s)
	return tuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (tuo *TopicUpdateOne) SetNillableTitle(s *string) *TopicUpdateOne {
	if s != nil {
		tuo.SetTitle(*s)
	}
	return tuo
}

// SetShortTitle sets the "short_title" field.
func (tuo *TopicUpdateOne) SetShortTitle(s string) *TopicUpdateOne {
	tuo.mutation.SetShortTitle(s)
	return tuo
}

// SetNillableShortTitle sets the "short_title" field if the given value is not nil.
func (tuo *TopicUpdateOne) SetNillableShortTitle(s *string) *TopicUpdateOne {
	if s != nil {
		tuo.SetShortTitle(*s)
	}
	return tuo
}

// SetDescription sets the "description" field.
func (tuo *TopicUpdateOne) SetDescription(s string) *TopicUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tuo *TopicUpdateOne) SetNillableDescription(s *string) *TopicUpdateOne {
	if s != nil {
		tuo.SetDescription(*s)
	}
	return tuo
}

// ClearDescription clears the value of the "description" field.
func (tuo *TopicUpdateOne) ClearDescription() *TopicUpdateOne {
	tuo.mutation.ClearDescription()
	return tuo
}

// SetModerators sets the "moderators" field.
func (tuo *TopicUpdateOne) SetModerators(s []string) *TopicUpdateOne {
	tuo.mutation.SetModerators(s)
	return tuo
}

// AppendModerators appends s to the "moderators" field.
func (tuo *TopicUpdateOne) AppendModerators(s []string) *TopicUpdateOne {
	tuo.mutation.AppendModerators(s)
	return tuo
}

// SetCreatedBy sets the "created_by" field.
func (tuo *TopicUpdateOne) SetCreatedBy(s string) *TopicUpdateOne {
	tuo.mutation.SetCreatedBy(s)
	return tuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (tuo *TopicUpdateOne) SetNillableCreatedBy(s *string) *TopicUpdateOne {
	if s != nil {
		tuo.SetCreatedBy(*s)
	}
	return tuo
}

// AddTopicThreadIDs adds the "topic_threads" edge to the Thread entity by IDs.
func (tuo *TopicUpdateOne) AddTopicThreadIDs(ids ...int) *TopicUpdateOne {
	tuo.mutation.AddTopicThreadIDs(ids...)
	return tuo
}

// AddTopicThreads adds the "topic_threads" edges to the Thread entity.
func (tuo *TopicUpdateOne) AddTopicThreads(t ...*Thread) *TopicUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTopicThreadIDs(ids...)
}

// Mutation returns the TopicMutation object of the builder.
func (tuo *TopicUpdateOne) Mutation() *TopicMutation {
	return tuo.mutation
}

// ClearTopicThreads clears all "topic_threads" edges to the Thread entity.
func (tuo *TopicUpdateOne) ClearTopicThreads() *TopicUpdateOne {
	tuo.mutation.ClearTopicThreads()
	return tuo
}

// RemoveTopicThreadIDs removes the "topic_threads" edge to Thread entities by IDs.
func (tuo *TopicUpdateOne) RemoveTopicThreadIDs(ids ...int) *TopicUpdateOne {
	tuo.mutation.RemoveTopicThreadIDs(ids...)
	return tuo
}

// RemoveTopicThreads removes "topic_threads" edges to Thread entities.
func (tuo *TopicUpdateOne) RemoveTopicThreads(t ...*Thread) *TopicUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTopicThreadIDs(ids...)
}

// Where appends a list predicates to the TopicUpdate builder.
func (tuo *TopicUpdateOne) Where(ps ...predicate.Topic) *TopicUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TopicUpdateOne) Select(field string, fields ...string) *TopicUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Topic entity.
func (tuo *TopicUpdateOne) Save(ctx context.Context) (*Topic, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TopicUpdateOne) SaveX(ctx context.Context) *Topic {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TopicUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TopicUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TopicUpdateOne) check() error {
	if v, ok := tuo.mutation.Title(); ok {
		if err := topic.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Topic.title": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.ShortTitle(); ok {
		if err := topic.ShortTitleValidator(v); err != nil {
			return &ValidationError{Name: "short_title", err: fmt.Errorf(`ent: validator failed for field "Topic.short_title": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Description(); ok {
		if err := topic.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Topic.description": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.CreatedBy(); ok {
		if err := topic.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`ent: validator failed for field "Topic.created_by": %w`, err)}
		}
	}
	return nil
}

func (tuo *TopicUpdateOne) sqlSave(ctx context.Context) (_node *Topic, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(topic.Table, topic.Columns, sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Topic.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, topic.FieldID)
		for _, f := range fields {
			if !topic.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != topic.FieldID {
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
		_spec.SetField(topic.FieldTitle, field.TypeString, value)
	}
	if value, ok := tuo.mutation.ShortTitle(); ok {
		_spec.SetField(topic.FieldShortTitle, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.SetField(topic.FieldDescription, field.TypeString, value)
	}
	if tuo.mutation.DescriptionCleared() {
		_spec.ClearField(topic.FieldDescription, field.TypeString)
	}
	if value, ok := tuo.mutation.Moderators(); ok {
		_spec.SetField(topic.FieldModerators, field.TypeJSON, value)
	}
	if value, ok := tuo.mutation.AppendedModerators(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, topic.FieldModerators, value)
		})
	}
	if value, ok := tuo.mutation.CreatedBy(); ok {
		_spec.SetField(topic.FieldCreatedBy, field.TypeString, value)
	}
	if tuo.mutation.TopicThreadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.TopicThreadsTable,
			Columns: []string{topic.TopicThreadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedTopicThreadsIDs(); len(nodes) > 0 && !tuo.mutation.TopicThreadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.TopicThreadsTable,
			Columns: []string{topic.TopicThreadsColumn},
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
	if nodes := tuo.mutation.TopicThreadsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   topic.TopicThreadsTable,
			Columns: []string{topic.TopicThreadsColumn},
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
	_node = &Topic{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{topic.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
