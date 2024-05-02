// Code generated by ent, DO NOT EDIT.

package ent

import (
	"dgut-icourse/ent/comment"
	"dgut-icourse/ent/coursecomment"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// CourseComment is the model entity for the CourseComment schema.
type CourseComment struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 难度
	Difficulty int8 `json:"difficulty,omitempty"`
	// 收获
	Quality int8 `json:"quality,omitempty"`
	// 作业量
	Workload int8 `json:"workload,omitempty"`
	// 评分好坏
	Mark int8 `json:"mark,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CourseCommentQuery when eager-loading is set.
	Edges                  CourseCommentEdges `json:"edges"`
	comment_course_comment *uuid.UUID
	selectValues           sql.SelectValues
}

// CourseCommentEdges holds the relations/edges for other nodes in the graph.
type CourseCommentEdges struct {
	// Comment holds the value of the comment edge.
	Comment *Comment `json:"comment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CommentOrErr returns the Comment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CourseCommentEdges) CommentOrErr() (*Comment, error) {
	if e.Comment != nil {
		return e.Comment, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: comment.Label}
	}
	return nil, &NotLoadedError{edge: "comment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CourseComment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case coursecomment.FieldID, coursecomment.FieldDifficulty, coursecomment.FieldQuality, coursecomment.FieldWorkload, coursecomment.FieldMark:
			values[i] = new(sql.NullInt64)
		case coursecomment.ForeignKeys[0]: // comment_course_comment
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CourseComment fields.
func (cc *CourseComment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case coursecomment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cc.ID = int(value.Int64)
		case coursecomment.FieldDifficulty:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field difficulty", values[i])
			} else if value.Valid {
				cc.Difficulty = int8(value.Int64)
			}
		case coursecomment.FieldQuality:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quality", values[i])
			} else if value.Valid {
				cc.Quality = int8(value.Int64)
			}
		case coursecomment.FieldWorkload:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field workload", values[i])
			} else if value.Valid {
				cc.Workload = int8(value.Int64)
			}
		case coursecomment.FieldMark:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mark", values[i])
			} else if value.Valid {
				cc.Mark = int8(value.Int64)
			}
		case coursecomment.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field comment_course_comment", values[i])
			} else if value.Valid {
				cc.comment_course_comment = new(uuid.UUID)
				*cc.comment_course_comment = *value.S.(*uuid.UUID)
			}
		default:
			cc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CourseComment.
// This includes values selected through modifiers, order, etc.
func (cc *CourseComment) Value(name string) (ent.Value, error) {
	return cc.selectValues.Get(name)
}

// QueryComment queries the "comment" edge of the CourseComment entity.
func (cc *CourseComment) QueryComment() *CommentQuery {
	return NewCourseCommentClient(cc.config).QueryComment(cc)
}

// Update returns a builder for updating this CourseComment.
// Note that you need to call CourseComment.Unwrap() before calling this method if this CourseComment
// was returned from a transaction, and the transaction was committed or rolled back.
func (cc *CourseComment) Update() *CourseCommentUpdateOne {
	return NewCourseCommentClient(cc.config).UpdateOne(cc)
}

// Unwrap unwraps the CourseComment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cc *CourseComment) Unwrap() *CourseComment {
	_tx, ok := cc.config.driver.(*txDriver)
	if !ok {
		panic("ent: CourseComment is not a transactional entity")
	}
	cc.config.driver = _tx.drv
	return cc
}

// String implements the fmt.Stringer.
func (cc *CourseComment) String() string {
	var builder strings.Builder
	builder.WriteString("CourseComment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cc.ID))
	builder.WriteString("difficulty=")
	builder.WriteString(fmt.Sprintf("%v", cc.Difficulty))
	builder.WriteString(", ")
	builder.WriteString("quality=")
	builder.WriteString(fmt.Sprintf("%v", cc.Quality))
	builder.WriteString(", ")
	builder.WriteString("workload=")
	builder.WriteString(fmt.Sprintf("%v", cc.Workload))
	builder.WriteString(", ")
	builder.WriteString("mark=")
	builder.WriteString(fmt.Sprintf("%v", cc.Mark))
	builder.WriteByte(')')
	return builder.String()
}

// CourseComments is a parsable slice of CourseComment.
type CourseComments []*CourseComment