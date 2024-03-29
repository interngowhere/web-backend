// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/interngowhere/web-backend/ent/comment"
	"github.com/interngowhere/web-backend/ent/commentkudo"
	"github.com/interngowhere/web-backend/ent/predicate"
	"github.com/interngowhere/web-backend/ent/user"
)

// CommentKudoQuery is the builder for querying CommentKudo entities.
type CommentKudoQuery struct {
	config
	ctx         *QueryContext
	order       []commentkudo.OrderOption
	inters      []Interceptor
	predicates  []predicate.CommentKudo
	withUser    *UserQuery
	withComment *CommentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CommentKudoQuery builder.
func (ckq *CommentKudoQuery) Where(ps ...predicate.CommentKudo) *CommentKudoQuery {
	ckq.predicates = append(ckq.predicates, ps...)
	return ckq
}

// Limit the number of records to be returned by this query.
func (ckq *CommentKudoQuery) Limit(limit int) *CommentKudoQuery {
	ckq.ctx.Limit = &limit
	return ckq
}

// Offset to start from.
func (ckq *CommentKudoQuery) Offset(offset int) *CommentKudoQuery {
	ckq.ctx.Offset = &offset
	return ckq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ckq *CommentKudoQuery) Unique(unique bool) *CommentKudoQuery {
	ckq.ctx.Unique = &unique
	return ckq
}

// Order specifies how the records should be ordered.
func (ckq *CommentKudoQuery) Order(o ...commentkudo.OrderOption) *CommentKudoQuery {
	ckq.order = append(ckq.order, o...)
	return ckq
}

// QueryUser chains the current query on the "user" edge.
func (ckq *CommentKudoQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: ckq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ckq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ckq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(commentkudo.Table, commentkudo.UserColumn, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, commentkudo.UserTable, commentkudo.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ckq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryComment chains the current query on the "comment" edge.
func (ckq *CommentKudoQuery) QueryComment() *CommentQuery {
	query := (&CommentClient{config: ckq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ckq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ckq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(commentkudo.Table, commentkudo.CommentColumn, selector),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, commentkudo.CommentTable, commentkudo.CommentColumn),
		)
		fromU = sqlgraph.SetNeighbors(ckq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CommentKudo entity from the query.
// Returns a *NotFoundError when no CommentKudo was found.
func (ckq *CommentKudoQuery) First(ctx context.Context) (*CommentKudo, error) {
	nodes, err := ckq.Limit(1).All(setContextOp(ctx, ckq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{commentkudo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ckq *CommentKudoQuery) FirstX(ctx context.Context) *CommentKudo {
	node, err := ckq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single CommentKudo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CommentKudo entity is found.
// Returns a *NotFoundError when no CommentKudo entities are found.
func (ckq *CommentKudoQuery) Only(ctx context.Context) (*CommentKudo, error) {
	nodes, err := ckq.Limit(2).All(setContextOp(ctx, ckq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{commentkudo.Label}
	default:
		return nil, &NotSingularError{commentkudo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ckq *CommentKudoQuery) OnlyX(ctx context.Context) *CommentKudo {
	node, err := ckq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of CommentKudos.
func (ckq *CommentKudoQuery) All(ctx context.Context) ([]*CommentKudo, error) {
	ctx = setContextOp(ctx, ckq.ctx, "All")
	if err := ckq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CommentKudo, *CommentKudoQuery]()
	return withInterceptors[[]*CommentKudo](ctx, ckq, qr, ckq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ckq *CommentKudoQuery) AllX(ctx context.Context) []*CommentKudo {
	nodes, err := ckq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (ckq *CommentKudoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ckq.ctx, "Count")
	if err := ckq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ckq, querierCount[*CommentKudoQuery](), ckq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ckq *CommentKudoQuery) CountX(ctx context.Context) int {
	count, err := ckq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ckq *CommentKudoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ckq.ctx, "Exist")
	switch _, err := ckq.First(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ckq *CommentKudoQuery) ExistX(ctx context.Context) bool {
	exist, err := ckq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CommentKudoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ckq *CommentKudoQuery) Clone() *CommentKudoQuery {
	if ckq == nil {
		return nil
	}
	return &CommentKudoQuery{
		config:      ckq.config,
		ctx:         ckq.ctx.Clone(),
		order:       append([]commentkudo.OrderOption{}, ckq.order...),
		inters:      append([]Interceptor{}, ckq.inters...),
		predicates:  append([]predicate.CommentKudo{}, ckq.predicates...),
		withUser:    ckq.withUser.Clone(),
		withComment: ckq.withComment.Clone(),
		// clone intermediate query.
		sql:  ckq.sql.Clone(),
		path: ckq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ckq *CommentKudoQuery) WithUser(opts ...func(*UserQuery)) *CommentKudoQuery {
	query := (&UserClient{config: ckq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ckq.withUser = query
	return ckq
}

// WithComment tells the query-builder to eager-load the nodes that are connected to
// the "comment" edge. The optional arguments are used to configure the query builder of the edge.
func (ckq *CommentKudoQuery) WithComment(opts ...func(*CommentQuery)) *CommentKudoQuery {
	query := (&CommentClient{config: ckq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ckq.withComment = query
	return ckq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CommentKudo.Query().
//		GroupBy(commentkudo.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ckq *CommentKudoQuery) GroupBy(field string, fields ...string) *CommentKudoGroupBy {
	ckq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CommentKudoGroupBy{build: ckq}
	grbuild.flds = &ckq.ctx.Fields
	grbuild.label = commentkudo.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//	}
//
//	client.CommentKudo.Query().
//		Select(commentkudo.FieldUserID).
//		Scan(ctx, &v)
func (ckq *CommentKudoQuery) Select(fields ...string) *CommentKudoSelect {
	ckq.ctx.Fields = append(ckq.ctx.Fields, fields...)
	sbuild := &CommentKudoSelect{CommentKudoQuery: ckq}
	sbuild.label = commentkudo.Label
	sbuild.flds, sbuild.scan = &ckq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CommentKudoSelect configured with the given aggregations.
func (ckq *CommentKudoQuery) Aggregate(fns ...AggregateFunc) *CommentKudoSelect {
	return ckq.Select().Aggregate(fns...)
}

func (ckq *CommentKudoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ckq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ckq); err != nil {
				return err
			}
		}
	}
	for _, f := range ckq.ctx.Fields {
		if !commentkudo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ckq.path != nil {
		prev, err := ckq.path(ctx)
		if err != nil {
			return err
		}
		ckq.sql = prev
	}
	return nil
}

func (ckq *CommentKudoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CommentKudo, error) {
	var (
		nodes       = []*CommentKudo{}
		_spec       = ckq.querySpec()
		loadedTypes = [2]bool{
			ckq.withUser != nil,
			ckq.withComment != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CommentKudo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CommentKudo{config: ckq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ckq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ckq.withUser; query != nil {
		if err := ckq.loadUser(ctx, query, nodes, nil,
			func(n *CommentKudo, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := ckq.withComment; query != nil {
		if err := ckq.loadComment(ctx, query, nodes, nil,
			func(n *CommentKudo, e *Comment) { n.Edges.Comment = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ckq *CommentKudoQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*CommentKudo, init func(*CommentKudo), assign func(*CommentKudo, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*CommentKudo)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ckq *CommentKudoQuery) loadComment(ctx context.Context, query *CommentQuery, nodes []*CommentKudo, init func(*CommentKudo), assign func(*CommentKudo, *Comment)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*CommentKudo)
	for i := range nodes {
		fk := nodes[i].CommentID
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
			return fmt.Errorf(`unexpected foreign-key "comment_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ckq *CommentKudoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ckq.querySpec()
	_spec.Unique = false
	_spec.Node.Columns = nil
	return sqlgraph.CountNodes(ctx, ckq.driver, _spec)
}

func (ckq *CommentKudoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(commentkudo.Table, commentkudo.Columns, nil)
	_spec.From = ckq.sql
	if unique := ckq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ckq.path != nil {
		_spec.Unique = true
	}
	if fields := ckq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
		}
		if ckq.withUser != nil {
			_spec.Node.AddColumnOnce(commentkudo.FieldUserID)
		}
		if ckq.withComment != nil {
			_spec.Node.AddColumnOnce(commentkudo.FieldCommentID)
		}
	}
	if ps := ckq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ckq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ckq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ckq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ckq *CommentKudoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ckq.driver.Dialect())
	t1 := builder.Table(commentkudo.Table)
	columns := ckq.ctx.Fields
	if len(columns) == 0 {
		columns = commentkudo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ckq.sql != nil {
		selector = ckq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ckq.ctx.Unique != nil && *ckq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ckq.predicates {
		p(selector)
	}
	for _, p := range ckq.order {
		p(selector)
	}
	if offset := ckq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ckq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CommentKudoGroupBy is the group-by builder for CommentKudo entities.
type CommentKudoGroupBy struct {
	selector
	build *CommentKudoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ckgb *CommentKudoGroupBy) Aggregate(fns ...AggregateFunc) *CommentKudoGroupBy {
	ckgb.fns = append(ckgb.fns, fns...)
	return ckgb
}

// Scan applies the selector query and scans the result into the given value.
func (ckgb *CommentKudoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ckgb.build.ctx, "GroupBy")
	if err := ckgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommentKudoQuery, *CommentKudoGroupBy](ctx, ckgb.build, ckgb, ckgb.build.inters, v)
}

func (ckgb *CommentKudoGroupBy) sqlScan(ctx context.Context, root *CommentKudoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ckgb.fns))
	for _, fn := range ckgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ckgb.flds)+len(ckgb.fns))
		for _, f := range *ckgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ckgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ckgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CommentKudoSelect is the builder for selecting fields of CommentKudo entities.
type CommentKudoSelect struct {
	*CommentKudoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cks *CommentKudoSelect) Aggregate(fns ...AggregateFunc) *CommentKudoSelect {
	cks.fns = append(cks.fns, fns...)
	return cks
}

// Scan applies the selector query and scans the result into the given value.
func (cks *CommentKudoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cks.ctx, "Select")
	if err := cks.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommentKudoQuery, *CommentKudoSelect](ctx, cks.CommentKudoQuery, cks, cks.inters, v)
}

func (cks *CommentKudoSelect) sqlScan(ctx context.Context, root *CommentKudoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cks.fns))
	for _, fn := range cks.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cks.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cks.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
