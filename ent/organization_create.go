// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"dgut-icourse/ent/object"
	"dgut-icourse/ent/organization"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// OrganizationCreate is the builder for creating a Organization entity.
type OrganizationCreate struct {
	config
	mutation *OrganizationMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (oc *OrganizationCreate) SetCreatedAt(i int64) *OrganizationCreate {
	oc.mutation.SetCreatedAt(i)
	return oc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableCreatedAt(i *int64) *OrganizationCreate {
	if i != nil {
		oc.SetCreatedAt(*i)
	}
	return oc
}

// SetUpdatedAt sets the "updated_at" field.
func (oc *OrganizationCreate) SetUpdatedAt(i int64) *OrganizationCreate {
	oc.mutation.SetUpdatedAt(i)
	return oc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableUpdatedAt(i *int64) *OrganizationCreate {
	if i != nil {
		oc.SetUpdatedAt(*i)
	}
	return oc
}

// SetType sets the "type" field.
func (oc *OrganizationCreate) SetType(i int8) *OrganizationCreate {
	oc.mutation.SetType(i)
	return oc
}

// SetAddress sets the "address" field.
func (oc *OrganizationCreate) SetAddress(s string) *OrganizationCreate {
	oc.mutation.SetAddress(s)
	return oc
}

// SetObjectID sets the "object" edge to the Object entity by ID.
func (oc *OrganizationCreate) SetObjectID(id uuid.UUID) *OrganizationCreate {
	oc.mutation.SetObjectID(id)
	return oc
}

// SetObject sets the "object" edge to the Object entity.
func (oc *OrganizationCreate) SetObject(o *Object) *OrganizationCreate {
	return oc.SetObjectID(o.ID)
}

// Mutation returns the OrganizationMutation object of the builder.
func (oc *OrganizationCreate) Mutation() *OrganizationMutation {
	return oc.mutation
}

// Save creates the Organization in the database.
func (oc *OrganizationCreate) Save(ctx context.Context) (*Organization, error) {
	oc.defaults()
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrganizationCreate) SaveX(ctx context.Context) *Organization {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrganizationCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrganizationCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OrganizationCreate) defaults() {
	if _, ok := oc.mutation.CreatedAt(); !ok {
		v := organization.DefaultCreatedAt()
		oc.mutation.SetCreatedAt(v)
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		v := organization.DefaultUpdatedAt()
		oc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrganizationCreate) check() error {
	if _, ok := oc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Organization.created_at"`)}
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Organization.updated_at"`)}
	}
	if _, ok := oc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Organization.type"`)}
	}
	if _, ok := oc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Organization.address"`)}
	}
	if _, ok := oc.mutation.ObjectID(); !ok {
		return &ValidationError{Name: "object", err: errors.New(`ent: missing required edge "Organization.object"`)}
	}
	return nil
}

func (oc *OrganizationCreate) sqlSave(ctx context.Context) (*Organization, error) {
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

func (oc *OrganizationCreate) createSpec() (*Organization, *sqlgraph.CreateSpec) {
	var (
		_node = &Organization{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(organization.Table, sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt))
	)
	if value, ok := oc.mutation.CreatedAt(); ok {
		_spec.SetField(organization.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := oc.mutation.UpdatedAt(); ok {
		_spec.SetField(organization.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := oc.mutation.GetType(); ok {
		_spec.SetField(organization.FieldType, field.TypeInt8, value)
		_node.Type = value
	}
	if value, ok := oc.mutation.Address(); ok {
		_spec.SetField(organization.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if nodes := oc.mutation.ObjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   organization.ObjectTable,
			Columns: []string{organization.ObjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(object.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.object_organization = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrganizationCreateBulk is the builder for creating many Organization entities in bulk.
type OrganizationCreateBulk struct {
	config
	err      error
	builders []*OrganizationCreate
}

// Save creates the Organization entities in the database.
func (ocb *OrganizationCreateBulk) Save(ctx context.Context) ([]*Organization, error) {
	if ocb.err != nil {
		return nil, ocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Organization, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrganizationMutation)
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
func (ocb *OrganizationCreateBulk) SaveX(ctx context.Context) []*Organization {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrganizationCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrganizationCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
