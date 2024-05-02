package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Course holds the schema definition for the Course entity.
type Course struct {
	ent.Schema
}

// Fields of the Course.
func (Course) Fields() []ent.Field {
	return []ent.Field{
		field.String("courseID").
			Comment("课程ID"),
	}
}

// Edges of the Course.
func (Course) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("object", Object.Type).
			Ref("course").
			Unique().
			Required(),
		edge.From("teacher", Object.Type).
			Ref("teach_courses").
			Unique().
			Comment("授课教师"),
		edge.From("organization", Object.Type).
			Ref("setup_courses").
			Unique().
			Comment("开课单位"),
	}
}

// Mixin of the Course.
func (Course) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
