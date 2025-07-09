package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Mixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Unique(),
		field.String("account").
			Unique().
			Comment("账号"),
		field.String("password").
			Comment("密码"),
		field.String("name").
			Comment("昵称"),
		field.String("icon").
			Comment("头像"), //创建时给一个头像选择
		field.String("email").
			Optional().
			Nillable().
			Unique().
			Comment("邮箱"),
		field.String("sign").
			Optional().
			Nillable().
			Comment("签名"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("artilces", Article.Type),
		edge.To("tags", Tag.Type),
	}
}
