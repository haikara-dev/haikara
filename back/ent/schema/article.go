package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

func (Article) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("url").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(2083)",
			}).
			NotEmpty(),
		field.Time("published_at"),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ogp_image", OGPImage.Type).
			Unique(),
		edge.From("site", Site.Type).
			Ref("articles").
			Unique().
			Required(),
	}
}
