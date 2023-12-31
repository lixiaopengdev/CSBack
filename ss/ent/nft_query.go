// Code generated by ent, DO NOT EDIT.

package ent

import (
	"CSBackendTmp/ent/nft"
	"CSBackendTmp/ent/predicate"
	"CSBackendTmp/ent/user"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NFTQuery is the builder for querying NFT entities.
type NFTQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.NFT
	withOwner  *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NFTQuery builder.
func (nq *NFTQuery) Where(ps ...predicate.NFT) *NFTQuery {
	nq.predicates = append(nq.predicates, ps...)
	return nq
}

// Limit adds a limit step to the query.
func (nq *NFTQuery) Limit(limit int) *NFTQuery {
	nq.limit = &limit
	return nq
}

// Offset adds an offset step to the query.
func (nq *NFTQuery) Offset(offset int) *NFTQuery {
	nq.offset = &offset
	return nq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nq *NFTQuery) Unique(unique bool) *NFTQuery {
	nq.unique = &unique
	return nq
}

// Order adds an order step to the query.
func (nq *NFTQuery) Order(o ...OrderFunc) *NFTQuery {
	nq.order = append(nq.order, o...)
	return nq
}

// QueryOwner chains the current query on the "owner" edge.
func (nq *NFTQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: nq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(nft.Table, nft.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, nft.OwnerTable, nft.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first NFT entity from the query.
// Returns a *NotFoundError when no NFT was found.
func (nq *NFTQuery) First(ctx context.Context) (*NFT, error) {
	nodes, err := nq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{nft.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nq *NFTQuery) FirstX(ctx context.Context) *NFT {
	node, err := nq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NFT ID from the query.
// Returns a *NotFoundError when no NFT ID was found.
func (nq *NFTQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = nq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{nft.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nq *NFTQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := nq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NFT entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one NFT entity is found.
// Returns a *NotFoundError when no NFT entities are found.
func (nq *NFTQuery) Only(ctx context.Context) (*NFT, error) {
	nodes, err := nq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{nft.Label}
	default:
		return nil, &NotSingularError{nft.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nq *NFTQuery) OnlyX(ctx context.Context) *NFT {
	node, err := nq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NFT ID in the query.
// Returns a *NotSingularError when more than one NFT ID is found.
// Returns a *NotFoundError when no entities are found.
func (nq *NFTQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = nq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{nft.Label}
	default:
		err = &NotSingularError{nft.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nq *NFTQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := nq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NFTs.
func (nq *NFTQuery) All(ctx context.Context) ([]*NFT, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return nq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (nq *NFTQuery) AllX(ctx context.Context) []*NFT {
	nodes, err := nq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NFT IDs.
func (nq *NFTQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := nq.Select(nft.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nq *NFTQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := nq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nq *NFTQuery) Count(ctx context.Context) (int, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return nq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (nq *NFTQuery) CountX(ctx context.Context) int {
	count, err := nq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nq *NFTQuery) Exist(ctx context.Context) (bool, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return nq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (nq *NFTQuery) ExistX(ctx context.Context) bool {
	exist, err := nq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NFTQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nq *NFTQuery) Clone() *NFTQuery {
	if nq == nil {
		return nil
	}
	return &NFTQuery{
		config:     nq.config,
		limit:      nq.limit,
		offset:     nq.offset,
		order:      append([]OrderFunc{}, nq.order...),
		predicates: append([]predicate.NFT{}, nq.predicates...),
		withOwner:  nq.withOwner.Clone(),
		// clone intermediate query.
		sql:    nq.sql.Clone(),
		path:   nq.path,
		unique: nq.unique,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NFTQuery) WithOwner(opts ...func(*UserQuery)) *NFTQuery {
	query := &UserQuery{config: nq.config}
	for _, opt := range opts {
		opt(query)
	}
	nq.withOwner = query
	return nq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.NFT.Query().
//		GroupBy(nft.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (nq *NFTQuery) GroupBy(field string, fields ...string) *NFTGroupBy {
	grbuild := &NFTGroupBy{config: nq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return nq.sqlQuery(ctx), nil
	}
	grbuild.label = nft.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.NFT.Query().
//		Select(nft.FieldCreateTime).
//		Scan(ctx, &v)
func (nq *NFTQuery) Select(fields ...string) *NFTSelect {
	nq.fields = append(nq.fields, fields...)
	selbuild := &NFTSelect{NFTQuery: nq}
	selbuild.label = nft.Label
	selbuild.flds, selbuild.scan = &nq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a NFTSelect configured with the given aggregations.
func (nq *NFTQuery) Aggregate(fns ...AggregateFunc) *NFTSelect {
	return nq.Select().Aggregate(fns...)
}

func (nq *NFTQuery) prepareQuery(ctx context.Context) error {
	for _, f := range nq.fields {
		if !nft.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nq.path != nil {
		prev, err := nq.path(ctx)
		if err != nil {
			return err
		}
		nq.sql = prev
	}
	return nil
}

func (nq *NFTQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*NFT, error) {
	var (
		nodes       = []*NFT{}
		withFKs     = nq.withFKs
		_spec       = nq.querySpec()
		loadedTypes = [1]bool{
			nq.withOwner != nil,
		}
	)
	if nq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, nft.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*NFT).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &NFT{config: nq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := nq.withOwner; query != nil {
		if err := nq.loadOwner(ctx, query, nodes, nil,
			func(n *NFT, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (nq *NFTQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*NFT, init func(*NFT), assign func(*NFT, *User)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*NFT)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
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

func (nq *NFTQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nq.querySpec()
	_spec.Node.Columns = nq.fields
	if len(nq.fields) > 0 {
		_spec.Unique = nq.unique != nil && *nq.unique
	}
	return sqlgraph.CountNodes(ctx, nq.driver, _spec)
}

func (nq *NFTQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := nq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (nq *NFTQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nft.Table,
			Columns: nft.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: nft.FieldID,
			},
		},
		From:   nq.sql,
		Unique: true,
	}
	if unique := nq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := nq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nft.FieldID)
		for i := range fields {
			if fields[i] != nft.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nq *NFTQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nq.driver.Dialect())
	t1 := builder.Table(nft.Table)
	columns := nq.fields
	if len(columns) == 0 {
		columns = nft.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nq.sql != nil {
		selector = nq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nq.unique != nil && *nq.unique {
		selector.Distinct()
	}
	for _, p := range nq.predicates {
		p(selector)
	}
	for _, p := range nq.order {
		p(selector)
	}
	if offset := nq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NFTGroupBy is the group-by builder for NFT entities.
type NFTGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ngb *NFTGroupBy) Aggregate(fns ...AggregateFunc) *NFTGroupBy {
	ngb.fns = append(ngb.fns, fns...)
	return ngb
}

// Scan applies the group-by query and scans the result into the given value.
func (ngb *NFTGroupBy) Scan(ctx context.Context, v any) error {
	query, err := ngb.path(ctx)
	if err != nil {
		return err
	}
	ngb.sql = query
	return ngb.sqlScan(ctx, v)
}

func (ngb *NFTGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range ngb.fields {
		if !nft.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ngb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ngb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ngb *NFTGroupBy) sqlQuery() *sql.Selector {
	selector := ngb.sql.Select()
	aggregation := make([]string, 0, len(ngb.fns))
	for _, fn := range ngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ngb.fields)+len(ngb.fns))
		for _, f := range ngb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ngb.fields...)...)
}

// NFTSelect is the builder for selecting fields of NFT entities.
type NFTSelect struct {
	*NFTQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ns *NFTSelect) Aggregate(fns ...AggregateFunc) *NFTSelect {
	ns.fns = append(ns.fns, fns...)
	return ns
}

// Scan applies the selector query and scans the result into the given value.
func (ns *NFTSelect) Scan(ctx context.Context, v any) error {
	if err := ns.prepareQuery(ctx); err != nil {
		return err
	}
	ns.sql = ns.NFTQuery.sqlQuery(ctx)
	return ns.sqlScan(ctx, v)
}

func (ns *NFTSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ns.fns))
	for _, fn := range ns.fns {
		aggregation = append(aggregation, fn(ns.sql))
	}
	switch n := len(*ns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ns.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ns.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ns.sql.Query()
	if err := ns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
