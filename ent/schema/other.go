package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Other holds the schema definition for the Other entity.
type Other struct {
	ent.Schema
}

// Fields of the Other.
func (Other) Fields() []ent.Field {
	return nil
}

// Edges of the Other.
func (Other) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("object", Object.Type).
			Ref("other").
			Unique().
			Required(),
	}
}
