// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/interngowhere/web-backend/ent/comment"
	"github.com/interngowhere/web-backend/ent/predicate"
	"github.com/interngowhere/web-backend/ent/tag"
	"github.com/interngowhere/web-backend/ent/thread"
	"github.com/interngowhere/web-backend/ent/threadkudo"
	"github.com/interngowhere/web-backend/ent/topic"
	"github.com/interngowhere/web-backend/ent/user"
)

// ThreadQuery is the builder for querying Thread entities.
type ThreadQuery struct {
	config
	ctx                *QueryContext
	order              []thread.OrderOption
	inters             []Interceptor
	predicates         []predicate.Thread
	withThreadComments *CommentQuery
	withTags           *TagQuery
	withTopics         *TopicQuery
	withUsers          *UserQuery
	withKudoedUsers    *UserQuery
	withThreadKudoes   *ThreadKudoQuery
	withFKs            bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ThreadQuery builder.
func (tq *ThreadQuery) Where(ps ...predicate.Thread) *ThreadQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *ThreadQuery) Limit(limit int) *ThreadQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *ThreadQuery) Offset(offset int) *ThreadQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *ThreadQuery) Unique(unique bool) *ThreadQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *ThreadQuery) Order(o ...thread.OrderOption) *ThreadQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryThreadComments chains the current query on the "thread_comments" edge.
func (tq *ThreadQuery) QueryThreadComments() *CommentQuery {
	query := (&CommentClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(thread.Table, thread.FieldID, selector),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, thread.ThreadCommentsTable, thread.ThreadCommentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTags chains the current query on the "tags" edge.
func (tq *ThreadQuery) QueryTags() *TagQuery {
	query := (&TagClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(thread.Table, thread.FieldID, selector),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, thread.TagsTable, thread.TagsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTopics chains the current query on the "topics" edge.
func (tq *ThreadQuery) QueryTopics() *TopicQuery {
	query := (&TopicClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(thread.Table, thread.FieldID, selector),
			sqlgraph.To(topic.Table, topic.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, thread.TopicsTable, thread.TopicsColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUsers chains the current query on the "users" edge.
func (tq *ThreadQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(thread.Table, thread.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, thread.UsersTable, thread.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryKudoedUsers chains the current query on the "kudoed_users" edge.
func (tq *ThreadQuery) QueryKudoedUsers() *UserQuery {
	query := (&UserClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(thread.Table, thread.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, thread.KudoedUsersTable, thread.KudoedUsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryThreadKudoes chains the current query on the "thread_kudoes" edge.
func (tq *ThreadQuery) QueryThreadKudoes() *ThreadKudoQuery {
	query := (&ThreadKudoClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(thread.Table, thread.FieldID, selector),
			sqlgraph.To(threadkudo.Table, threadkudo.ThreadColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, thread.ThreadKudoesTable, thread.ThreadKudoesColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Thread entity from the query.
// Returns a *NotFoundError when no Thread was found.
func (tq *ThreadQuery) First(ctx context.Context) (*Thread, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{thread.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *ThreadQuery) FirstX(ctx context.Context) *Thread {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Thread ID from the query.
// Returns a *NotFoundError when no Thread ID was found.
func (tq *ThreadQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{thread.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *ThreadQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Thread entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Thread entity is found.
// Returns a *NotFoundError when no Thread entities are found.
func (tq *ThreadQuery) Only(ctx context.Context) (*Thread, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{thread.Label}
	default:
		return nil, &NotSingularError{thread.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *ThreadQuery) OnlyX(ctx context.Context) *Thread {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Thread ID in the query.
// Returns a *NotSingularError when more than one Thread ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *ThreadQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{thread.Label}
	default:
		err = &NotSingularError{thread.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *ThreadQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Threads.
func (tq *ThreadQuery) All(ctx context.Context) ([]*Thread, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Thread, *ThreadQuery]()
	return withInterceptors[[]*Thread](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *ThreadQuery) AllX(ctx context.Context) []*Thread {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Thread IDs.
func (tq *ThreadQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err = tq.Select(thread.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *ThreadQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *ThreadQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*ThreadQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *ThreadQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *ThreadQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, "Exist")
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *ThreadQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ThreadQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *ThreadQuery) Clone() *ThreadQuery {
	if tq == nil {
		return nil
	}
	return &ThreadQuery{
		config:             tq.config,
		ctx:                tq.ctx.Clone(),
		order:              append([]thread.OrderOption{}, tq.order...),
		inters:             append([]Interceptor{}, tq.inters...),
		predicates:         append([]predicate.Thread{}, tq.predicates...),
		withThreadComments: tq.withThreadComments.Clone(),
		withTags:           tq.withTags.Clone(),
		withTopics:         tq.withTopics.Clone(),
		withUsers:          tq.withUsers.Clone(),
		withKudoedUsers:    tq.withKudoedUsers.Clone(),
		withThreadKudoes:   tq.withThreadKudoes.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithThreadComments tells the query-builder to eager-load the nodes that are connected to
// the "thread_comments" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *ThreadQuery) WithThreadComments(opts ...func(*CommentQuery)) *ThreadQuery {
	query := (&CommentClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withThreadComments = query
	return tq
}

// WithTags tells the query-builder to eager-load the nodes that are connected to
// the "tags" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *ThreadQuery) WithTags(opts ...func(*TagQuery)) *ThreadQuery {
	query := (&TagClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withTags = query
	return tq
}

// WithTopics tells the query-builder to eager-load the nodes that are connected to
// the "topics" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *ThreadQuery) WithTopics(opts ...func(*TopicQuery)) *ThreadQuery {
	query := (&TopicClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withTopics = query
	return tq
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *ThreadQuery) WithUsers(opts ...func(*UserQuery)) *ThreadQuery {
	query := (&UserClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withUsers = query
	return tq
}

// WithKudoedUsers tells the query-builder to eager-load the nodes that are connected to
// the "kudoed_users" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *ThreadQuery) WithKudoedUsers(opts ...func(*UserQuery)) *ThreadQuery {
	query := (&UserClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withKudoedUsers = query
	return tq
}

// WithThreadKudoes tells the query-builder to eager-load the nodes that are connected to
// the "thread_kudoes" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *ThreadQuery) WithThreadKudoes(opts ...func(*ThreadKudoQuery)) *ThreadQuery {
	query := (&ThreadKudoClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withThreadKudoes = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Thread.Query().
//		GroupBy(thread.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *ThreadQuery) GroupBy(field string, fields ...string) *ThreadGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ThreadGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = thread.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.Thread.Query().
//		Select(thread.FieldTitle).
//		Scan(ctx, &v)
func (tq *ThreadQuery) Select(fields ...string) *ThreadSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &ThreadSelect{ThreadQuery: tq}
	sbuild.label = thread.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ThreadSelect configured with the given aggregations.
func (tq *ThreadQuery) Aggregate(fns ...AggregateFunc) *ThreadSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *ThreadQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !thread.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *ThreadQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Thread, error) {
	var (
		nodes       = []*Thread{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [6]bool{
			tq.withThreadComments != nil,
			tq.withTags != nil,
			tq.withTopics != nil,
			tq.withUsers != nil,
			tq.withKudoedUsers != nil,
			tq.withThreadKudoes != nil,
		}
	)
	if tq.withTopics != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, thread.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Thread).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Thread{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withThreadComments; query != nil {
		if err := tq.loadThreadComments(ctx, query, nodes,
			func(n *Thread) { n.Edges.ThreadComments = []*Comment{} },
			func(n *Thread, e *Comment) { n.Edges.ThreadComments = append(n.Edges.ThreadComments, e) }); err != nil {
			return nil, err
		}
	}
	if query := tq.withTags; query != nil {
		if err := tq.loadTags(ctx, query, nodes,
			func(n *Thread) { n.Edges.Tags = []*Tag{} },
			func(n *Thread, e *Tag) { n.Edges.Tags = append(n.Edges.Tags, e) }); err != nil {
			return nil, err
		}
	}
	if query := tq.withTopics; query != nil {
		if err := tq.loadTopics(ctx, query, nodes, nil,
			func(n *Thread, e *Topic) { n.Edges.Topics = e }); err != nil {
			return nil, err
		}
	}
	if query := tq.withUsers; query != nil {
		if err := tq.loadUsers(ctx, query, nodes, nil,
			func(n *Thread, e *User) { n.Edges.Users = e }); err != nil {
			return nil, err
		}
	}
	if query := tq.withKudoedUsers; query != nil {
		if err := tq.loadKudoedUsers(ctx, query, nodes,
			func(n *Thread) { n.Edges.KudoedUsers = []*User{} },
			func(n *Thread, e *User) { n.Edges.KudoedUsers = append(n.Edges.KudoedUsers, e) }); err != nil {
			return nil, err
		}
	}
	if query := tq.withThreadKudoes; query != nil {
		if err := tq.loadThreadKudoes(ctx, query, nodes,
			func(n *Thread) { n.Edges.ThreadKudoes = []*ThreadKudo{} },
			func(n *Thread, e *ThreadKudo) { n.Edges.ThreadKudoes = append(n.Edges.ThreadKudoes, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *ThreadQuery) loadThreadComments(ctx context.Context, query *CommentQuery, nodes []*Thread, init func(*Thread), assign func(*Thread, *Comment)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Thread)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(thread.ThreadCommentsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.thread_thread_comments
		if fk == nil {
			return fmt.Errorf(`foreign-key "thread_thread_comments" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "thread_thread_comments" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (tq *ThreadQuery) loadTags(ctx context.Context, query *TagQuery, nodes []*Thread, init func(*Thread), assign func(*Thread, *Tag)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Thread)
	nids := make(map[int]map[*Thread]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(thread.TagsTable)
		s.Join(joinT).On(s.C(tag.FieldID), joinT.C(thread.TagsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(thread.TagsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(thread.TagsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Thread]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Tag](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "tags" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (tq *ThreadQuery) loadTopics(ctx context.Context, query *TopicQuery, nodes []*Thread, init func(*Thread), assign func(*Thread, *Topic)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Thread)
	for i := range nodes {
		if nodes[i].topic_topic_threads == nil {
			continue
		}
		fk := *nodes[i].topic_topic_threads
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(topic.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "topic_topic_threads" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tq *ThreadQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*Thread, init func(*Thread), assign func(*Thread, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Thread)
	for i := range nodes {
		fk := nodes[i].CreatedBy
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
			return fmt.Errorf(`unexpected foreign-key "created_by" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tq *ThreadQuery) loadKudoedUsers(ctx context.Context, query *UserQuery, nodes []*Thread, init func(*Thread), assign func(*Thread, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Thread)
	nids := make(map[uuid.UUID]map[*Thread]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(thread.KudoedUsersTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(thread.KudoedUsersPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(thread.KudoedUsersPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(thread.KudoedUsersPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Thread]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "kudoed_users" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (tq *ThreadQuery) loadThreadKudoes(ctx context.Context, query *ThreadKudoQuery, nodes []*Thread, init func(*Thread), assign func(*Thread, *ThreadKudo)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Thread)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(threadkudo.FieldThreadID)
	}
	query.Where(predicate.ThreadKudo(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(thread.ThreadKudoesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ThreadID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "thread_id" returned %v for node %v`, fk, n)
		}
		assign(node, n)
	}
	return nil
}

func (tq *ThreadQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *ThreadQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(thread.Table, thread.Columns, sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, thread.FieldID)
		for i := range fields {
			if fields[i] != thread.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if tq.withUsers != nil {
			_spec.Node.AddColumnOnce(thread.FieldCreatedBy)
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *ThreadQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(thread.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = thread.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ThreadGroupBy is the group-by builder for Thread entities.
type ThreadGroupBy struct {
	selector
	build *ThreadQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *ThreadGroupBy) Aggregate(fns ...AggregateFunc) *ThreadGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *ThreadGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ThreadQuery, *ThreadGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *ThreadGroupBy) sqlScan(ctx context.Context, root *ThreadQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ThreadSelect is the builder for selecting fields of Thread entities.
type ThreadSelect struct {
	*ThreadQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *ThreadSelect) Aggregate(fns ...AggregateFunc) *ThreadSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *ThreadSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ThreadQuery, *ThreadSelect](ctx, ts.ThreadQuery, ts, ts.inters, v)
}

func (ts *ThreadSelect) sqlScan(ctx context.Context, root *ThreadQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
