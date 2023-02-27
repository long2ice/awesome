// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/long2ice/awesome/ent/predicate"
	"github.com/long2ice/awesome/ent/repo"
	"github.com/long2ice/awesome/ent/topic"
)

// RepoQuery is the builder for querying Repo entities.
type RepoQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.Repo
	withTopics *TopicQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RepoQuery builder.
func (rq *RepoQuery) Where(ps ...predicate.Repo) *RepoQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *RepoQuery) Limit(limit int) *RepoQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *RepoQuery) Offset(offset int) *RepoQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RepoQuery) Unique(unique bool) *RepoQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *RepoQuery) Order(o ...OrderFunc) *RepoQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryTopics chains the current query on the "topics" edge.
func (rq *RepoQuery) QueryTopics() *TopicQuery {
	query := (&TopicClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(repo.Table, repo.FieldID, selector),
			sqlgraph.To(topic.Table, topic.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, repo.TopicsTable, repo.TopicsColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Repo entity from the query.
// Returns a *NotFoundError when no Repo was found.
func (rq *RepoQuery) First(ctx context.Context) (*Repo, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{repo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RepoQuery) FirstX(ctx context.Context) *Repo {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Repo ID from the query.
// Returns a *NotFoundError when no Repo ID was found.
func (rq *RepoQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{repo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *RepoQuery) FirstIDX(ctx context.Context) int {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Repo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Repo entity is found.
// Returns a *NotFoundError when no Repo entities are found.
func (rq *RepoQuery) Only(ctx context.Context) (*Repo, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{repo.Label}
	default:
		return nil, &NotSingularError{repo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RepoQuery) OnlyX(ctx context.Context) *Repo {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Repo ID in the query.
// Returns a *NotSingularError when more than one Repo ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *RepoQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{repo.Label}
	default:
		err = &NotSingularError{repo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RepoQuery) OnlyIDX(ctx context.Context) int {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Repos.
func (rq *RepoQuery) All(ctx context.Context) ([]*Repo, error) {
	ctx = setContextOp(ctx, rq.ctx, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Repo, *RepoQuery]()
	return withInterceptors[[]*Repo](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *RepoQuery) AllX(ctx context.Context) []*Repo {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Repo IDs.
func (rq *RepoQuery) IDs(ctx context.Context) (ids []int, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, "IDs")
	if err = rq.Select(repo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RepoQuery) IDsX(ctx context.Context) []int {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RepoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*RepoQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RepoQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RepoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, "Exist")
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RepoQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RepoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RepoQuery) Clone() *RepoQuery {
	if rq == nil {
		return nil
	}
	return &RepoQuery{
		config:     rq.config,
		ctx:        rq.ctx.Clone(),
		order:      append([]OrderFunc{}, rq.order...),
		inters:     append([]Interceptor{}, rq.inters...),
		predicates: append([]predicate.Repo{}, rq.predicates...),
		withTopics: rq.withTopics.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithTopics tells the query-builder to eager-load the nodes that are connected to
// the "topics" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RepoQuery) WithTopics(opts ...func(*TopicQuery)) *RepoQuery {
	query := (&TopicClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withTopics = query
	return rq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Repo.Query().
//		GroupBy(repo.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *RepoQuery) GroupBy(field string, fields ...string) *RepoGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RepoGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = repo.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Repo.Query().
//		Select(repo.FieldName).
//		Scan(ctx, &v)
func (rq *RepoQuery) Select(fields ...string) *RepoSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &RepoSelect{RepoQuery: rq}
	sbuild.label = repo.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RepoSelect configured with the given aggregations.
func (rq *RepoQuery) Aggregate(fns ...AggregateFunc) *RepoSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *RepoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !repo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *RepoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Repo, error) {
	var (
		nodes       = []*Repo{}
		_spec       = rq.querySpec()
		loadedTypes = [1]bool{
			rq.withTopics != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Repo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Repo{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(rq.modifiers) > 0 {
		_spec.Modifiers = rq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withTopics; query != nil {
		if err := rq.loadTopics(ctx, query, nodes, nil,
			func(n *Repo, e *Topic) { n.Edges.Topics = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *RepoQuery) loadTopics(ctx context.Context, query *TopicQuery, nodes []*Repo, init func(*Repo), assign func(*Repo, *Topic)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Repo)
	for i := range nodes {
		fk := nodes[i].TopicID
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
			return fmt.Errorf(`unexpected foreign-key "topic_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (rq *RepoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	if len(rq.modifiers) > 0 {
		_spec.Modifiers = rq.modifiers
	}
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RepoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(repo.Table, repo.Columns, sqlgraph.NewFieldSpec(repo.FieldID, field.TypeInt))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, repo.FieldID)
		for i := range fields {
			if fields[i] != repo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RepoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(repo.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = repo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range rq.modifiers {
		m(selector)
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (rq *RepoQuery) ForUpdate(opts ...sql.LockOption) *RepoQuery {
	if rq.driver.Dialect() == dialect.Postgres {
		rq.Unique(false)
	}
	rq.modifiers = append(rq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return rq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (rq *RepoQuery) ForShare(opts ...sql.LockOption) *RepoQuery {
	if rq.driver.Dialect() == dialect.Postgres {
		rq.Unique(false)
	}
	rq.modifiers = append(rq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return rq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (rq *RepoQuery) Modify(modifiers ...func(s *sql.Selector)) *RepoSelect {
	rq.modifiers = append(rq.modifiers, modifiers...)
	return rq.Select()
}

// RepoGroupBy is the group-by builder for Repo entities.
type RepoGroupBy struct {
	selector
	build *RepoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RepoGroupBy) Aggregate(fns ...AggregateFunc) *RepoGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *RepoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RepoQuery, *RepoGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *RepoGroupBy) sqlScan(ctx context.Context, root *RepoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RepoSelect is the builder for selecting fields of Repo entities.
type RepoSelect struct {
	*RepoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *RepoSelect) Aggregate(fns ...AggregateFunc) *RepoSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RepoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RepoQuery, *RepoSelect](ctx, rs.RepoQuery, rs, rs.inters, v)
}

func (rs *RepoSelect) sqlScan(ctx context.Context, root *RepoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (rs *RepoSelect) Modify(modifiers ...func(s *sql.Selector)) *RepoSelect {
	rs.modifiers = append(rs.modifiers, modifiers...)
	return rs
}
