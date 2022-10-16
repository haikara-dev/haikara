package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OGPImage holds the schema definition for the OGPImage entity.
type OGPImage struct {
	ent.Schema
}

func (OGPImage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the OGPImage.
func (OGPImage) Fields() []ent.Field {
	return []ent.Field{
		field.String("file_name"),
		field.String("file_path"),
		field.String("origin_url").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(2083)",
			}),
	}
}

// Edges of the OGPImage.
func (OGPImage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("article", Article.Type).
			Ref("ogp_image").
			Unique().
			Required(),
	}
}
