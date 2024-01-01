// Code generated by ent, DO NOT EDIT.

package threadkudo

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/interngowhere/web-backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ThreadKudo {
	return predicate.ThreadKudo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ThreadKudo {
	return predicate.ThreadKudo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ThreadKudo {
	return predicate.ThreadKudo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ThreadKudo {
	return predicate.ThreadKudo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ThreadKudo {
	return predicate.ThreadKudo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ThreadKudo {
	return predicate.ThreadKudo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ThreadKudo {
	return predicate.ThreadKudo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ThreadKudo {
	return predicate.ThreadKudo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ThreadKudo {
	return predicate.ThreadKudo(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.ThreadKudo {
	return predicate.ThreadKudo(sql.FieldEQ(FieldUserID, v))
}

// ThreadID applies equality check predicate on the "thread_id" field. It's identical to ThreadIDEQ.
func ThreadID(v int) predicate.ThreadKudo {
	return predicate.ThreadKudo(sql.FieldEQ(FieldThreadID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.ThreadKudo {
	return predicate.ThreadKudo(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.ThreadKudo {
	return predicate.ThreadKudo(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.ThreadKudo {
	return predicate.ThreadKudo(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.ThreadKudo {
	return predicate.ThreadKudo(sql.FieldNotIn(FieldUserID, vs...))
}

// ThreadIDEQ applies the EQ predicate on the "thread_id" field.
func ThreadIDEQ(v int) predicate.ThreadKudo {
	return predicate.ThreadKudo(sql.FieldEQ(FieldThreadID, v))
}

// ThreadIDNEQ applies the NEQ predicate on the "thread_id" field.
func ThreadIDNEQ(v int) predicate.ThreadKudo {
	return predicate.ThreadKudo(sql.FieldNEQ(FieldThreadID, v))
}

// ThreadIDIn applies the In predicate on the "thread_id" field.
func ThreadIDIn(vs ...int) predicate.ThreadKudo {
	return predicate.ThreadKudo(sql.FieldIn(FieldThreadID, vs...))
}

// ThreadIDNotIn applies the NotIn predicate on the "thread_id" field.
func ThreadIDNotIn(vs ...int) predicate.ThreadKudo {
	return predicate.ThreadKudo(sql.FieldNotIn(FieldThreadID, vs...))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.ThreadKudo {
	return predicate.ThreadKudo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.ThreadKudo {
	return predicate.ThreadKudo(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasThread applies the HasEdge predicate on the "thread" edge.
func HasThread() predicate.ThreadKudo {
	return predicate.ThreadKudo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ThreadTable, ThreadColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasThreadWith applies the HasEdge predicate on the "thread" edge with a given conditions (other predicates).
func HasThreadWith(preds ...predicate.Thread) predicate.ThreadKudo {
	return predicate.ThreadKudo(func(s *sql.Selector) {
		step := newThreadStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ThreadKudo) predicate.ThreadKudo {
	return predicate.ThreadKudo(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ThreadKudo) predicate.ThreadKudo {
	return predicate.ThreadKudo(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ThreadKudo) predicate.ThreadKudo {
	return predicate.ThreadKudo(sql.NotPredicates(p))
}
