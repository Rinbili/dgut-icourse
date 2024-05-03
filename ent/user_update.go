// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"dgut-icourse/ent/comment"
	"dgut-icourse/ent/predicate"
	"dgut-icourse/ent/user"
	"dgut-icourse/ent/userinfo"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(i int64) *UserUpdate {
	uu.mutation.ResetUpdatedAt()
	uu.mutation.SetUpdatedAt(i)
	return uu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (uu *UserUpdate) AddUpdatedAt(i int64) *UserUpdate {
	uu.mutation.AddUpdatedAt(i)
	return uu
}

// SetNickName sets the "nickName" field.
func (uu *UserUpdate) SetNickName(s string) *UserUpdate {
	uu.mutation.SetNickName(s)
	return uu
}

// SetNillableNickName sets the "nickName" field if the given value is not nil.
func (uu *UserUpdate) SetNillableNickName(s *string) *UserUpdate {
	if s != nil {
		uu.SetNickName(*s)
	}
	return uu
}

// ClearNickName clears the value of the "nickName" field.
func (uu *UserUpdate) ClearNickName() *UserUpdate {
	uu.mutation.ClearNickName()
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// ClearPassword clears the value of the "password" field.
func (uu *UserUpdate) ClearPassword() *UserUpdate {
	uu.mutation.ClearPassword()
	return uu
}

// SetPhone sets the "phone" field.
func (uu *UserUpdate) SetPhone(s string) *UserUpdate {
	uu.mutation.SetPhone(s)
	return uu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePhone(s *string) *UserUpdate {
	if s != nil {
		uu.SetPhone(*s)
	}
	return uu
}

// ClearPhone clears the value of the "phone" field.
func (uu *UserUpdate) ClearPhone() *UserUpdate {
	uu.mutation.ClearPhone()
	return uu
}

// SetIcon sets the "icon" field.
func (uu *UserUpdate) SetIcon(s string) *UserUpdate {
	uu.mutation.SetIcon(s)
	return uu
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIcon(s *string) *UserUpdate {
	if s != nil {
		uu.SetIcon(*s)
	}
	return uu
}

// ClearIcon clears the value of the "icon" field.
func (uu *UserUpdate) ClearIcon() *UserUpdate {
	uu.mutation.ClearIcon()
	return uu
}

// SetOpenid sets the "openid" field.
func (uu *UserUpdate) SetOpenid(s string) *UserUpdate {
	uu.mutation.SetOpenid(s)
	return uu
}

// SetNillableOpenid sets the "openid" field if the given value is not nil.
func (uu *UserUpdate) SetNillableOpenid(s *string) *UserUpdate {
	if s != nil {
		uu.SetOpenid(*s)
	}
	return uu
}

// SetInfoID sets the "info" edge to the UserInfo entity by ID.
func (uu *UserUpdate) SetInfoID(id int) *UserUpdate {
	uu.mutation.SetInfoID(id)
	return uu
}

// SetNillableInfoID sets the "info" edge to the UserInfo entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableInfoID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetInfoID(*id)
	}
	return uu
}

// SetInfo sets the "info" edge to the UserInfo entity.
func (uu *UserUpdate) SetInfo(u *UserInfo) *UserUpdate {
	return uu.SetInfoID(u.ID)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (uu *UserUpdate) AddCommentIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddCommentIDs(ids...)
	return uu
}

// AddComments adds the "comments" edges to the Comment entity.
func (uu *UserUpdate) AddComments(c ...*Comment) *UserUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddCommentIDs(ids...)
}

// AddLikedCommentIDs adds the "liked_comments" edge to the Comment entity by IDs.
func (uu *UserUpdate) AddLikedCommentIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddLikedCommentIDs(ids...)
	return uu
}

// AddLikedComments adds the "liked_comments" edges to the Comment entity.
func (uu *UserUpdate) AddLikedComments(c ...*Comment) *UserUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddLikedCommentIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearInfo clears the "info" edge to the UserInfo entity.
func (uu *UserUpdate) ClearInfo() *UserUpdate {
	uu.mutation.ClearInfo()
	return uu
}

// ClearComments clears all "comments" edges to the Comment entity.
func (uu *UserUpdate) ClearComments() *UserUpdate {
	uu.mutation.ClearComments()
	return uu
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (uu *UserUpdate) RemoveCommentIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveCommentIDs(ids...)
	return uu
}

// RemoveComments removes "comments" edges to Comment entities.
func (uu *UserUpdate) RemoveComments(c ...*Comment) *UserUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveCommentIDs(ids...)
}

// ClearLikedComments clears all "liked_comments" edges to the Comment entity.
func (uu *UserUpdate) ClearLikedComments() *UserUpdate {
	uu.mutation.ClearLikedComments()
	return uu
}

