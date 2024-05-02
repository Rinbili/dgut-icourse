package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CourseComment holds the schema definition for the CourseComment entity.
type CourseComment struct {
	ent.Schema
}

// Fields of the CourseComment.
func (CourseComment) Fields() []ent.Field {
	return []ent.Field{
		field.Int8("difficulty").
			Comment("难度"),
		field.Int8("quality").
			Comment("收获"),
		field.Int8("workload").
			Comment("作业量"),
		field.Int8("mark").
			Comment("评分好坏"),
	}
}

// Edges of the CourseComment.
func (CourseComment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("comment", Comment.Type).
			Ref("course_comment").
			Unique().
			Required(),
	}
}
