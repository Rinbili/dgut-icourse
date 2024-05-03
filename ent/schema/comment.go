package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			StructTag(`json:"id"`),
		field.String("content").
			MaxLen(255).
			Comment("评论内容"),
		field.Int("score").
			Comment("评分"),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("author", User.Type).
			Required().
			Comment("评论作者").
			Unique(),
		edge.From("object", Object.Type).
			Ref("comments").
			Unique(),
		edge.To("children", Comment.Type).
			From("parent").
			Unique(),
		edge.To("course_comment", CourseComment.Type).
			Unique(),
		edge.From("liked_users", User.Type).
			Ref("liked_comments"),
	}
}

// Mixin of the Comment.
func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
