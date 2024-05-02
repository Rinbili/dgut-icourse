// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"dgut-icourse/ent/comment"
	"dgut-icourse/ent/coursecomment"
	"dgut-icourse/ent/predicate"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CourseCommentQuery is the builder for querying CourseComment entities.
type CourseCommentQuery struct {
	config
	ctx         *QueryContext
	order       []coursecomment.OrderOption
	inters      []Interceptor
	predicates  []predicate.CourseComment
	withComment *CommentQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CourseCommentQuery builder.
func (ccq *CourseCommentQuery) Where(ps ...predicate.CourseComment) *CourseCommentQuery {
	ccq.predicates = append(ccq.predicates, ps...)
	return ccq
}

// Limit the number of records to be returned by this query.
func (ccq *CourseCommentQuery) Limit(limit int) *CourseCommentQuery {
	ccq.ctx.Limit = &limit
	return ccq
}

// Offset to start from.
func (ccq *CourseCommentQuery) Offset(offset int) *CourseCommentQuery {
	ccq.ctx.Offset = &offset
	return ccq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ccq *CourseCommentQuery) Unique(unique bool) *CourseCommentQuery {
	ccq.ctx.Unique = &unique
	return ccq
}

// Order specifies how the records should be ordered.
func (ccq *CourseCommentQuery) Order(o ...coursecomment.OrderOption) *CourseCommentQuery {
	ccq.order = append(ccq.order, o...)
	return ccq
}

// QueryComment chains the current query on the "comment" edge.
func (ccq *CourseCommentQuery) QueryComment() *CommentQuery {
	query := (&CommentClient{config: ccq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ccq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ccq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(coursecomment.Table, coursecomment.FieldID, selector),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, coursecomment.CommentTable, coursecomment.CommentColumn),
		)
		fromU = sqlgraph.SetNeighbors(ccq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CourseComment entity from the query.
// Returns a *NotFoundError when no CourseComment was found.
func (ccq *CourseCommentQuery) First(ctx context.Context) (*CourseComment, error) {
	nodes, err := ccq.Limit(1).All(setContextOp(ctx, ccq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{coursecomment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ccq *CourseCommentQuery) FirstX(ctx context.Context) *CourseComment {
	node, err := ccq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CourseComment ID from the query.
// Returns a *NotFoundError when no CourseComment ID was found.
func (ccq *CourseCommentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ccq.Limit(1).IDs(setContextOp(ctx, ccq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{coursecomment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ccq *CourseCommentQuery) FirstIDX(ctx context.Context) int {
	id, err := ccq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CourseComment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CourseComment entity is found.
// Returns a *NotFoundError when no CourseComment entities are found.
func (ccq *CourseCommentQuery) Only(ctx context.Context) (*CourseComment, error) {
	nodes, err := ccq.Limit(2).All(setContextOp(ctx, ccq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{coursecomment.Label}
	default:
		return nil, &NotSingularError{coursecomment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ccq *CourseCommentQuery) OnlyX(ctx context.Context) *CourseComment {
	node, err := ccq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CourseComment ID in the query.
// Returns a *NotSingularError when more than one CourseComment ID is found.
// Returns a *NotFoundError when no entities are found.
func (ccq *CourseCommentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ccq.Limit(2).IDs(setContextOp(ctx, ccq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{coursecomment.Label}
	default:
		err = &NotSingularError{coursecomment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ccq *CourseCommentQuery) OnlyIDX(ctx context.Context) int {
	id, err := ccq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CourseComments.
func (ccq *CourseCommentQuery) All(ctx context.Context) ([]*CourseComment, error) {
	ctx = setContextOp(ctx, ccq.ctx, "All")
	if err := ccq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CourseComment, *CourseCommentQuery]()
	return withInterceptors[[]*CourseComment](ctx, ccq, qr, ccq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ccq *CourseCommentQuery) AllX(ctx context.Context) []*CourseComment {
	nodes, err := ccq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CourseComment IDs.
func (ccq *CourseCommentQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ccq.ctx.Unique == nil && ccq.path != nil {
		ccq.Unique(true)
	}
	ctx = setContextOp(ctx, ccq.ctx, "IDs")
	if err = ccq.Select(coursecomment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ccq *CourseCommentQuery) IDsX(ctx context.Context) []int {
	ids, err := ccq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ccq *CourseCommentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ccq.ctx, "Count")
	if err := ccq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ccq, querierCount[*CourseCommentQuery](), ccq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ccq *CourseCommentQuery) CountX(ctx context.Context) int {
	count, err := ccq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ccq *CourseCommentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ccq.ctx, "Exist")
	switch _, err := ccq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ccq *CourseCommentQuery) ExistX(ctx context.Context) bool {
	exist, err := ccq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CourseCommentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ccq *CourseCommentQuery) Clone() *CourseCommentQuery {
	if ccq == nil {
		return nil
	}
	return &CourseCommentQuery{
		config:      ccq.config,
		ctx:         ccq.ctx.Clone(),
		order:       append([]coursecomment.OrderOption{}, ccq.order...),
		inters:      append([]Interceptor{}, ccq.inters...),
		predicates:  append([]predicate.CourseComment{}, ccq.predicates...),
		withComment: ccq.withComment.Clone(),
		// clone intermediate query.
		sql:  ccq.sql.Clone(),
		path: ccq.path,
	}
}

// WithComment tells the query-builder to eager-load the nodes that are connected to
// the "comment" edge. The optional arguments are used to configure the query builder of the edge.
func (ccq *CourseCommentQuery) WithComment(opts ...func(*CommentQuery)) *CourseCommentQuery {
	query := (&CommentClient{config: ccq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ccq.withComment = query
	return ccq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Difficulty int8 `json:"difficulty,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CourseComment.Query().
//		GroupBy(coursecomment.FieldDifficulty).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ccq *CourseCommentQuery) GroupBy(field string, fields ...string) *CourseCommentGroupBy {
	ccq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CourseCommentGroupBy{build: ccq}
	grbuild.flds = &ccq.ctx.Fields
	grbuild.label = coursecomment.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Difficulty int8 `json:"difficulty,omitempty"`
//	}
//
//	client.CourseComment.Query().
//		Select(coursecomment.FieldDifficulty).
//		Scan(ctx, &v)
func (ccq *CourseCommentQuery) Select(fields ...string) *CourseCommentSelect {
	ccq.ctx.Fields = append(ccq.ctx.Fields, fields...)
	sbuild := &CourseCommentSelect{CourseCommentQuery: ccq}
	sbuild.label = coursecomment.Label
	sbuild.flds, sbuild.scan = &ccq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CourseCommentSelect configured with the given aggregations.
func (ccq *CourseCommentQuery) Aggregate(fns ...AggregateFunc) *CourseCommentSelect {
	return ccq.Select().Aggregate(fns...)
}

func (ccq *CourseCommentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ccq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ccq); err != nil {
				return err
			}
		}
	}
	for _, f := range ccq.ctx.Fields {
		if !coursecomment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ccq.path != nil {
		prev, err := ccq.path(ctx)
		if err != nil {
			return err
		}
		ccq.sql = prev
	}
	return nil
}

func (ccq *CourseCommentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CourseComment, error) {
	var (
		nodes       = []*CourseComment{}
		withFKs     = ccq.withFKs
		_spec       = ccq.querySpec()
		loadedTypes = [1]bool{
			ccq.withComment != nil,
		}
	)
	if ccq.withComment != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, coursecomment.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CourseComment).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CourseComment{config: ccq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ccq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ccq.withComment; query != nil {
		if err := ccq.loadComment(ctx, query, nodes, nil,
			func(n *CourseComment, e *Comment) { n.Edges.Comment = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ccq *CourseCommentQuery) loadComment(ctx context.Context, query *CommentQuery, nodes []*CourseComment, init func(*CourseComment), assign func(*CourseComment, *Comment)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*CourseComment)
	for i := range nodes {
		if nodes[i].comment_course_comment == nil {
			continue
		}
		fk := *nodes[i].comment_course_comment
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(comment.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "comment_course_comment" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ccq *CourseCommentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ccq.querySpec()
	_spec.Node.Columns = ccq.ctx.Fields
	if len(ccq.ctx.Fields) > 0 {
		_spec.Unique = ccq.ctx.Unique != nil && *ccq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ccq.driver, _spec)
}

func (ccq *CourseCommentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(coursecomment.Table, coursecomment.Columns, sqlgraph.NewFieldSpec(coursecomment.FieldID, field.TypeInt))
	_spec.From = ccq.sql
	if unique := ccq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ccq.path != nil {
		_spec.Unique = true
	}
	if fields := ccq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coursecomment.FieldID)
		for i := range fields {
			if fields[i] != coursecomment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ccq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ccq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ccq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ccq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ccq *CourseCommentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ccq.driver.Dialect())
	t1 := builder.Table(coursecomment.Table)
	columns := ccq.ctx.Fields
	if len(columns) == 0 {
		columns = coursecomment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ccq.sql != nil {
		selector = ccq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ccq.ctx.Unique != nil && *ccq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ccq.predicates {
		p(selector)
	}
	for _, p := range ccq.order {
		p(selector)
	}
	if offset := ccq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ccq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CourseCommentGroupBy is the group-by builder for CourseComment entities.
type CourseCommentGroupBy struct {
	selector
	build *CourseCommentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ccgb *CourseCommentGroupBy) Aggregate(fns ...AggregateFunc) *CourseCommentGroupBy {
	ccgb.fns = append(ccgb.fns, fns...)
	return ccgb
}

// Scan applies the selector query and scans the result into the given value.
func (ccgb *CourseCommentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ccgb.build.ctx, "GroupBy")
	if err := ccgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CourseCommentQuery, *CourseCommentGroupBy](ctx, ccgb.build, ccgb, ccgb.build.inters, v)
}

func (ccgb *CourseCommentGroupBy) sqlScan(ctx context.Context, root *CourseCommentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ccgb.fns))
	for _, fn := range ccgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ccgb.flds)+len(ccgb.fns))
		for _, f := range *ccgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ccgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ccgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CourseCommentSelect is the builder for selecting fields of CourseComment entities.
type CourseCommentSelect struct {
	*CourseCommentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ccs *CourseCommentSelect) Aggregate(fns ...AggregateFunc) *CourseCommentSelect {
	ccs.fns = append(ccs.fns, fns...)
	return ccs
}

// Scan applies the selector query and scans the result into the given value.
func (ccs *CourseCommentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ccs.ctx, "Select")
	if err := ccs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CourseCommentQuery, *CourseCommentSelect](ctx, ccs.CourseCommentQuery, ccs, ccs.inters, v)
}

func (ccs *CourseCommentSelect) sqlScan(ctx context.Context, root *CourseCommentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ccs.fns))
	for _, fn := range ccs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ccs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ccs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}