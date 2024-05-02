package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Object holds the schema definition for the Object entity.
type Object struct {
	ent.Schema
}

// Fields of the Object.
func (Object) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			StructTag(`json:"id"`),
		field.String("name").
			MaxLen(30).
			Comment("名称"),
		field.String("desc").
			MaxLen(255).
			Comment("描述"),
		field.Int8("type").
			Comment("类型"),
	}
}

// Edges of the Object.
func (Object) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("course", Course.Type).Unique(),
		edge.To("organization", Organization.Type).Unique(),
		edge.To("person", Person.Type).Unique(),
		edge.To("other", Other.Type).Unique(),
		edge.To("comments", Comment.Type),
		edge.To("teach_courses", Course.Type),
		edge.To("setup_courses", Course.Type),
	}
}

func (Object) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
