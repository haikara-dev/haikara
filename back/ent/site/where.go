// Code generated by ent, DO NOT EDIT.

package site

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/haikara-dev/haikara/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// FeedURL applies equality check predicate on the "feed_url" field. It's identical to FeedURLEQ.
func FeedURL(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFeedURL), v))
	})
}

// Active applies equality check predicate on the "active" field. It's identical to ActiveEQ.
func Active(v bool) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActive), v))
	})
}

// CannotCrawlAt applies equality check predicate on the "cannot_crawl_at" field. It's identical to CannotCrawlAtEQ.
func CannotCrawlAt(v time.Time) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCannotCrawlAt), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Site {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Site {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Site {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Site {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Site {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Site {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldURL), v))
	})
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Site {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldURL), v...))
	})
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Site {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldURL), v...))
	})
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldURL), v))
	})
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldURL), v))
	})
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldURL), v))
	})
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldURL), v))
	})
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldURL), v))
	})
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldURL), v))
	})
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldURL), v))
	})
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldURL), v))
	})
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldURL), v))
	})
}

// FeedURLEQ applies the EQ predicate on the "feed_url" field.
func FeedURLEQ(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFeedURL), v))
	})
}

// FeedURLNEQ applies the NEQ predicate on the "feed_url" field.
func FeedURLNEQ(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFeedURL), v))
	})
}

// FeedURLIn applies the In predicate on the "feed_url" field.
func FeedURLIn(vs ...string) predicate.Site {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFeedURL), v...))
	})
}

// FeedURLNotIn applies the NotIn predicate on the "feed_url" field.
func FeedURLNotIn(vs ...string) predicate.Site {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFeedURL), v...))
	})
}

// FeedURLGT applies the GT predicate on the "feed_url" field.
func FeedURLGT(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFeedURL), v))
	})
}

// FeedURLGTE applies the GTE predicate on the "feed_url" field.
func FeedURLGTE(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFeedURL), v))
	})
}

// FeedURLLT applies the LT predicate on the "feed_url" field.
func FeedURLLT(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFeedURL), v))
	})
}

// FeedURLLTE applies the LTE predicate on the "feed_url" field.
func FeedURLLTE(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFeedURL), v))
	})
}

// FeedURLContains applies the Contains predicate on the "feed_url" field.
func FeedURLContains(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFeedURL), v))
	})
}

// FeedURLHasPrefix applies the HasPrefix predicate on the "feed_url" field.
func FeedURLHasPrefix(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFeedURL), v))
	})
}

// FeedURLHasSuffix applies the HasSuffix predicate on the "feed_url" field.
func FeedURLHasSuffix(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFeedURL), v))
	})
}

// FeedURLEqualFold applies the EqualFold predicate on the "feed_url" field.
func FeedURLEqualFold(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFeedURL), v))
	})
}

// FeedURLContainsFold applies the ContainsFold predicate on the "feed_url" field.
func FeedURLContainsFold(v string) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFeedURL), v))
	})
}

// ActiveEQ applies the EQ predicate on the "active" field.
func ActiveEQ(v bool) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActive), v))
	})
}

// ActiveNEQ applies the NEQ predicate on the "active" field.
func ActiveNEQ(v bool) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldActive), v))
	})
}

// CannotCrawlAtEQ applies the EQ predicate on the "cannot_crawl_at" field.
func CannotCrawlAtEQ(v time.Time) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCannotCrawlAt), v))
	})
}

// CannotCrawlAtNEQ applies the NEQ predicate on the "cannot_crawl_at" field.
func CannotCrawlAtNEQ(v time.Time) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCannotCrawlAt), v))
	})
}

// CannotCrawlAtIn applies the In predicate on the "cannot_crawl_at" field.
func CannotCrawlAtIn(vs ...time.Time) predicate.Site {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCannotCrawlAt), v...))
	})
}

// CannotCrawlAtNotIn applies the NotIn predicate on the "cannot_crawl_at" field.
func CannotCrawlAtNotIn(vs ...time.Time) predicate.Site {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCannotCrawlAt), v...))
	})
}

// CannotCrawlAtGT applies the GT predicate on the "cannot_crawl_at" field.
func CannotCrawlAtGT(v time.Time) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCannotCrawlAt), v))
	})
}

// CannotCrawlAtGTE applies the GTE predicate on the "cannot_crawl_at" field.
func CannotCrawlAtGTE(v time.Time) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCannotCrawlAt), v))
	})
}

// CannotCrawlAtLT applies the LT predicate on the "cannot_crawl_at" field.
func CannotCrawlAtLT(v time.Time) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCannotCrawlAt), v))
	})
}

// CannotCrawlAtLTE applies the LTE predicate on the "cannot_crawl_at" field.
func CannotCrawlAtLTE(v time.Time) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCannotCrawlAt), v))
	})
}

// CannotCrawlAtIsNil applies the IsNil predicate on the "cannot_crawl_at" field.
func CannotCrawlAtIsNil() predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCannotCrawlAt)))
	})
}

// CannotCrawlAtNotNil applies the NotNil predicate on the "cannot_crawl_at" field.
func CannotCrawlAtNotNil() predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCannotCrawlAt)))
	})
}

// HasArticles applies the HasEdge predicate on the "articles" edge.
func HasArticles() predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ArticlesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ArticlesTable, ArticlesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasArticlesWith applies the HasEdge predicate on the "articles" edge with a given conditions (other predicates).
func HasArticlesWith(preds ...predicate.Article) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ArticlesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ArticlesTable, ArticlesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFeeds applies the HasEdge predicate on the "feeds" edge.
func HasFeeds() predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FeedsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FeedsTable, FeedsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFeedsWith applies the HasEdge predicate on the "feeds" edge with a given conditions (other predicates).
func HasFeedsWith(preds ...predicate.Feed) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FeedsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FeedsTable, FeedsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSiteCrawlRule applies the HasEdge predicate on the "site_crawl_rule" edge.
func HasSiteCrawlRule() predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SiteCrawlRuleTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, SiteCrawlRuleTable, SiteCrawlRuleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSiteCrawlRuleWith applies the HasEdge predicate on the "site_crawl_rule" edge with a given conditions (other predicates).
func HasSiteCrawlRuleWith(preds ...predicate.SiteCrawlRule) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SiteCrawlRuleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, SiteCrawlRuleTable, SiteCrawlRuleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSiteCategories applies the HasEdge predicate on the "site_categories" edge.
func HasSiteCategories() predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SiteCategoriesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, SiteCategoriesTable, SiteCategoriesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSiteCategoriesWith applies the HasEdge predicate on the "site_categories" edge with a given conditions (other predicates).
func HasSiteCategoriesWith(preds ...predicate.SiteCategory) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SiteCategoriesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, SiteCategoriesTable, SiteCategoriesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Site) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Site) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
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
func Not(p predicate.Site) predicate.Site {
	return predicate.Site(func(s *sql.Selector) {
		p(s.Not())
	})
}
