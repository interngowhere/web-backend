// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/interngowhere/web-backend/ent/moderator"
	"github.com/interngowhere/web-backend/ent/predicate"
)

// ModeratorUpdate is the builder for updating Moderator entities.
type ModeratorUpdate struct {
	config
	hooks    []Hook
	mutation *ModeratorMutation
}

// Where appends a list predicates to the ModeratorUpdate builder.
func (mu *ModeratorUpdate) Where(ps ...predicate.Moderator) *ModeratorUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// Mutation returns the ModeratorMutation object of the builder.
func (mu *ModeratorUpdate) Mutation() *ModeratorMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *ModeratorUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *ModeratorUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *ModeratorUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *ModeratorUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *ModeratorUpdate) check() error {
	if _, ok := mu.mutation.ModeratorID(); mu.mutation.ModeratorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Moderator.moderator"`)
	}
	if _, ok := mu.mutation.TopicID(); mu.mutation.TopicCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Moderator.topic"`)
	}
	return nil
}

func (mu *ModeratorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(moderator.Table, moderator.Columns, sqlgraph.NewFieldSpec(moderator.FieldModeratorID, field.TypeUUID), sqlgraph.NewFieldSpec(moderator.FieldTopicID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{moderator.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// ModeratorUpdateOne is the builder for updating a single Moderator entity.
type ModeratorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ModeratorMutation
}

// Mutation returns the ModeratorMutation object of the builder.
func (muo *ModeratorUpdateOne) Mutation() *ModeratorMutation {
	return muo.mutation
}

// Where appends a list predicates to the ModeratorUpdate builder.
func (muo *ModeratorUpdateOne) Where(ps ...predicate.Moderator) *ModeratorUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *ModeratorUpdateOne) Select(field string, fields ...string) *ModeratorUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Moderator entity.
func (muo *ModeratorUpdateOne) Save(ctx context.Context) (*Moderator, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *ModeratorUpdateOne) SaveX(ctx context.Context) *Moderator {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *ModeratorUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *ModeratorUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *ModeratorUpdateOne) check() error {
	if _, ok := muo.mutation.ModeratorID(); muo.mutation.ModeratorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Moderator.moderator"`)
	}
	if _, ok := muo.mutation.TopicID(); muo.mutation.TopicCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Moderator.topic"`)
	}
	return nil
}

func (muo *ModeratorUpdateOne) sqlSave(ctx context.Context) (_node *Moderator, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(moderator.Table, moderator.Columns, sqlgraph.NewFieldSpec(moderator.FieldModeratorID, field.TypeUUID), sqlgraph.NewFieldSpec(moderator.FieldTopicID, field.TypeInt))
	if id, ok := muo.mutation.ModeratorID(); !ok {
		return nil, &ValidationError{Name: "moderator_id", err: errors.New(`ent: missing "Moderator.moderator_id" for update`)}
	} else {
		_spec.Node.CompositeID[0].Value = id
	}
	if id, ok := muo.mutation.TopicID(); !ok {
		return nil, &ValidationError{Name: "topic_id", err: errors.New(`ent: missing "Moderator.topic_id" for update`)}
	} else {
		_spec.Node.CompositeID[1].Value = id
	}
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, len(fields))
		for i, f := range fields {
			if !moderator.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			_spec.Node.Columns[i] = f
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &Moderator{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{moderator.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
