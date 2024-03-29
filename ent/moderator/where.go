// Code generated by ent, DO NOT EDIT.

package moderator

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/interngowhere/web-backend/ent/predicate"
)

// ModeratorID applies equality check predicate on the "moderator_id" field. It's identical to ModeratorIDEQ.
func ModeratorID(v uuid.UUID) predicate.Moderator {
	return predicate.Moderator(sql.FieldEQ(FieldModeratorID, v))
}

// TopicID applies equality check predicate on the "topic_id" field. It's identical to TopicIDEQ.
func TopicID(v int) predicate.Moderator {
	return predicate.Moderator(sql.FieldEQ(FieldTopicID, v))
}

// ModeratorIDEQ applies the EQ predicate on the "moderator_id" field.
func ModeratorIDEQ(v uuid.UUID) predicate.Moderator {
	return predicate.Moderator(sql.FieldEQ(FieldModeratorID, v))
}

// ModeratorIDNEQ applies the NEQ predicate on the "moderator_id" field.
func ModeratorIDNEQ(v uuid.UUID) predicate.Moderator {
	return predicate.Moderator(sql.FieldNEQ(FieldModeratorID, v))
}

// ModeratorIDIn applies the In predicate on the "moderator_id" field.
func ModeratorIDIn(vs ...uuid.UUID) predicate.Moderator {
	return predicate.Moderator(sql.FieldIn(FieldModeratorID, vs...))
}

// ModeratorIDNotIn applies the NotIn predicate on the "moderator_id" field.
func ModeratorIDNotIn(vs ...uuid.UUID) predicate.Moderator {
	return predicate.Moderator(sql.FieldNotIn(FieldModeratorID, vs...))
}

// TopicIDEQ applies the EQ predicate on the "topic_id" field.
func TopicIDEQ(v int) predicate.Moderator {
	return predicate.Moderator(sql.FieldEQ(FieldTopicID, v))
}

// TopicIDNEQ applies the NEQ predicate on the "topic_id" field.
func TopicIDNEQ(v int) predicate.Moderator {
	return predicate.Moderator(sql.FieldNEQ(FieldTopicID, v))
}

// TopicIDIn applies the In predicate on the "topic_id" field.
func TopicIDIn(vs ...int) predicate.Moderator {
	return predicate.Moderator(sql.FieldIn(FieldTopicID, vs...))
}

// TopicIDNotIn applies the NotIn predicate on the "topic_id" field.
func TopicIDNotIn(vs ...int) predicate.Moderator {
	return predicate.Moderator(sql.FieldNotIn(FieldTopicID, vs...))
}

// HasModerator applies the HasEdge predicate on the "moderator" edge.
func HasModerator() predicate.Moderator {
	return predicate.Moderator(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, ModeratorColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, ModeratorTable, ModeratorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasModeratorWith applies the HasEdge predicate on the "moderator" edge with a given conditions (other predicates).
func HasModeratorWith(preds ...predicate.User) predicate.Moderator {
	return predicate.Moderator(func(s *sql.Selector) {
		step := newModeratorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTopic applies the HasEdge predicate on the "topic" edge.
func HasTopic() predicate.Moderator {
	return predicate.Moderator(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, TopicColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, TopicTable, TopicColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTopicWith applies the HasEdge predicate on the "topic" edge with a given conditions (other predicates).
func HasTopicWith(preds ...predicate.Topic) predicate.Moderator {
	return predicate.Moderator(func(s *sql.Selector) {
		step := newTopicStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Moderator) predicate.Moderator {
	return predicate.Moderator(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Moderator) predicate.Moderator {
	return predicate.Moderator(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Moderator) predicate.Moderator {
	return predicate.Moderator(sql.NotPredicates(p))
}
