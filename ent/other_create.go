// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"dgut-icourse/ent/object"
	"dgut-icourse/ent/other"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// OtherCreate is the builder for creating a Other entity.
type OtherCreate struct {
	config
	mutation *OtherMutation
	hooks    []Hook
}

// SetObjectID sets the "object" edge to the Object entity by ID.
func (oc *OtherCreate) SetObjectID(id uuid.UUID) *OtherCreate {
	oc.mutation.SetObjectID(id)
	return oc
}

// SetObject sets the "object" edge to the Object entity.
func (oc *OtherCreate) SetObject(o *Object) *OtherCreate {
	return oc.SetObjectID(o.ID)
}

// Mutation returns the OtherMutation object of the builder.
func (oc *OtherCreate) Mutation() *OtherMutation {
	return oc.mutation
}

// Save creates the Other in the database.
func (oc *OtherCreate) Save(ctx context.Context) (*Other, error) {
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OtherCreate) SaveX(ctx context.Context) *Other {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OtherCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OtherCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OtherCreate) check() error {
	if _, ok := oc.mutation.ObjectID(); !ok {
		return &ValidationError{Name: "object", err: errors.New(`ent: missing required edge "Other.object"`)}
	}
	return nil
}

func (oc *OtherCreate) sqlSave(ctx context.Context) (*Other, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OtherCreate) createSpec() (*Other, *sqlgraph.CreateSpec) {
	var (
		_node = &Other{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(other.Table, sqlgraph.NewFieldSpec(other.FieldID, field.TypeInt))
	)
	if nodes := oc.mutation.ObjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   other.ObjectTable,
			Columns: []string{other.ObjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(object.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.object_other = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OtherCreateBulk is the builder for creating many Other entities in bulk.
type OtherCreateBulk struct {
	config
	err      error
	builders []*OtherCreate
}

// Save creates the Other entities in the database.
func (ocb *OtherCreateBulk) Save(ctx context.Context) ([]*Other, error) {
	if ocb.err != nil {
		return nil, ocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Other, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OtherMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OtherCreateBulk) SaveX(ctx context.Context) []*Other {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OtherCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OtherCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
