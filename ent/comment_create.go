// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"dgut-icourse/ent/comment"
	"dgut-icourse/ent/coursecomment"
	"dgut-icourse/ent/object"
	"dgut-icourse/ent/user"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CommentCreate is the builder for creating a Comment entity.
type CommentCreate struct {
	config
	mutation *CommentMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *CommentCreate) SetCreatedAt(i int64) *CommentCreate {
	cc.mutation.SetCreatedAt(i)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CommentCreate) SetNillableCreatedAt(i *int64) *CommentCreate {
	if i != nil {
		cc.SetCreatedAt(*i)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CommentCreate) SetUpdatedAt(i int64) *CommentCreate {
	cc.mutation.SetUpdatedAt(i)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CommentCreate) SetNillableUpdatedAt(i *int64) *CommentCreate {
	if i != nil {
		cc.SetUpdatedAt(*i)
	}
	return cc
}

// SetContent sets the "content" field.
func (cc *CommentCreate) SetContent(s string) *CommentCreate {
	cc.mutation.SetContent(s)
	return cc
}

// SetScore sets the "score" field.
func (cc *CommentCreate) SetScore(i int) *CommentCreate {
	cc.mutation.SetScore(i)
	return cc
}

// SetID sets the "id" field.
func (cc *CommentCreate) SetID(u uuid.UUID) *CommentCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *CommentCreate) SetNillableID(u *uuid.UUID) *CommentCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (cc *CommentCreate) SetAuthorID(id uuid.UUID) *CommentCreate {
	cc.mutation.SetAuthorID(id)
	return cc
}

// SetAuthor sets the "author" edge to the User entity.
func (cc *CommentCreate) SetAuthor(u *User) *CommentCreate {
	return cc.SetAuthorID(u.ID)
}

// SetObjectID sets the "object" edge to the Object entity by ID.
func (cc *CommentCreate) SetObjectID(id uuid.UUID) *CommentCreate {
	cc.mutation.SetObjectID(id)
	return cc
}

// SetNillableObjectID sets the "object" edge to the Object entity by ID if the given value is not nil.
func (cc *CommentCreate) SetNillableObjectID(id *uuid.UUID) *CommentCreate {
	if id != nil {
		cc = cc.SetObjectID(*id)
	}
	return cc
}

// SetObject sets the "object" edge to the Object entity.
func (cc *CommentCreate) SetObject(o *Object) *CommentCreate {
	return cc.SetObjectID(o.ID)
}

// SetParentID sets the "parent" edge to the Comment entity by ID.
func (cc *CommentCreate) SetParentID(id uuid.UUID) *CommentCreate {
	cc.mutation.SetParentID(id)
	return cc
}

// SetNillableParentID sets the "parent" edge to the Comment entity by ID if the given value is not nil.
func (cc *CommentCreate) SetNillableParentID(id *uuid.UUID) *CommentCreate {
	if id != nil {
		cc = cc.SetParentID(*id)
	}
	return cc
}

// SetParent sets the "parent" edge to the Comment entity.
func (cc *CommentCreate) SetParent(c *Comment) *CommentCreate {
	return cc.SetParentID(c.ID)
}

// AddChildIDs adds the "children" edge to the Comment entity by IDs.
func (cc *CommentCreate) AddChildIDs(ids ...uuid.UUID) *CommentCreate {
	cc.mutation.AddChildIDs(ids...)
	return cc
}

// AddChildren adds the "children" edges to the Comment entity.
func (cc *CommentCreate) AddChildren(c ...*Comment) *CommentCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddChildIDs(ids...)
}

// SetCourseCommentID sets the "course_comment" edge to the CourseComment entity by ID.
func (cc *CommentCreate) SetCourseCommentID(id int) *CommentCreate {
	cc.mutation.SetCourseCommentID(id)
	return cc
}

// SetNillableCourseCommentID sets the "course_comment" edge to the CourseComment entity by ID if the given value is not nil.
func (cc *CommentCreate) SetNillableCourseCommentID(id *int) *CommentCreate {
	if id != nil {
		cc = cc.SetCourseCommentID(*id)
	}
	return cc
}

// SetCourseComment sets the "course_comment" edge to the CourseComment entity.
func (cc *CommentCreate) SetCourseComment(c *CourseComment) *CommentCreate {
	return cc.SetCourseCommentID(c.ID)
}

// AddLikedUserIDs adds the "liked_users" edge to the User entity by IDs.
func (cc *CommentCreate) AddLikedUserIDs(ids ...uuid.UUID) *CommentCreate {
	cc.mutation.AddLikedUserIDs(ids...)
	return cc
}

// AddLikedUsers adds the "liked_users" edges to the User entity.
func (cc *CommentCreate) AddLikedUsers(u ...*User) *CommentCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cc.AddLikedUserIDs(ids...)
}

// Mutation returns the CommentMutation object of the builder.
func (cc *CommentCreate) Mutation() *CommentMutation {
	return cc.mutation
}

// Save creates the Comment in the database.
func (cc *CommentCreate) Save(ctx context.Context) (*Comment, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CommentCreate) SaveX(ctx context.Context) *Comment {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CommentCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CommentCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CommentCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := comment.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := comment.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := comment.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CommentCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Comment.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Comment.updated_at"`)}
	}
	if _, ok := cc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Comment.content"`)}
	}
	if v, ok := cc.mutation.Content(); ok {
		if err := comment.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Comment.content": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Score(); !ok {
		return &ValidationError{Name: "score", err: errors.New(`ent: missing required field "Comment.score"`)}
	}
	if _, ok := cc.mutation.AuthorID(); !ok {
		return &ValidationError{Name: "author", err: errors.New(`ent: missing required edge "Comment.author"`)}
	}
	return nil
}

func (cc *CommentCreate) sqlSave(ctx context.Context) (*Comment, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CommentCreate) createSpec() (*Comment, *sqlgraph.CreateSpec) {
	var (
		_node = &Comment{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(comment.Table, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(comment.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(comment.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.Content(); ok {
		_spec.SetField(comment.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := cc.mutation.Score(); ok {
		_spec.SetField(comment.FieldScore, field.TypeInt, value)
		_node.Score = value
	}
	if nodes := cc.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   comment.AuthorTable,
			Columns: []string{comment.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.comment_author = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ObjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.ObjectTable,
			Columns: []string{comment.ObjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(object.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.object_comments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.ParentTable,
			Columns: []string{comment.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.comment_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.ChildrenTable,
			Columns: []string{comment.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.CourseCommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   comment.CourseCommentTable,
			Columns: []string{comment.CourseCommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(coursecomment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.LikedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comment.LikedUsersTable,
			Columns: comment.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CommentCreateBulk is the builder for creating many Comment entities in bulk.
type CommentCreateBulk struct {
	config
	err      error
	builders []*CommentCreate
}

// Save creates the Comment entities in the database.
func (ccb *CommentCreateBulk) Save(ctx context.Context) ([]*Comment, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Comment, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommentMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CommentCreateBulk) SaveX(ctx context.Context) []*Comment {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CommentCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CommentCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
