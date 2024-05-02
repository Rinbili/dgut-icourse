package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Person holds the schema definition for the Person entity.
type Person struct {
	ent.Schema
}

// Fields of the Person.
func (Person) Fields() []ent.Field {
	return nil
}

// Edges of the Person.
func (Person) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("object", Object.Type).
			Ref("person").
			Unique().
			Required(),
	}
}
