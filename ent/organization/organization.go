// Code generated by ent, DO NOT EDIT.

package organization

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the organization type in the database.
	Label = "organization"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// EdgeObject holds the string denoting the object edge name in mutations.
	EdgeObject = "object"
	// Table holds the table name of the organization in the database.
	Table = "organizations"
	// ObjectTable is the table that holds the object relation/edge.
	ObjectTable = "organizations"
	// ObjectInverseTable is the table name for the Object entity.
	// It exists in this package in order to avoid circular dependency with the "object" package.
	ObjectInverseTable = "objects"
	// ObjectColumn is the table column denoting the object relation/edge.
	ObjectColumn = "object_organization"
)

// Columns holds all SQL columns for organization fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldType,
	FieldAddress,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "organizations"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"object_organization",
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
)

// OrderOption defines the ordering options for the Organization queries.
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

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByObjectField orders the results by object field.
func ByObjectField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newObjectStep(), sql.OrderByField(field, opts...))
	}
}
func newObjectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ObjectInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, ObjectTable, ObjectColumn),
	)
}
