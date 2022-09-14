package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SiteCategory holds the schema definition for the SiteCategory entity.
type SiteCategory struct {
	ent.Schema
}

func (SiteCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the SiteCategory.
func (SiteCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("label").Unique().NotEmpty(),
	}
}

// Edges of the SiteCategory.
func (SiteCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sites", Site.Type),
	}
}
