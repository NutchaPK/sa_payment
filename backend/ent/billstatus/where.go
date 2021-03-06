// Code generated by entc, DO NOT EDIT.

package billstatus

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/theuo/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BillStatus applies equality check predicate on the "BillStatus" field. It's identical to BillStatusEQ.
func BillStatus(v string) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBillStatus), v))
	})
}

// BillStatusEQ applies the EQ predicate on the "BillStatus" field.
func BillStatusEQ(v string) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBillStatus), v))
	})
}

// BillStatusNEQ applies the NEQ predicate on the "BillStatus" field.
func BillStatusNEQ(v string) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBillStatus), v))
	})
}

// BillStatusIn applies the In predicate on the "BillStatus" field.
func BillStatusIn(vs ...string) predicate.BillStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillStatus(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBillStatus), v...))
	})
}

// BillStatusNotIn applies the NotIn predicate on the "BillStatus" field.
func BillStatusNotIn(vs ...string) predicate.BillStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillStatus(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBillStatus), v...))
	})
}

// BillStatusGT applies the GT predicate on the "BillStatus" field.
func BillStatusGT(v string) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBillStatus), v))
	})
}

// BillStatusGTE applies the GTE predicate on the "BillStatus" field.
func BillStatusGTE(v string) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBillStatus), v))
	})
}

// BillStatusLT applies the LT predicate on the "BillStatus" field.
func BillStatusLT(v string) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBillStatus), v))
	})
}

// BillStatusLTE applies the LTE predicate on the "BillStatus" field.
func BillStatusLTE(v string) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBillStatus), v))
	})
}

// BillStatusContains applies the Contains predicate on the "BillStatus" field.
func BillStatusContains(v string) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBillStatus), v))
	})
}

// BillStatusHasPrefix applies the HasPrefix predicate on the "BillStatus" field.
func BillStatusHasPrefix(v string) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBillStatus), v))
	})
}

// BillStatusHasSuffix applies the HasSuffix predicate on the "BillStatus" field.
func BillStatusHasSuffix(v string) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBillStatus), v))
	})
}

// BillStatusEqualFold applies the EqualFold predicate on the "BillStatus" field.
func BillStatusEqualFold(v string) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBillStatus), v))
	})
}

// BillStatusContainsFold applies the ContainsFold predicate on the "BillStatus" field.
func BillStatusContainsFold(v string) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBillStatus), v))
	})
}

// HasBills applies the HasEdge predicate on the "bills" edge.
func HasBills() predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BillsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BillsTable, BillsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBillsWith applies the HasEdge predicate on the "bills" edge with a given conditions (other predicates).
func HasBillsWith(preds ...predicate.Bill) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BillsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BillsTable, BillsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.BillStatus) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.BillStatus) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
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
func Not(p predicate.BillStatus) predicate.BillStatus {
	return predicate.BillStatus(func(s *sql.Selector) {
		p(s.Not())
	})
}
