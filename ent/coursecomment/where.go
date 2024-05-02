// Code generated by ent, DO NOT EDIT.

package coursecomment

import (
	"dgut-icourse/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldLTE(FieldID, id))
}

// Difficulty applies equality check predicate on the "difficulty" field. It's identical to DifficultyEQ.
func Difficulty(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldEQ(FieldDifficulty, v))
}

// Quality applies equality check predicate on the "quality" field. It's identical to QualityEQ.
func Quality(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldEQ(FieldQuality, v))
}

// Workload applies equality check predicate on the "workload" field. It's identical to WorkloadEQ.
func Workload(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldEQ(FieldWorkload, v))
}

// Mark applies equality check predicate on the "mark" field. It's identical to MarkEQ.
func Mark(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldEQ(FieldMark, v))
}

// DifficultyEQ applies the EQ predicate on the "difficulty" field.
func DifficultyEQ(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldEQ(FieldDifficulty, v))
}

// DifficultyNEQ applies the NEQ predicate on the "difficulty" field.
func DifficultyNEQ(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldNEQ(FieldDifficulty, v))
}

// DifficultyIn applies the In predicate on the "difficulty" field.
func DifficultyIn(vs ...int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldIn(FieldDifficulty, vs...))
}

// DifficultyNotIn applies the NotIn predicate on the "difficulty" field.
func DifficultyNotIn(vs ...int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldNotIn(FieldDifficulty, vs...))
}

// DifficultyGT applies the GT predicate on the "difficulty" field.
func DifficultyGT(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldGT(FieldDifficulty, v))
}

// DifficultyGTE applies the GTE predicate on the "difficulty" field.
func DifficultyGTE(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldGTE(FieldDifficulty, v))
}

// DifficultyLT applies the LT predicate on the "difficulty" field.
func DifficultyLT(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldLT(FieldDifficulty, v))
}

// DifficultyLTE applies the LTE predicate on the "difficulty" field.
func DifficultyLTE(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldLTE(FieldDifficulty, v))
}

// QualityEQ applies the EQ predicate on the "quality" field.
func QualityEQ(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldEQ(FieldQuality, v))
}

// QualityNEQ applies the NEQ predicate on the "quality" field.
func QualityNEQ(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldNEQ(FieldQuality, v))
}

// QualityIn applies the In predicate on the "quality" field.
func QualityIn(vs ...int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldIn(FieldQuality, vs...))
}

// QualityNotIn applies the NotIn predicate on the "quality" field.
func QualityNotIn(vs ...int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldNotIn(FieldQuality, vs...))
}

// QualityGT applies the GT predicate on the "quality" field.
func QualityGT(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldGT(FieldQuality, v))
}

// QualityGTE applies the GTE predicate on the "quality" field.
func QualityGTE(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldGTE(FieldQuality, v))
}

// QualityLT applies the LT predicate on the "quality" field.
func QualityLT(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldLT(FieldQuality, v))
}

// QualityLTE applies the LTE predicate on the "quality" field.
func QualityLTE(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldLTE(FieldQuality, v))
}

// WorkloadEQ applies the EQ predicate on the "workload" field.
func WorkloadEQ(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldEQ(FieldWorkload, v))
}

// WorkloadNEQ applies the NEQ predicate on the "workload" field.
func WorkloadNEQ(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldNEQ(FieldWorkload, v))
}

// WorkloadIn applies the In predicate on the "workload" field.
func WorkloadIn(vs ...int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldIn(FieldWorkload, vs...))
}

// WorkloadNotIn applies the NotIn predicate on the "workload" field.
func WorkloadNotIn(vs ...int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldNotIn(FieldWorkload, vs...))
}

// WorkloadGT applies the GT predicate on the "workload" field.
func WorkloadGT(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldGT(FieldWorkload, v))
}

// WorkloadGTE applies the GTE predicate on the "workload" field.
func WorkloadGTE(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldGTE(FieldWorkload, v))
}

// WorkloadLT applies the LT predicate on the "workload" field.
func WorkloadLT(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldLT(FieldWorkload, v))
}

// WorkloadLTE applies the LTE predicate on the "workload" field.
func WorkloadLTE(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldLTE(FieldWorkload, v))
}

// MarkEQ applies the EQ predicate on the "mark" field.
func MarkEQ(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldEQ(FieldMark, v))
}

// MarkNEQ applies the NEQ predicate on the "mark" field.
func MarkNEQ(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldNEQ(FieldMark, v))
}

// MarkIn applies the In predicate on the "mark" field.
func MarkIn(vs ...int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldIn(FieldMark, vs...))
}

// MarkNotIn applies the NotIn predicate on the "mark" field.
func MarkNotIn(vs ...int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldNotIn(FieldMark, vs...))
}

// MarkGT applies the GT predicate on the "mark" field.
func MarkGT(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldGT(FieldMark, v))
}

// MarkGTE applies the GTE predicate on the "mark" field.
func MarkGTE(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldGTE(FieldMark, v))
}

// MarkLT applies the LT predicate on the "mark" field.
func MarkLT(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldLT(FieldMark, v))
}

// MarkLTE applies the LTE predicate on the "mark" field.
func MarkLTE(v int8) predicate.CourseComment {
	return predicate.CourseComment(sql.FieldLTE(FieldMark, v))
}

// HasComment applies the HasEdge predicate on the "comment" edge.
func HasComment() predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, CommentTable, CommentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentWith applies the HasEdge predicate on the "comment" edge with a given conditions (other predicates).
func HasCommentWith(preds ...predicate.Comment) predicate.CourseComment {
	return predicate.CourseComment(func(s *sql.Selector) {
		step := newCommentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CourseComment) predicate.CourseComment {
	return predicate.CourseComment(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CourseComment) predicate.CourseComment {
	return predicate.CourseComment(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CourseComment) predicate.CourseComment {
	return predicate.CourseComment(sql.NotPredicates(p))
}