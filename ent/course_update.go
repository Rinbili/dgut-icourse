// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"dgut-icourse/ent/course"
	"dgut-icourse/ent/object"
	"dgut-icourse/ent/predicate"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CourseUpdate is the builder for updating Course entities.
type CourseUpdate struct {
	config
	hooks    []Hook
	mutation *CourseMutation
}

// Where appends a list predicates to the CourseUpdate builder.
func (cu *CourseUpdate) Where(ps ...predicate.Course) *CourseUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CourseUpdate) SetUpdatedAt(i int64) *CourseUpdate {
	cu.mutation.ResetUpdatedAt()
	cu.mutation.SetUpdatedAt(i)
	return cu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (cu *CourseUpdate) AddUpdatedAt(i int64) *CourseUpdate {
	cu.mutation.AddUpdatedAt(i)
	return cu
}

// SetCourseID sets the "courseID" field.
func (cu *CourseUpdate) SetCourseID(s string) *CourseUpdate {
	cu.mutation.SetCourseID(s)
	return cu
}

// SetNillableCourseID sets the "courseID" field if the given value is not nil.
func (cu *CourseUpdate) SetNillableCourseID(s *string) *CourseUpdate {
	if s != nil {
		cu.SetCourseID(*s)
	}
	return cu
}

// SetObjectID sets the "object" edge to the Object entity by ID.
func (cu *CourseUpdate) SetObjectID(id uuid.UUID) *CourseUpdate {
	cu.mutation.SetObjectID(id)
	return cu
}

// SetObject sets the "object" edge to the Object entity.
func (cu *CourseUpdate) SetObject(o *Object) *CourseUpdate {
	return cu.SetObjectID(o.ID)
}

// SetTeacherID sets the "teacher" edge to the Object entity by ID.
func (cu *CourseUpdate) SetTeacherID(id uuid.UUID) *CourseUpdate {
	cu.mutation.SetTeacherID(id)
	return cu
}

// SetNillableTeacherID sets the "teacher" edge to the Object entity by ID if the given value is not nil.
func (cu *CourseUpdate) SetNillableTeacherID(id *uuid.UUID) *CourseUpdate {
	if id != nil {
		cu = cu.SetTeacherID(*id)
	}
	return cu
}

// SetTeacher sets the "teacher" edge to the Object entity.
func (cu *CourseUpdate) SetTeacher(o *Object) *CourseUpdate {
	return cu.SetTeacherID(o.ID)
}

// SetOrganizationID sets the "organization" edge to the Object entity by ID.
func (cu *CourseUpdate) SetOrganizationID(id uuid.UUID) *CourseUpdate {
	cu.mutation.SetOrganizationID(id)
	return cu
}

// SetNillableOrganizationID sets the "organization" edge to the Object entity by ID if the given value is not nil.
func (cu *CourseUpdate) SetNillableOrganizationID(id *uuid.UUID) *CourseUpdate {
	if id != nil {
		cu = cu.SetOrganizationID(*id)
	}
	return cu
}

// SetOrganization sets the "organization" edge to the Object entity.
func (cu *CourseUpdate) SetOrganization(o *Object) *CourseUpdate {
	return cu.SetOrganizationID(o.ID)
}

// Mutation returns the CourseMutation object of the builder.
func (cu *CourseUpdate) Mutation() *CourseMutation {
	return cu.mutation
}

// ClearObject clears the "object" edge to the Object entity.
func (cu *CourseUpdate) ClearObject() *CourseUpdate {
	cu.mutation.ClearObject()
	return cu
}

// ClearTeacher clears the "teacher" edge to the Object entity.
func (cu *CourseUpdate) ClearTeacher() *CourseUpdate {
	cu.mutation.ClearTeacher()
	return cu
}

// ClearOrganization clears the "organization" edge to the Object entity.
func (cu *CourseUpdate) ClearOrganization() *CourseUpdate {
	cu.mutation.ClearOrganization()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CourseUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CourseUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CourseUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CourseUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CourseUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := course.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CourseUpdate) check() error {
	if _, ok := cu.mutation.ObjectID(); cu.mutation.ObjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Course.object"`)
	}
	return nil
}

