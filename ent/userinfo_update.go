// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"dgut-icourse/ent/predicate"
	"dgut-icourse/ent/userinfo"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserInfoUpdate is the builder for updating UserInfo entities.
type UserInfoUpdate struct {
	config
	hooks    []Hook
	mutation *UserInfoMutation
}

// Where appends a list predicates to the UserInfoUpdate builder.
func (uiu *UserInfoUpdate) Where(ps ...predicate.UserInfo) *UserInfoUpdate {
	uiu.mutation.Where(ps...)
	return uiu
}

// SetUpdatedAt sets the "updated_at" field.
func (uiu *UserInfoUpdate) SetUpdatedAt(i int64) *UserInfoUpdate {
	uiu.mutation.ResetUpdatedAt()
	uiu.mutation.SetUpdatedAt(i)
	return uiu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (uiu *UserInfoUpdate) AddUpdatedAt(i int64) *UserInfoUpdate {
	uiu.mutation.AddUpdatedAt(i)
	return uiu
}

// Mutation returns the UserInfoMutation object of the builder.
func (uiu *UserInfoUpdate) Mutation() *UserInfoMutation {
	return uiu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uiu *UserInfoUpdate) Save(ctx context.Context) (int, error) {
	uiu.defaults()
	return withHooks(ctx, uiu.sqlSave, uiu.mutation, uiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uiu *UserInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := uiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uiu *UserInfoUpdate) Exec(ctx context.Context) error {
	_, err := uiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uiu *UserInfoUpdate) ExecX(ctx context.Context) {
	if err := uiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uiu *UserInfoUpdate) defaults() {
	if _, ok := uiu.mutation.UpdatedAt(); !ok {
		v := userinfo.UpdateDefaultUpdatedAt()
		uiu.mutation.SetUpdatedAt(v)
	}
}

func (uiu *UserInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(userinfo.Table, userinfo.Columns, sqlgraph.NewFieldSpec(userinfo.FieldID, field.TypeInt))
	if ps := uiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uiu.mutation.UpdatedAt(); ok {
		_spec.SetField(userinfo.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := uiu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(userinfo.FieldUpdatedAt, field.TypeInt64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uiu.mutation.done = true
	return n, nil
}

// UserInfoUpdateOne is the builder for updating a single UserInfo entity.
type UserInfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserInfoMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (uiuo *UserInfoUpdateOne) SetUpdatedAt(i int64) *UserInfoUpdateOne {
	uiuo.mutation.ResetUpdatedAt()
	uiuo.mutation.SetUpdatedAt(i)
	return uiuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (uiuo *UserInfoUpdateOne) AddUpdatedAt(i int64) *UserInfoUpdateOne {
	uiuo.mutation.AddUpdatedAt(i)
	return uiuo
}

// Mutation returns the UserInfoMutation object of the builder.
func (uiuo *UserInfoUpdateOne) Mutation() *UserInfoMutation {
	return uiuo.mutation
}

// Where appends a list predicates to the UserInfoUpdate builder.
func (uiuo *UserInfoUpdateOne) Where(ps ...predicate.UserInfo) *UserInfoUpdateOne {
	uiuo.mutation.Where(ps...)
	return uiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uiuo *UserInfoUpdateOne) Select(field string, fields ...string) *UserInfoUpdateOne {
	uiuo.fields = append([]string{field}, fields...)
	return uiuo
}

// Save executes the query and returns the updated UserInfo entity.
func (uiuo *UserInfoUpdateOne) Save(ctx context.Context) (*UserInfo, error) {
	uiuo.defaults()
	return withHooks(ctx, uiuo.sqlSave, uiuo.mutation, uiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uiuo *UserInfoUpdateOne) SaveX(ctx context.Context) *UserInfo {
	node, err := uiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uiuo *UserInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := uiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uiuo *UserInfoUpdateOne) ExecX(ctx context.Context) {
	if err := uiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uiuo *UserInfoUpdateOne) defaults() {
	if _, ok := uiuo.mutation.UpdatedAt(); !ok {
		v := userinfo.UpdateDefaultUpdatedAt()
		uiuo.mutation.SetUpdatedAt(v)
	}
}

func (uiuo *UserInfoUpdateOne) sqlSave(ctx context.Context) (_node *UserInfo, err error) {
	_spec := sqlgraph.NewUpdateSpec(userinfo.Table, userinfo.Columns, sqlgraph.NewFieldSpec(userinfo.FieldID, field.TypeInt))
	id, ok := uiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userinfo.FieldID)
		for _, f := range fields {
			if !userinfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uiuo.mutation.UpdatedAt(); ok {
		_spec.SetField(userinfo.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := uiuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(userinfo.FieldUpdatedAt, field.TypeInt64, value)
	}
	_node = &UserInfo{config: uiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uiuo.mutation.done = true
	return _node, nil
}
