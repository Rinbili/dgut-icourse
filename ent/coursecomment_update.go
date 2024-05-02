// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"dgut-icourse/ent/comment"
	"dgut-icourse/ent/coursecomment"
	"dgut-icourse/ent/predicate"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CourseCommentUpdate is the builder for updating CourseComment entities.
type CourseCommentUpdate struct {
	config
	hooks    []Hook
	mutation *CourseCommentMutation
}

// Where appends a list predicates to the CourseCommentUpdate builder.
func (ccu *CourseCommentUpdate) Where(ps ...predicate.CourseComment) *CourseCommentUpdate {
	ccu.mutation.Where(ps...)
	return ccu
}

// SetDifficulty sets the "difficulty" field.
func (ccu *CourseCommentUpdate) SetDifficulty(i int8) *CourseCommentUpdate {
	ccu.mutation.ResetDifficulty()
	ccu.mutation.SetDifficulty(i)
	return ccu
}

// SetNillableDifficulty sets the "difficulty" field if the given value is not nil.
func (ccu *CourseCommentUpdate) SetNillableDifficulty(i *int8) *CourseCommentUpdate {
	if i != nil {
		ccu.SetDifficulty(*i)
	}
	return ccu
}

// AddDifficulty adds i to the "difficulty" field.
func (ccu *CourseCommentUpdate) AddDifficulty(i int8) *CourseCommentUpdate {
	ccu.mutation.AddDifficulty(i)
	return ccu
}

// SetQuality sets the "quality" field.
func (ccu *CourseCommentUpdate) SetQuality(i int8) *CourseCommentUpdate {
	ccu.mutation.ResetQuality()
	ccu.mutation.SetQuality(i)
	return ccu
}

// SetNillableQuality sets the "quality" field if the given value is not nil.
func (ccu *CourseCommentUpdate) SetNillableQuality(i *int8) *CourseCommentUpdate {
	if i != nil {
		ccu.SetQuality(*i)
	}
	return ccu
}

// AddQuality adds i to the "quality" field.
func (ccu *CourseCommentUpdate) AddQuality(i int8) *CourseCommentUpdate {
	ccu.mutation.AddQuality(i)
	return ccu
}

// SetWorkload sets the "workload" field.
func (ccu *CourseCommentUpdate) SetWorkload(i int8) *CourseCommentUpdate {
	ccu.mutation.ResetWorkload()
	ccu.mutation.SetWorkload(i)
	return ccu
}

// SetNillableWorkload sets the "workload" field if the given value is not nil.
func (ccu *CourseCommentUpdate) SetNillableWorkload(i *int8) *CourseCommentUpdate {
	if i != nil {
		ccu.SetWorkload(*i)
	}
	return ccu
}

// AddWorkload adds i to the "workload" field.
func (ccu *CourseCommentUpdate) AddWorkload(i int8) *CourseCommentUpdate {
	ccu.mutation.AddWorkload(i)
	return ccu
}

// SetMark sets the "mark" field.
func (ccu *CourseCommentUpdate) SetMark(i int8) *CourseCommentUpdate {
	ccu.mutation.ResetMark()
	ccu.mutation.SetMark(i)
	return ccu
}

// SetNillableMark sets the "mark" field if the given value is not nil.
func (ccu *CourseCommentUpdate) SetNillableMark(i *int8) *CourseCommentUpdate {
	if i != nil {
		ccu.SetMark(*i)
	}
	return ccu
}

// AddMark adds i to the "mark" field.
func (ccu *CourseCommentUpdate) AddMark(i int8) *CourseCommentUpdate {
	ccu.mutation.AddMark(i)
	return ccu
}

// SetCommentID sets the "comment" edge to the Comment entity by ID.
func (ccu *CourseCommentUpdate) SetCommentID(id uuid.UUID) *CourseCommentUpdate {
	ccu.mutation.SetCommentID(id)
	return ccu
}

// SetComment sets the "comment" edge to the Comment entity.
func (ccu *CourseCommentUpdate) SetComment(c *Comment) *CourseCommentUpdate {
	return ccu.SetCommentID(c.ID)
}

// Mutation returns the CourseCommentMutation object of the builder.
func (ccu *CourseCommentUpdate) Mutation() *CourseCommentMutation {
	return ccu.mutation
}

// ClearComment clears the "comment" edge to the Comment entity.
func (ccu *CourseCommentUpdate) ClearComment() *CourseCommentUpdate {
	ccu.mutation.ClearComment()
	return ccu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ccu *CourseCommentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ccu.sqlSave, ccu.mutation, ccu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ccu *CourseCommentUpdate) SaveX(ctx context.Context) int {
	affected, err := ccu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ccu *CourseCommentUpdate) Exec(ctx context.Context) error {
	_, err := ccu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccu *CourseCommentUpdate) ExecX(ctx context.Context) {
	if err := ccu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ccu *CourseCommentUpdate) check() error {
	if _, ok := ccu.mutation.CommentID(); ccu.mutation.CommentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CourseComment.comment"`)
	}
	return nil
}

