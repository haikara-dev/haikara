package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SiteCrawlRule holds the schema definition for the SiteCrawlRule entity.
type SiteCrawlRule struct {
	ent.Schema
}

func (SiteCrawlRule) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the SiteCrawlRule.
func (SiteCrawlRule) Fields() []ent.Field {
	return []ent.Field{
		field.String("article_selector"),
		field.String("title_selector"),
		field.String("link_selector"),
		field.String("description_selector"),
		field.Bool("has_data_to_list").Default(true).
			StructTag(`json:"has_data_to_list"`),
		field.String("date_selector"),
		field.String("date_layout"),
		field.Bool("is_time_humanize").Default(false).
			StructTag(`json:"is_time_humanize"`),
		field.Bool("is_spa").Default(false).
			StructTag(`json:"is_spa"`),
	}
}

// Edges of the SiteCrawlRule.
func (SiteCrawlRule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("site", Site.Type).
			Ref("site_crawl_rule").
			Unique().
			Required(),
	}
}
