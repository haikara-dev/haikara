// Code generated by ent, DO NOT EDIT.

package sitecategory

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/haikara-dev/haikara/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SiteCategory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SiteCategory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SiteCategory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SiteCategory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// LabelEQ applies the EQ predicate on the "label" field.
func LabelEQ(v string) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLabel), v))
	})
}

// LabelNEQ applies the NEQ predicate on the "label" field.
func LabelNEQ(v string) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLabel), v))
	})
}

// LabelIn applies the In predicate on the "label" field.
func LabelIn(vs ...string) predicate.SiteCategory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLabel), v...))
	})
}

// LabelNotIn applies the NotIn predicate on the "label" field.
func LabelNotIn(vs ...string) predicate.SiteCategory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLabel), v...))
	})
}

// LabelGT applies the GT predicate on the "label" field.
func LabelGT(v string) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLabel), v))
	})
}

// LabelGTE applies the GTE predicate on the "label" field.
func LabelGTE(v string) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLabel), v))
	})
}

// LabelLT applies the LT predicate on the "label" field.
func LabelLT(v string) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLabel), v))
	})
}

// LabelLTE applies the LTE predicate on the "label" field.
func LabelLTE(v string) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLabel), v))
	})
}

// LabelContains applies the Contains predicate on the "label" field.
func LabelContains(v string) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLabel), v))
	})
}

// LabelHasPrefix applies the HasPrefix predicate on the "label" field.
func LabelHasPrefix(v string) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLabel), v))
	})
}

// LabelHasSuffix applies the HasSuffix predicate on the "label" field.
func LabelHasSuffix(v string) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLabel), v))
	})
}

// LabelEqualFold applies the EqualFold predicate on the "label" field.
func LabelEqualFold(v string) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLabel), v))
	})
}

// LabelContainsFold applies the ContainsFold predicate on the "label" field.
func LabelContainsFold(v string) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLabel), v))
	})
}

// HasSites applies the HasEdge predicate on the "sites" edge.
func HasSites() predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SitesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, SitesTable, SitesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSitesWith applies the HasEdge predicate on the "sites" edge with a given conditions (other predicates).
func HasSitesWith(preds ...predicate.Site) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SitesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, SitesTable, SitesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SiteCategory) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SiteCategory) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
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
func Not(p predicate.SiteCategory) predicate.SiteCategory {
	return predicate.SiteCategory(func(s *sql.Selector) {
		p(s.Not())
	})
}