func (ccu *CourseCommentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ccu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(coursecomment.Table, coursecomment.Columns, sqlgraph.NewFieldSpec(coursecomment.FieldID, field.TypeInt))
	if ps := ccu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccu.mutation.Difficulty(); ok {
		_spec.SetField(coursecomment.FieldDifficulty, field.TypeInt8, value)
	}
	if value, ok := ccu.mutation.AddedDifficulty(); ok {
		_spec.AddField(coursecomment.FieldDifficulty, field.TypeInt8, value)
	}
	if value, ok := ccu.mutation.Quality(); ok {
		_spec.SetField(coursecomment.FieldQuality, field.TypeInt8, value)
	}
	if value, ok := ccu.mutation.AddedQuality(); ok {
		_spec.AddField(coursecomment.FieldQuality, field.TypeInt8, value)
	}
	if value, ok := ccu.mutation.Workload(); ok {
		_spec.SetField(coursecomment.FieldWorkload, field.TypeInt8, value)
	}
	if value, ok := ccu.mutation.AddedWorkload(); ok {
		_spec.AddField(coursecomment.FieldWorkload, field.TypeInt8, value)
	}
	if value, ok := ccu.mutation.Mark(); ok {
		_spec.SetField(coursecomment.FieldMark, field.TypeInt8, value)
	}
	if value, ok := ccu.mutation.AddedMark(); ok {
		_spec.AddField(coursecomment.FieldMark, field.TypeInt8, value)
	}
	if ccu.mutation.CommentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   coursecomment.CommentTable,
			Columns: []string{coursecomment.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccu.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   coursecomment.CommentTable,
			Columns: []string{coursecomment.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ccu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coursecomment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ccu.mutation.done = true
	return n, nil
}

// CourseCommentUpdateOne is the builder for updating a single CourseComment entity.
type CourseCommentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CourseCommentMutation
}

// SetDifficulty sets the "difficulty" field.
func (ccuo *CourseCommentUpdateOne) SetDifficulty(i int8) *CourseCommentUpdateOne {
	ccuo.mutation.ResetDifficulty()
	ccuo.mutation.SetDifficulty(i)
	return ccuo
}

// SetNillableDifficulty sets the "difficulty" field if the given value is not nil.
func (ccuo *CourseCommentUpdateOne) SetNillableDifficulty(i *int8) *CourseCommentUpdateOne {
	if i != nil {
		ccuo.SetDifficulty(*i)
	}
	return ccuo
}

// AddDifficulty adds i to the "difficulty" field.
func (ccuo *CourseCommentUpdateOne) AddDifficulty(i int8) *CourseCommentUpdateOne {
	ccuo.mutation.AddDifficulty(i)
	return ccuo
}

// SetQuality sets the "quality" field.
func (ccuo *CourseCommentUpdateOne) SetQuality(i int8) *CourseCommentUpdateOne {
	ccuo.mutation.ResetQuality()
	ccuo.mutation.SetQuality(i)
	return ccuo
}

// SetNillableQuality sets the "quality" field if the given value is not nil.
func (ccuo *CourseCommentUpdateOne) SetNillableQuality(i *int8) *CourseCommentUpdateOne {
	if i != nil {
		ccuo.SetQuality(*i)
	}
	return ccuo
}

// AddQuality adds i to the "quality" field.
func (ccuo *CourseCommentUpdateOne) AddQuality(i int8) *CourseCommentUpdateOne {
	ccuo.mutation.AddQuality(i)
	return ccuo
}

// SetWorkload sets the "workload" field.
func (ccuo *CourseCommentUpdateOne) SetWorkload(i int8) *CourseCommentUpdateOne {
	ccuo.mutation.ResetWorkload()
	ccuo.mutation.SetWorkload(i)
	return ccuo
}

// SetNillableWorkload sets the "workload" field if the given value is not nil.
func (ccuo *CourseCommentUpdateOne) SetNillableWorkload(i *int8) *CourseCommentUpdateOne {
	if i != nil {
		ccuo.SetWorkload(*i)
	}
	return ccuo
}

// AddWorkload adds i to the "workload" field.
func (ccuo *CourseCommentUpdateOne) AddWorkload(i int8) *CourseCommentUpdateOne {
	ccuo.mutation.AddWorkload(i)
	return ccuo
}

// SetMark sets the "mark" field.
func (ccuo *CourseCommentUpdateOne) SetMark(i int8) *CourseCommentUpdateOne {
	ccuo.mutation.ResetMark()
	ccuo.mutation.SetMark(i)
	return ccuo
}

// SetNillableMark sets the "mark" field if the given value is not nil.
func (ccuo *CourseCommentUpdateOne) SetNillableMark(i *int8) *CourseCommentUpdateOne {
	if i != nil {
		ccuo.SetMark(*i)
	}
	return ccuo
}

// AddMark adds i to the "mark" field.
func (ccuo *CourseCommentUpdateOne) AddMark(i int8) *CourseCommentUpdateOne {
	ccuo.mutation.AddMark(i)
	return ccuo
}

// SetCommentID sets the "comment" edge to the Comment entity by ID.
func (ccuo *CourseCommentUpdateOne) SetCommentID(id uuid.UUID) *CourseCommentUpdateOne {
	ccuo.mutation.SetCommentID(id)
	return ccuo
}

// SetComment sets the "comment" edge to the Comment entity.
func (ccuo *CourseCommentUpdateOne) SetComment(c *Comment) *CourseCommentUpdateOne {
	return ccuo.SetCommentID(c.ID)
}

// Mutation returns the CourseCommentMutation object of the builder.
func (ccuo *CourseCommentUpdateOne) Mutation() *CourseCommentMutation {
	return ccuo.mutation
}

// ClearComment clears the "comment" edge to the Comment entity.
func (ccuo *CourseCommentUpdateOne) ClearComment() *CourseCommentUpdateOne {
	ccuo.mutation.ClearComment()
	return ccuo
}

// Where appends a list predicates to the CourseCommentUpdate builder.
func (ccuo *CourseCommentUpdateOne) Where(ps ...predicate.CourseComment) *CourseCommentUpdateOne {
	ccuo.mutation.Where(ps...)
	return ccuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ccuo *CourseCommentUpdateOne) Select(field string, fields ...string) *CourseCommentUpdateOne {
	ccuo.fields = append([]string{field}, fields...)
	return ccuo
}

// Save executes the query and returns the updated CourseComment entity.
func (ccuo *CourseCommentUpdateOne) Save(ctx context.Context) (*CourseComment, error) {
	return withHooks(ctx, ccuo.sqlSave, ccuo.mutation, ccuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ccuo *CourseCommentUpdateOne) SaveX(ctx context.Context) *CourseComment {
	node, err := ccuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ccuo *CourseCommentUpdateOne) Exec(ctx context.Context) error {
	_, err := ccuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccuo *CourseCommentUpdateOne) ExecX(ctx context.Context) {
	if err := ccuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ccuo *CourseCommentUpdateOne) check() error {
	if _, ok := ccuo.mutation.CommentID(); ccuo.mutation.CommentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CourseComment.comment"`)
	}
	return nil
}

func (ccuo *CourseCommentUpdateOne) sqlSave(ctx context.Context) (_node *CourseComment, err error) {
	if err := ccuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(coursecomment.Table, coursecomment.Columns, sqlgraph.NewFieldSpec(coursecomment.FieldID, field.TypeInt))
	id, ok := ccuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CourseComment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ccuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coursecomment.FieldID)
		for _, f := range fields {
			if !coursecomment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != coursecomment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ccuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccuo.mutation.Difficulty(); ok {
		_spec.SetField(coursecomment.FieldDifficulty, field.TypeInt8, value)
	}
	if value, ok := ccuo.mutation.AddedDifficulty(); ok {
		_spec.AddField(coursecomment.FieldDifficulty, field.TypeInt8, value)
	}
	if value, ok := ccuo.mutation.Quality(); ok {
		_spec.SetField(coursecomment.FieldQuality, field.TypeInt8, value)
	}
	if value, ok := ccuo.mutation.AddedQuality(); ok {
		_spec.AddField(coursecomment.FieldQuality, field.TypeInt8, value)
	}
	if value, ok := ccuo.mutation.Workload(); ok {
		_spec.SetField(coursecomment.FieldWorkload, field.TypeInt8, value)
	}
	if value, ok := ccuo.mutation.AddedWorkload(); ok {
		_spec.AddField(coursecomment.FieldWorkload, field.TypeInt8, value)
	}
	if value, ok := ccuo.mutation.Mark(); ok {
		_spec.SetField(coursecomment.FieldMark, field.TypeInt8, value)
	}
	if value, ok := ccuo.mutation.AddedMark(); ok {
		_spec.AddField(coursecomment.FieldMark, field.TypeInt8, value)
	}
	if ccuo.mutation.CommentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   coursecomment.CommentTable,
			Columns: []string{coursecomment.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccuo.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   coursecomment.CommentTable,
			Columns: []string{coursecomment.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CourseComment{config: ccuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ccuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coursecomment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ccuo.mutation.done = true
	return _node, nil
}