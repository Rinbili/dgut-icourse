// Code generated by ent, DO NOT EDIT.

package comment

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the comment type in the database.
	Label = "comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldScore holds the string denoting the score field in the database.
	FieldScore = "score"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// EdgeObject holds the string denoting the object edge name in mutations.
	EdgeObject = "object"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeCourseComment holds the string denoting the course_comment edge name in mutations.
	EdgeCourseComment = "course_comment"
	// Table holds the table name of the comment in the database.
	Table = "comments"
	// AuthorTable is the table that holds the author relation/edge.
	AuthorTable = "comments"
	// AuthorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AuthorInverseTable = "users"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "comment_author"
	// ObjectTable is the table that holds the object relation/edge.
	ObjectTable = "comments"
	// ObjectInverseTable is the table name for the Object entity.
	// It exists in this package in order to avoid circular dependency with the "object" package.
	ObjectInverseTable = "objects"
	// ObjectColumn is the table column denoting the object relation/edge.
	ObjectColumn = "object_comments"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "comments"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "comment_children"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "comments"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "comment_children"
	// CourseCommentTable is the table that holds the course_comment relation/edge.
	CourseCommentTable = "course_comments"
	// CourseCommentInverseTable is the table name for the CourseComment entity.
	// It exists in this package in order to avoid circular dependency with the "coursecomment" package.
	CourseCommentInverseTable = "course_comments"
	// CourseCommentColumn is the table column denoting the course_comment relation/edge.
	CourseCommentColumn = "comment_course_comment"
)

// Columns holds all SQL columns for comment fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldContent,
	FieldScore,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "comments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"comment_author",
	"comment_children",
	"object_comments",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() int64
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() int64
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() int64
	// ContentValidator is a validator for the "content" field. It is called by the builders before save.
	ContentValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Comment queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByScore orders the results by the score field.
func ByScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScore, opts...).ToFunc()
}

// ByAuthorField orders the results by author field.
func ByAuthorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthorStep(), sql.OrderByField(field, opts...))
	}
}

// ByObjectField orders the results by object field.
func ByObjectField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newObjectStep(), sql.OrderByField(field, opts...))
	}
}

// ByParentField orders the results by parent field.
func ByParentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentStep(), sql.OrderByField(field, opts...))
	}
}

// ByChildrenCount orders the results by children count.
func ByChildrenCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChildrenStep(), opts...)
	}
}

// ByChildren orders the results by children terms.
func ByChildren(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChildrenStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCourseCommentField orders the results by course_comment field.
func ByCourseCommentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCourseCommentStep(), sql.OrderByField(field, opts...))
	}
}
func newAuthorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, AuthorTable, AuthorColumn),
	)
}
func newObjectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ObjectInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ObjectTable, ObjectColumn),
	)
}
func newParentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
	)
}
func newChildrenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
	)
}
func newCourseCommentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CourseCommentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, CourseCommentTable, CourseCommentColumn),
	)
}