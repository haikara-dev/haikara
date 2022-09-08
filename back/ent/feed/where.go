// Code generated by ent, DO NOT EDIT.

package feed

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/cubdesign/dailyfj/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Contents applies equality check predicate on the "contents" field. It's identical to ContentsEQ.
func Contents(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContents), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Feed {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Feed {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Feed {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Feed {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// ContentsEQ applies the EQ predicate on the "contents" field.
func ContentsEQ(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContents), v))
	})
}

// ContentsNEQ applies the NEQ predicate on the "contents" field.
func ContentsNEQ(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContents), v))
	})
}

// ContentsIn applies the In predicate on the "contents" field.
func ContentsIn(vs ...string) predicate.Feed {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContents), v...))
	})
}

// ContentsNotIn applies the NotIn predicate on the "contents" field.
func ContentsNotIn(vs ...string) predicate.Feed {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContents), v...))
	})
}

// ContentsGT applies the GT predicate on the "contents" field.
func ContentsGT(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContents), v))
	})
}

// ContentsGTE applies the GTE predicate on the "contents" field.
func ContentsGTE(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContents), v))
	})
}

// ContentsLT applies the LT predicate on the "contents" field.
func ContentsLT(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContents), v))
	})
}

// ContentsLTE applies the LTE predicate on the "contents" field.
func ContentsLTE(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContents), v))
	})
}

// ContentsContains applies the Contains predicate on the "contents" field.
func ContentsContains(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContents), v))
	})
}

// ContentsHasPrefix applies the HasPrefix predicate on the "contents" field.
func ContentsHasPrefix(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContents), v))
	})
}

// ContentsHasSuffix applies the HasSuffix predicate on the "contents" field.
func ContentsHasSuffix(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContents), v))
	})
}

// ContentsEqualFold applies the EqualFold predicate on the "contents" field.
func ContentsEqualFold(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContents), v))
	})
}

// ContentsContainsFold applies the ContainsFold predicate on the "contents" field.
func ContentsContainsFold(v string) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContents), v))
	})
}

// HasSite applies the HasEdge predicate on the "site" edge.
func HasSite() predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SiteTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SiteTable, SiteColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSiteWith applies the HasEdge predicate on the "site" edge with a given conditions (other predicates).
func HasSiteWith(preds ...predicate.Site) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SiteInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SiteTable, SiteColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Feed) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Feed) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
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
func Not(p predicate.Feed) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		p(s.Not())
	})
}
