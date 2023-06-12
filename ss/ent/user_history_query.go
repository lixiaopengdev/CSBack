// Code generated by ent, DO NOT EDIT.

package ent

import (
	"CSBackendTmp/ent/predicate"
	"CSBackendTmp/ent/user"
	"CSBackendTmp/ent/user_history"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserHistoryQuery is the builder for querying User_history entities.
type UserHistoryQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.User_history
	withOwner  *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserHistoryQuery builder.
func (uhq *UserHistoryQuery) Where(ps ...predicate.User_history) *UserHistoryQuery {
	uhq.predicates = append(uhq.predicates, ps...)
	return uhq
}

// Limit adds a limit step to the query.
func (uhq *UserHistoryQuery) Limit(limit int) *UserHistoryQuery {
	uhq.limit = &limit
	return uhq
}

// Offset adds an offset step to the query.
func (uhq *UserHistoryQuery) Offset(offset int) *UserHistoryQuery {
	uhq.offset = &offset
	return uhq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uhq *UserHistoryQuery) Unique(unique bool) *UserHistoryQuery {
	uhq.unique = &unique
	return uhq
}

// Order adds an order step to the query.
func (uhq *UserHistoryQuery) Order(o ...OrderFunc) *UserHistoryQuery {
	uhq.order = append(uhq.order, o...)
	return uhq
}

// QueryOwner chains the current query on the "owner" edge.
func (uhq *UserHistoryQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: uhq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uhq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uhq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user_history.Table, user_history.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user_history.OwnerTable, user_history.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(uhq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first User_history entity from the query.
// Returns a *NotFoundError when no User_history was found.
func (uhq *UserHistoryQuery) First(ctx context.Context) (*User_history, error) {
	nodes, err := uhq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{user_history.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uhq *UserHistoryQuery) FirstX(ctx context.Context) *User_history {
	node, err := uhq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first User_history ID from the query.
// Returns a *NotFoundError when no User_history ID was found.
func (uhq *UserHistoryQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = uhq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{user_history.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uhq *UserHistoryQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := uhq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single User_history entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one User_history entity is found.
// Returns a *NotFoundError when no User_history entities are found.
func (uhq *UserHistoryQuery) Only(ctx context.Context) (*User_history, error) {
	nodes, err := uhq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{user_history.Label}
	default:
		return nil, &NotSingularError{user_history.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uhq *UserHistoryQuery) OnlyX(ctx context.Context) *User_history {
	node, err := uhq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only User_history ID in the query.
// Returns a *NotSingularError when more than one User_history ID is found.
// Returns a *NotFoundError when no entities are found.
func (uhq *UserHistoryQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = uhq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{user_history.Label}
	default:
		err = &NotSingularError{user_history.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uhq *UserHistoryQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := uhq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of User_histories.
func (uhq *UserHistoryQuery) All(ctx context.Context) ([]*User_history, error) {
	if err := uhq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return uhq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uhq *UserHistoryQuery) AllX(ctx context.Context) []*User_history {
	nodes, err := uhq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of User_history IDs.
func (uhq *UserHistoryQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := uhq.Select(user_history.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uhq *UserHistoryQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := uhq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uhq *UserHistoryQuery) Count(ctx context.Context) (int, error) {
	if err := uhq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return uhq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uhq *UserHistoryQuery) CountX(ctx context.Context) int {
	count, err := uhq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uhq *UserHistoryQuery) Exist(ctx context.Context) (bool, error) {
	if err := uhq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return uhq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uhq *UserHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := uhq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uhq *UserHistoryQuery) Clone() *UserHistoryQuery {
	if uhq == nil {
		return nil
	}
	return &UserHistoryQuery{
		config:     uhq.config,
		limit:      uhq.limit,
		offset:     uhq.offset,
		order:      append([]OrderFunc{}, uhq.order...),
		predicates: append([]predicate.User_history{}, uhq.predicates...),
		withOwner:  uhq.withOwner.Clone(),
		// clone intermediate query.
		sql:    uhq.sql.Clone(),
		path:   uhq.path,
		unique: uhq.unique,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (uhq *UserHistoryQuery) WithOwner(opts ...func(*UserQuery)) *UserHistoryQuery {
	query := &UserQuery{config: uhq.config}
	for _, opt := range opts {
		opt(query)
	}
	uhq.withOwner = query
	return uhq
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
//	client.UserHistory.Query().
//		GroupBy(user_history.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uhq *UserHistoryQuery) GroupBy(field string, fields ...string) *UserHistoryGroupBy {
	grbuild := &UserHistoryGroupBy{config: uhq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := uhq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return uhq.sqlQuery(ctx), nil
	}
	grbuild.label = user_history.Label
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
//	client.UserHistory.Query().
//		Select(user_history.FieldCreateTime).
//		Scan(ctx, &v)
func (uhq *UserHistoryQuery) Select(fields ...string) *UserHistorySelect {
	uhq.fields = append(uhq.fields, fields...)
	selbuild := &UserHistorySelect{UserHistoryQuery: uhq}
	selbuild.label = user_history.Label
	selbuild.flds, selbuild.scan = &uhq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a UserHistorySelect configured with the given aggregations.
func (uhq *UserHistoryQuery) Aggregate(fns ...AggregateFunc) *UserHistorySelect {
	return uhq.Select().Aggregate(fns...)
}

func (uhq *UserHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range uhq.fields {
		if !user_history.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uhq.path != nil {
		prev, err := uhq.path(ctx)
		if err != nil {
			return err
		}
		uhq.sql = prev
	}
	return nil
}

func (uhq *UserHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*User_history, error) {
	var (
		nodes       = []*User_history{}
		_spec       = uhq.querySpec()
		loadedTypes = [1]bool{
			uhq.withOwner != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*User_history).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &User_history{config: uhq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uhq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uhq.withOwner; query != nil {
		if err := uhq.loadOwner(ctx, query, nodes, nil,
			func(n *User_history, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uhq *UserHistoryQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*User_history, init func(*User_history), assign func(*User_history, *User)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*User_history)
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

func (uhq *UserHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uhq.querySpec()
	_spec.Node.Columns = uhq.fields
	if len(uhq.fields) > 0 {
		_spec.Unique = uhq.unique != nil && *uhq.unique
	}
	return sqlgraph.CountNodes(ctx, uhq.driver, _spec)
}

func (uhq *UserHistoryQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := uhq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (uhq *UserHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user_history.Table,
			Columns: user_history.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: user_history.FieldID,
			},
		},
		From:   uhq.sql,
		Unique: true,
	}
	if unique := uhq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := uhq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user_history.FieldID)
		for i := range fields {
			if fields[i] != user_history.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uhq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uhq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uhq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uhq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uhq *UserHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uhq.driver.Dialect())
	t1 := builder.Table(user_history.Table)
	columns := uhq.fields
	if len(columns) == 0 {
		columns = user_history.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uhq.sql != nil {
		selector = uhq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uhq.unique != nil && *uhq.unique {
		selector.Distinct()
	}
	for _, p := range uhq.predicates {
		p(selector)
	}
	for _, p := range uhq.order {
		p(selector)
	}
	if offset := uhq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uhq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserHistoryGroupBy is the group-by builder for User_history entities.
type UserHistoryGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uhgb *UserHistoryGroupBy) Aggregate(fns ...AggregateFunc) *UserHistoryGroupBy {
	uhgb.fns = append(uhgb.fns, fns...)
	return uhgb
}

// Scan applies the group-by query and scans the result into the given value.
func (uhgb *UserHistoryGroupBy) Scan(ctx context.Context, v any) error {
	query, err := uhgb.path(ctx)
	if err != nil {
		return err
	}
	uhgb.sql = query
	return uhgb.sqlScan(ctx, v)
}

func (uhgb *UserHistoryGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range uhgb.fields {
		if !user_history.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := uhgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uhgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uhgb *UserHistoryGroupBy) sqlQuery() *sql.Selector {
	selector := uhgb.sql.Select()
	aggregation := make([]string, 0, len(uhgb.fns))
	for _, fn := range uhgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(uhgb.fields)+len(uhgb.fns))
		for _, f := range uhgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(uhgb.fields...)...)
}

// UserHistorySelect is the builder for selecting fields of UserHistory entities.
type UserHistorySelect struct {
	*UserHistoryQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (uhs *UserHistorySelect) Aggregate(fns ...AggregateFunc) *UserHistorySelect {
	uhs.fns = append(uhs.fns, fns...)
	return uhs
}

// Scan applies the selector query and scans the result into the given value.
func (uhs *UserHistorySelect) Scan(ctx context.Context, v any) error {
	if err := uhs.prepareQuery(ctx); err != nil {
		return err
	}
	uhs.sql = uhs.UserHistoryQuery.sqlQuery(ctx)
	return uhs.sqlScan(ctx, v)
}

func (uhs *UserHistorySelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(uhs.fns))
	for _, fn := range uhs.fns {
		aggregation = append(aggregation, fn(uhs.sql))
	}
	switch n := len(*uhs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		uhs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		uhs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := uhs.sql.Query()
	if err := uhs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
