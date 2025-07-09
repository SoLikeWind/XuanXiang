package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

func (Article) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Mixin{},
	}
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Unique(),
		field.String("title").
			MinLen(0).
			MaxLen(255),
		field.String("summary").
			MinLen(0).
			MaxLen(255),
		field.String("image").
			Optional().
			MaxLen(255),
		field.String("content_md").
			MinLen(0).
			MaxLen(100000),
		field.String("content_html").
			MinLen(0).
			MaxLen(100000),
		field.Int64("views").
			Default(0).
			Min(0),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type),
	}
}
