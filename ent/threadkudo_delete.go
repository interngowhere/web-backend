// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/interngowhere/web-backend/ent/predicate"
	"github.com/interngowhere/web-backend/ent/threadkudo"
)

// ThreadKudoDelete is the builder for deleting a ThreadKudo entity.
type ThreadKudoDelete struct {
	config
	hooks    []Hook
	mutation *ThreadKudoMutation
}

// Where appends a list predicates to the ThreadKudoDelete builder.
func (tkd *ThreadKudoDelete) Where(ps ...predicate.ThreadKudo) *ThreadKudoDelete {
	tkd.mutation.Where(ps...)
	return tkd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tkd *ThreadKudoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tkd.sqlExec, tkd.mutation, tkd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tkd *ThreadKudoDelete) ExecX(ctx context.Context) int {
	n, err := tkd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tkd *ThreadKudoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(threadkudo.Table, sqlgraph.NewFieldSpec(threadkudo.FieldID, field.TypeInt))
	if ps := tkd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tkd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tkd.mutation.done = true
	return affected, err
}

// ThreadKudoDeleteOne is the builder for deleting a single ThreadKudo entity.
type ThreadKudoDeleteOne struct {
	tkd *ThreadKudoDelete
}

// Where appends a list predicates to the ThreadKudoDelete builder.
func (tkdo *ThreadKudoDeleteOne) Where(ps ...predicate.ThreadKudo) *ThreadKudoDeleteOne {
	tkdo.tkd.mutation.Where(ps...)
	return tkdo
}

// Exec executes the deletion query.
func (tkdo *ThreadKudoDeleteOne) Exec(ctx context.Context) error {
	n, err := tkdo.tkd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{threadkudo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tkdo *ThreadKudoDeleteOne) ExecX(ctx context.Context) {
	if err := tkdo.Exec(ctx); err != nil {
		panic(err)
	}
}
