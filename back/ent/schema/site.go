package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Site holds the schema definition for the Site entity.
type Site struct {
	ent.Schema
}

func (Site) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Site.
func (Site) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("url").NotEmpty(),
		field.String("feed_url"),
		field.Bool("active").Default(false).
			StructTag(`json:"active"`),
		field.Time("cannot_crawl_at").
			Optional().
			Nillable(),
	}
}

// Edges of the Site.
func (Site) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("articles", Article.Type),
		edge.To("feeds", Feed.Type),
		edge.To("site_crawl_rule", SiteCrawlRule.Type).Unique(),
		edge.From("site_categories", SiteCategory.Type).
			Ref("sites"),
	}
}
