package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			StructTag(`json:"id"`),
		field.Text("nickName").
			Optional().
			Nillable().
			MaxLen(30).
			Comment("用户名"),
		field.Text("password").
			Optional().
			Nillable().
			Sensitive().
			Comment("密码"),
		field.Text("phone").
			MinLen(5).
			MaxLen(15).
			Optional().
			Nillable().
			Unique().
			Comment("手机号码"),
		field.Text("icon").
			Optional().
			Nillable().
			Comment("头像URL"),
		field.Text("openid").
			MaxLen(30).
			Comment("微信openid"),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("info", UserInfo.Type).
			Unique(),
		edge.From("comments", Comment.Type).
			Ref("author").
			Comment("评论"),
		edge.To("liked_comments", Comment.Type),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
