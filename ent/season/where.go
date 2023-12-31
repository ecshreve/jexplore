// Code generated by ent, DO NOT EDIT.

package season

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ecshreve/jexplore/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Season {
	return predicate.Season(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Season {
	return predicate.Season(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Season {
	return predicate.Season(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Season {
	return predicate.Season(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Season {
	return predicate.Season(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Season {
	return predicate.Season(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Season {
	return predicate.Season(sql.FieldLTE(FieldID, id))
}

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v int) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldNumber, v))
}

// StartDate applies equality check predicate on the "start_date" field. It's identical to StartDateEQ.
func StartDate(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldStartDate, v))
}

// EndDate applies equality check predicate on the "end_date" field. It's identical to EndDateEQ.
func EndDate(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldEndDate, v))
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v int) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldNumber, v))
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v int) predicate.Season {
	return predicate.Season(sql.FieldNEQ(FieldNumber, v))
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...int) predicate.Season {
	return predicate.Season(sql.FieldIn(FieldNumber, vs...))
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...int) predicate.Season {
	return predicate.Season(sql.FieldNotIn(FieldNumber, vs...))
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v int) predicate.Season {
	return predicate.Season(sql.FieldGT(FieldNumber, v))
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v int) predicate.Season {
	return predicate.Season(sql.FieldGTE(FieldNumber, v))
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v int) predicate.Season {
	return predicate.Season(sql.FieldLT(FieldNumber, v))
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v int) predicate.Season {
	return predicate.Season(sql.FieldLTE(FieldNumber, v))
}

// StartDateEQ applies the EQ predicate on the "start_date" field.
func StartDateEQ(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldStartDate, v))
}

// StartDateNEQ applies the NEQ predicate on the "start_date" field.
func StartDateNEQ(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldNEQ(FieldStartDate, v))
}

// StartDateIn applies the In predicate on the "start_date" field.
func StartDateIn(vs ...time.Time) predicate.Season {
	return predicate.Season(sql.FieldIn(FieldStartDate, vs...))
}

// StartDateNotIn applies the NotIn predicate on the "start_date" field.
func StartDateNotIn(vs ...time.Time) predicate.Season {
	return predicate.Season(sql.FieldNotIn(FieldStartDate, vs...))
}

// StartDateGT applies the GT predicate on the "start_date" field.
func StartDateGT(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldGT(FieldStartDate, v))
}

// StartDateGTE applies the GTE predicate on the "start_date" field.
func StartDateGTE(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldGTE(FieldStartDate, v))
}

// StartDateLT applies the LT predicate on the "start_date" field.
func StartDateLT(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldLT(FieldStartDate, v))
}

// StartDateLTE applies the LTE predicate on the "start_date" field.
func StartDateLTE(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldLTE(FieldStartDate, v))
}

// EndDateEQ applies the EQ predicate on the "end_date" field.
func EndDateEQ(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldEndDate, v))
}

// EndDateNEQ applies the NEQ predicate on the "end_date" field.
func EndDateNEQ(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldNEQ(FieldEndDate, v))
}

// EndDateIn applies the In predicate on the "end_date" field.
func EndDateIn(vs ...time.Time) predicate.Season {
	return predicate.Season(sql.FieldIn(FieldEndDate, vs...))
}

// EndDateNotIn applies the NotIn predicate on the "end_date" field.
func EndDateNotIn(vs ...time.Time) predicate.Season {
	return predicate.Season(sql.FieldNotIn(FieldEndDate, vs...))
}

// EndDateGT applies the GT predicate on the "end_date" field.
func EndDateGT(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldGT(FieldEndDate, v))
}

// EndDateGTE applies the GTE predicate on the "end_date" field.
func EndDateGTE(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldGTE(FieldEndDate, v))
}

// EndDateLT applies the LT predicate on the "end_date" field.
func EndDateLT(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldLT(FieldEndDate, v))
}

// EndDateLTE applies the LTE predicate on the "end_date" field.
func EndDateLTE(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldLTE(FieldEndDate, v))
}

// HasGames applies the HasEdge predicate on the "games" edge.
func HasGames() predicate.Season {
	return predicate.Season(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, GamesTable, GamesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGamesWith applies the HasEdge predicate on the "games" edge with a given conditions (other predicates).
func HasGamesWith(preds ...predicate.Game) predicate.Season {
	return predicate.Season(func(s *sql.Selector) {
		step := newGamesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Season) predicate.Season {
	return predicate.Season(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Season) predicate.Season {
	return predicate.Season(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Season) predicate.Season {
	return predicate.Season(func(s *sql.Selector) {
		p(s.Not())
	})
}