func (cu *CourseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(course.Table, course.Columns, sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(course.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(course.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.CourseID(); ok {
		_spec.SetField(course.FieldCourseID, field.TypeString, value)
	}
	if cu.mutation.ObjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   course.ObjectTable,
			Columns: []string{course.ObjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(object.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ObjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   course.ObjectTable,
			Columns: []string{course.ObjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(object.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.TeacherCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   course.TeacherTable,
			Columns: []string{course.TeacherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(object.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.TeacherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   course.TeacherTable,
			Columns: []string{course.TeacherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(object.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   course.OrganizationTable,
			Columns: []string{course.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(object.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   course.OrganizationTable,
			Columns: []string{course.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(object.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{course.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CourseUpdateOne is the builder for updating a single Course entity.
type CourseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CourseMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CourseUpdateOne) SetUpdatedAt(i int64) *CourseUpdateOne {
	cuo.mutation.ResetUpdatedAt()
	cuo.mutation.SetUpdatedAt(i)
	return cuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (cuo *CourseUpdateOne) AddUpdatedAt(i int64) *CourseUpdateOne {
	cuo.mutation.AddUpdatedAt(i)
	return cuo
}

// SetCourseID sets the "courseID" field.
func (cuo *CourseUpdateOne) SetCourseID(s string) *CourseUpdateOne {
	cuo.mutation.SetCourseID(s)
	return cuo
}

// SetNillableCourseID sets the "courseID" field if the given value is not nil.
func (cuo *CourseUpdateOne) SetNillableCourseID(s *string) *CourseUpdateOne {
	if s != nil {
		cuo.SetCourseID(*s)
	}
	return cuo
}

// SetObjectID sets the "object" edge to the Object entity by ID.
func (cuo *CourseUpdateOne) SetObjectID(id uuid.UUID) *CourseUpdateOne {
	cuo.mutation.SetObjectID(id)
	return cuo
}

// SetObject sets the "object" edge to the Object entity.
func (cuo *CourseUpdateOne) SetObject(o *Object) *CourseUpdateOne {
	return cuo.SetObjectID(o.ID)
}

// SetTeacherID sets the "teacher" edge to the Object entity by ID.
func (cuo *CourseUpdateOne) SetTeacherID(id uuid.UUID) *CourseUpdateOne {
	cuo.mutation.SetTeacherID(id)
	return cuo
}

// SetNillableTeacherID sets the "teacher" edge to the Object entity by ID if the given value is not nil.
func (cuo *CourseUpdateOne) SetNillableTeacherID(id *uuid.UUID) *CourseUpdateOne {
	if id != nil {
		cuo = cuo.SetTeacherID(*id)
	}
	return cuo
}

// SetTeacher sets the "teacher" edge to the Object entity.
func (cuo *CourseUpdateOne) SetTeacher(o *Object) *CourseUpdateOne {
	return cuo.SetTeacherID(o.ID)
}

// SetOrganizationID sets the "organization" edge to the Object entity by ID.
func (cuo *CourseUpdateOne) SetOrganizationID(id uuid.UUID) *CourseUpdateOne {
	cuo.mutation.SetOrganizationID(id)
	return cuo
}

// SetNillableOrganizationID sets the "organization" edge to the Object entity by ID if the given value is not nil.
func (cuo *CourseUpdateOne) SetNillableOrganizationID(id *uuid.UUID) *CourseUpdateOne {
	if id != nil {
		cuo = cuo.SetOrganizationID(*id)
	}
	return cuo
}

// SetOrganization sets the "organization" edge to the Object entity.
func (cuo *CourseUpdateOne) SetOrganization(o *Object) *CourseUpdateOne {
	return cuo.SetOrganizationID(o.ID)
}

// Mutation returns the CourseMutation object of the builder.
func (cuo *CourseUpdateOne) Mutation() *CourseMutation {
	return cuo.mutation
}

// ClearObject clears the "object" edge to the Object entity.
func (cuo *CourseUpdateOne) ClearObject() *CourseUpdateOne {
	cuo.mutation.ClearObject()
	return cuo
}

// ClearTeacher clears the "teacher" edge to the Object entity.
func (cuo *CourseUpdateOne) ClearTeacher() *CourseUpdateOne {
	cuo.mutation.ClearTeacher()
	return cuo
}

// ClearOrganization clears the "organization" edge to the Object entity.
func (cuo *CourseUpdateOne) ClearOrganization() *CourseUpdateOne {
	cuo.mutation.ClearOrganization()
	return cuo
}

// Where appends a list predicates to the CourseUpdate builder.
func (cuo *CourseUpdateOne) Where(ps ...predicate.Course) *CourseUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CourseUpdateOne) Select(field string, fields ...string) *CourseUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Course entity.
func (cuo *CourseUpdateOne) Save(ctx context.Context) (*Course, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CourseUpdateOne) SaveX(ctx context.Context) *Course {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CourseUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CourseUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CourseUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := course.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CourseUpdateOne) check() error {
	if _, ok := cuo.mutation.ObjectID(); cuo.mutation.ObjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Course.object"`)
	}
	return nil
}

func (cuo *CourseUpdateOne) sqlSave(ctx context.Context) (_node *Course, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(course.Table, course.Columns, sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Course.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, course.FieldID)
		for _, f := range fields {
			if !course.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != course.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(course.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(course.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.CourseID(); ok {
		_spec.SetField(course.FieldCourseID, field.TypeString, value)
	}
	if cuo.mutation.ObjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   course.ObjectTable,
			Columns: []string{course.ObjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(object.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ObjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   course.ObjectTable,
			Columns: []string{course.ObjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(object.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.TeacherCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   course.TeacherTable,
			Columns: []string{course.TeacherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(object.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.TeacherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   course.TeacherTable,
			Columns: []string{course.TeacherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(object.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   course.OrganizationTable,
			Columns: []string{course.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(object.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   course.OrganizationTable,
			Columns: []string{course.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(object.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Course{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{course.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
