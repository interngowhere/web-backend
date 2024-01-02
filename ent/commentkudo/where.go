// Code generated by ent, DO NOT EDIT.

package commentkudo

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/interngowhere/web-backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CommentKudo {
	return predicate.CommentKudo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CommentKudo {
	return predicate.CommentKudo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CommentKudo {
	return predicate.CommentKudo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CommentKudo {
	return predicate.CommentKudo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CommentKudo {
	return predicate.CommentKudo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CommentKudo {
	return predicate.CommentKudo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CommentKudo {
	return predicate.CommentKudo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CommentKudo {
	return predicate.CommentKudo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CommentKudo {
	return predicate.CommentKudo(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.CommentKudo {
	return predicate.CommentKudo(sql.FieldEQ(FieldUserID, v))
}

// CommentID applies equality check predicate on the "comment_id" field. It's identical to CommentIDEQ.
func CommentID(v int) predicate.CommentKudo {
	return predicate.CommentKudo(sql.FieldEQ(FieldCommentID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.CommentKudo {
	return predicate.CommentKudo(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.CommentKudo {
	return predicate.CommentKudo(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.CommentKudo {
	return predicate.CommentKudo(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.CommentKudo {
	return predicate.CommentKudo(sql.FieldNotIn(FieldUserID, vs...))
}

// CommentIDEQ applies the EQ predicate on the "comment_id" field.
func CommentIDEQ(v int) predicate.CommentKudo {
	return predicate.CommentKudo(sql.FieldEQ(FieldCommentID, v))
}

// CommentIDNEQ applies the NEQ predicate on the "comment_id" field.
func CommentIDNEQ(v int) predicate.CommentKudo {
	return predicate.CommentKudo(sql.FieldNEQ(FieldCommentID, v))
}

// CommentIDIn applies the In predicate on the "comment_id" field.
func CommentIDIn(vs ...int) predicate.CommentKudo {
	return predicate.CommentKudo(sql.FieldIn(FieldCommentID, vs...))
}

// CommentIDNotIn applies the NotIn predicate on the "comment_id" field.
func CommentIDNotIn(vs ...int) predicate.CommentKudo {
	return predicate.CommentKudo(sql.FieldNotIn(FieldCommentID, vs...))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.CommentKudo {
	return predicate.CommentKudo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.CommentKudo {
	return predicate.CommentKudo(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasComment applies the HasEdge predicate on the "comment" edge.
func HasComment() predicate.CommentKudo {
	return predicate.CommentKudo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CommentTable, CommentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentWith applies the HasEdge predicate on the "comment" edge with a given conditions (other predicates).
func HasCommentWith(preds ...predicate.Comment) predicate.CommentKudo {
	return predicate.CommentKudo(func(s *sql.Selector) {
		step := newCommentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CommentKudo) predicate.CommentKudo {
	return predicate.CommentKudo(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CommentKudo) predicate.CommentKudo {
	return predicate.CommentKudo(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CommentKudo) predicate.CommentKudo {
	return predicate.CommentKudo(sql.NotPredicates(p))
}