package schema

import (
	"entgo.io/ent"
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
		field.Bool("active").Default(false).
			StructTag(`json:"active"`),
	}
}

// Edges of the Site.
func (Site) Edges() []ent.Edge {
	return nil
}