// RemoveLikedCommentIDs removes the "liked_comments" edge to Comment entities by IDs.
func (uu *UserUpdate) RemoveLikedCommentIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveLikedCommentIDs(ids...)
	return uu
}

// RemoveLikedComments removes "liked_comments" edges to Comment entities.
func (uu *UserUpdate) RemoveLikedComments(c ...*Comment) *UserUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveLikedCommentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.NickName(); ok {
		if err := user.NickNameValidator(v); err != nil {
			return &ValidationError{Name: "nickName", err: fmt.Errorf(`ent: validator failed for field "User.nickName": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Phone(); ok {
		if err := user.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "User.phone": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Openid(); ok {
		if err := user.OpenidValidator(v); err != nil {
			return &ValidationError{Name: "openid", err: fmt.Errorf(`ent: validator failed for field "User.openid": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(user.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.NickName(); ok {
		_spec.SetField(user.FieldNickName, field.TypeString, value)
	}
	if uu.mutation.NickNameCleared() {
		_spec.ClearField(user.FieldNickName, field.TypeString)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if uu.mutation.PasswordCleared() {
		_spec.ClearField(user.FieldPassword, field.TypeString)
	}
	if value, ok := uu.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
	}
	if uu.mutation.PhoneCleared() {
		_spec.ClearField(user.FieldPhone, field.TypeString)
	}
	if value, ok := uu.mutation.Icon(); ok {
		_spec.SetField(user.FieldIcon, field.TypeString, value)
	}
	if uu.mutation.IconCleared() {
		_spec.ClearField(user.FieldIcon, field.TypeString)
	}
	if value, ok := uu.mutation.Openid(); ok {
		_spec.SetField(user.FieldOpenid, field.TypeString, value)
	}
	if uu.mutation.InfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.InfoTable,
			Columns: []string{user.InfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userinfo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.InfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.InfoTable,
			Columns: []string{user.InfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userinfo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.CommentsTable,
			Columns: []string{user.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !uu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.CommentsTable,
			Columns: []string{user.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.CommentsTable,
			Columns: []string{user.CommentsColumn},
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
	if uu.mutation.LikedCommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.LikedCommentsTable,
			Columns: user.LikedCommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedLikedCommentsIDs(); len(nodes) > 0 && !uu.mutation.LikedCommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.LikedCommentsTable,
			Columns: user.LikedCommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.LikedCommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.LikedCommentsTable,
			Columns: user.LikedCommentsPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(i int64) *UserUpdateOne {
	uuo.mutation.ResetUpdatedAt()
	uuo.mutation.SetUpdatedAt(i)
	return uuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (uuo *UserUpdateOne) AddUpdatedAt(i int64) *UserUpdateOne {
	uuo.mutation.AddUpdatedAt(i)
	return uuo
}

// SetNickName sets the "nickName" field.
func (uuo *UserUpdateOne) SetNickName(s string) *UserUpdateOne {
	uuo.mutation.SetNickName(s)
	return uuo
}

// SetNillableNickName sets the "nickName" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableNickName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetNickName(*s)
	}
	return uuo
}

// ClearNickName clears the value of the "nickName" field.
func (uuo *UserUpdateOne) ClearNickName() *UserUpdateOne {
	uuo.mutation.ClearNickName()
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// ClearPassword clears the value of the "password" field.
func (uuo *UserUpdateOne) ClearPassword() *UserUpdateOne {
	uuo.mutation.ClearPassword()
	return uuo
}

// SetPhone sets the "phone" field.
func (uuo *UserUpdateOne) SetPhone(s string) *UserUpdateOne {
	uuo.mutation.SetPhone(s)
	return uuo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePhone(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPhone(*s)
	}
	return uuo
}

// ClearPhone clears the value of the "phone" field.
func (uuo *UserUpdateOne) ClearPhone() *UserUpdateOne {
	uuo.mutation.ClearPhone()
	return uuo
}

// SetIcon sets the "icon" field.
func (uuo *UserUpdateOne) SetIcon(s string) *UserUpdateOne {
	uuo.mutation.SetIcon(s)
	return uuo
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIcon(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetIcon(*s)
	}
	return uuo
}

// ClearIcon clears the value of the "icon" field.
func (uuo *UserUpdateOne) ClearIcon() *UserUpdateOne {
	uuo.mutation.ClearIcon()
	return uuo
}

// SetOpenid sets the "openid" field.
func (uuo *UserUpdateOne) SetOpenid(s string) *UserUpdateOne {
	uuo.mutation.SetOpenid(s)
	return uuo
}

// SetNillableOpenid sets the "openid" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableOpenid(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetOpenid(*s)
	}
	return uuo
}

// SetInfoID sets the "info" edge to the UserInfo entity by ID.
func (uuo *UserUpdateOne) SetInfoID(id int) *UserUpdateOne {
	uuo.mutation.SetInfoID(id)
	return uuo
}

// SetNillableInfoID sets the "info" edge to the UserInfo entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableInfoID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetInfoID(*id)
	}
	return uuo
}

// SetInfo sets the "info" edge to the UserInfo entity.
func (uuo *UserUpdateOne) SetInfo(u *UserInfo) *UserUpdateOne {
	return uuo.SetInfoID(u.ID)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (uuo *UserUpdateOne) AddCommentIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddCommentIDs(ids...)
	return uuo
}

// AddComments adds the "comments" edges to the Comment entity.
func (uuo *UserUpdateOne) AddComments(c ...*Comment) *UserUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddCommentIDs(ids...)
}

// AddLikedCommentIDs adds the "liked_comments" edge to the Comment entity by IDs.
func (uuo *UserUpdateOne) AddLikedCommentIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddLikedCommentIDs(ids...)
	return uuo
}

// AddLikedComments adds the "liked_comments" edges to the Comment entity.
func (uuo *UserUpdateOne) AddLikedComments(c ...*Comment) *UserUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddLikedCommentIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearInfo clears the "info" edge to the UserInfo entity.
func (uuo *UserUpdateOne) ClearInfo() *UserUpdateOne {
	uuo.mutation.ClearInfo()
	return uuo
}

// ClearComments clears all "comments" edges to the Comment entity.
func (uuo *UserUpdateOne) ClearComments() *UserUpdateOne {
	uuo.mutation.ClearComments()
	return uuo
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (uuo *UserUpdateOne) RemoveCommentIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveCommentIDs(ids...)
	return uuo
}

// RemoveComments removes "comments" edges to Comment entities.
func (uuo *UserUpdateOne) RemoveComments(c ...*Comment) *UserUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveCommentIDs(ids...)
}

// ClearLikedComments clears all "liked_comments" edges to the Comment entity.
func (uuo *UserUpdateOne) ClearLikedComments() *UserUpdateOne {
	uuo.mutation.ClearLikedComments()
	return uuo
}

// RemoveLikedCommentIDs removes the "liked_comments" edge to Comment entities by IDs.
func (uuo *UserUpdateOne) RemoveLikedCommentIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveLikedCommentIDs(ids...)
	return uuo
}

// RemoveLikedComments removes "liked_comments" edges to Comment entities.
func (uuo *UserUpdateOne) RemoveLikedComments(c ...*Comment) *UserUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveLikedCommentIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.NickName(); ok {
		if err := user.NickNameValidator(v); err != nil {
			return &ValidationError{Name: "nickName", err: fmt.Errorf(`ent: validator failed for field "User.nickName": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Phone(); ok {
		if err := user.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "User.phone": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Openid(); ok {
		if err := user.OpenidValidator(v); err != nil {
			return &ValidationError{Name: "openid", err: fmt.Errorf(`ent: validator failed for field "User.openid": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(user.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.NickName(); ok {
		_spec.SetField(user.FieldNickName, field.TypeString, value)
	}
	if uuo.mutation.NickNameCleared() {
		_spec.ClearField(user.FieldNickName, field.TypeString)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if uuo.mutation.PasswordCleared() {
		_spec.ClearField(user.FieldPassword, field.TypeString)
	}
	if value, ok := uuo.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
	}
	if uuo.mutation.PhoneCleared() {
		_spec.ClearField(user.FieldPhone, field.TypeString)
	}
	if value, ok := uuo.mutation.Icon(); ok {
		_spec.SetField(user.FieldIcon, field.TypeString, value)
	}
	if uuo.mutation.IconCleared() {
		_spec.ClearField(user.FieldIcon, field.TypeString)
	}
	if value, ok := uuo.mutation.Openid(); ok {
		_spec.SetField(user.FieldOpenid, field.TypeString, value)
	}
	if uuo.mutation.InfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.InfoTable,
			Columns: []string{user.InfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userinfo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.InfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.InfoTable,
			Columns: []string{user.InfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userinfo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.CommentsTable,
			Columns: []string{user.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !uuo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.CommentsTable,
			Columns: []string{user.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.CommentsTable,
			Columns: []string{user.CommentsColumn},
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
	if uuo.mutation.LikedCommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.LikedCommentsTable,
			Columns: user.LikedCommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedLikedCommentsIDs(); len(nodes) > 0 && !uuo.mutation.LikedCommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.LikedCommentsTable,
			Columns: user.LikedCommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.LikedCommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.LikedCommentsTable,
			Columns: user.LikedCommentsPrimaryKey,
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
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
