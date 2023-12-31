// Code generated by ent, DO NOT EDIT.

package ent

import (
	"CSBackendTmp/ent/predicate"
	"CSBackendTmp/ent/user"
	"CSBackendTmp/ent/user_auth"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserAuthQuery is the builder for querying User_auth entities.
type UserAuthQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.User_auth
	withOwner  *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserAuthQuery builder.
func (uaq *UserAuthQuery) Where(ps ...predicate.User_auth) *UserAuthQuery {
	uaq.predicates = append(uaq.predicates, ps...)
	return uaq
}

// Limit adds a limit step to the query.
func (uaq *UserAuthQuery) Limit(limit int) *UserAuthQuery {
	uaq.limit = &limit
	return uaq
}

// Offset adds an offset step to the query.
func (uaq *UserAuthQuery) Offset(offset int) *UserAuthQuery {
	uaq.offset = &offset
	return uaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uaq *UserAuthQuery) Unique(unique bool) *UserAuthQuery {
	uaq.unique = &unique
	return uaq
}

// Order adds an order step to the query.
func (uaq *UserAuthQuery) Order(o ...OrderFunc) *UserAuthQuery {
	uaq.order = append(uaq.order, o...)
	return uaq
}

// QueryOwner chains the current query on the "owner" edge.
func (uaq *UserAuthQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: uaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user_auth.Table, user_auth.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user_auth.OwnerTable, user_auth.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(uaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first User_auth entity from the query.
// Returns a *NotFoundError when no User_auth was found.
func (uaq *UserAuthQuery) First(ctx context.Context) (*User_auth, error) {
	nodes, err := uaq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{user_auth.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uaq *UserAuthQuery) FirstX(ctx context.Context) *User_auth {
	node, err := uaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first User_auth ID from the query.
// Returns a *NotFoundError when no User_auth ID was found.
func (uaq *UserAuthQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uaq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{user_auth.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uaq *UserAuthQuery) FirstIDX(ctx context.Context) int {
	id, err := uaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single User_auth entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one User_auth entity is found.
// Returns a *NotFoundError when no User_auth entities are found.
func (uaq *UserAuthQuery) Only(ctx context.Context) (*User_auth, error) {
	nodes, err := uaq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{user_auth.Label}
	default:
		return nil, &NotSingularError{user_auth.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uaq *UserAuthQuery) OnlyX(ctx context.Context) *User_auth {
	node, err := uaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only User_auth ID in the query.
// Returns a *NotSingularError when more than one User_auth ID is found.
// Returns a *NotFoundError when no entities are found.
func (uaq *UserAuthQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uaq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{user_auth.Label}
	default:
		err = &NotSingularError{user_auth.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uaq *UserAuthQuery) OnlyIDX(ctx context.Context) int {
	id, err := uaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of User_auths.
func (uaq *UserAuthQuery) All(ctx context.Context) ([]*User_auth, error) {
	if err := uaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return uaq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uaq *UserAuthQuery) AllX(ctx context.Context) []*User_auth {
	nodes, err := uaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of User_auth IDs.
func (uaq *UserAuthQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := uaq.Select(user_auth.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uaq *UserAuthQuery) IDsX(ctx context.Context) []int {
	ids, err := uaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uaq *UserAuthQuery) Count(ctx context.Context) (int, error) {
	if err := uaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return uaq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uaq *UserAuthQuery) CountX(ctx context.Context) int {
	count, err := uaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uaq *UserAuthQuery) Exist(ctx context.Context) (bool, error) {
	if err := uaq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return uaq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uaq *UserAuthQuery) ExistX(ctx context.Context) bool {
	exist, err := uaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserAuthQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uaq *UserAuthQuery) Clone() *UserAuthQuery {
	if uaq == nil {
		return nil
	}
	return &UserAuthQuery{
		config:     uaq.config,
		limit:      uaq.limit,
		offset:     uaq.offset,
		order:      append([]OrderFunc{}, uaq.order...),
		predicates: append([]predicate.User_auth{}, uaq.predicates...),
		withOwner:  uaq.withOwner.Clone(),
		// clone intermediate query.
		sql:    uaq.sql.Clone(),
		path:   uaq.path,
		unique: uaq.unique,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (uaq *UserAuthQuery) WithOwner(opts ...func(*UserQuery)) *UserAuthQuery {
	query := &UserQuery{config: uaq.config}
	for _, opt := range opts {
		opt(query)
	}
	uaq.withOwner = query
	return uaq
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
//	client.UserAuth.Query().
//		GroupBy(user_auth.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uaq *UserAuthQuery) GroupBy(field string, fields ...string) *UserAuthGroupBy {
	grbuild := &UserAuthGroupBy{config: uaq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := uaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return uaq.sqlQuery(ctx), nil
	}
	grbuild.label = user_auth.Label
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
//	client.UserAuth.Query().
//		Select(user_auth.FieldCreateTime).
//		Scan(ctx, &v)
func (uaq *UserAuthQuery) Select(fields ...string) *UserAuthSelect {
	uaq.fields = append(uaq.fields, fields...)
	selbuild := &UserAuthSelect{UserAuthQuery: uaq}
	selbuild.label = user_auth.Label
	selbuild.flds, selbuild.scan = &uaq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a UserAuthSelect configured with the given aggregations.
func (uaq *UserAuthQuery) Aggregate(fns ...AggregateFunc) *UserAuthSelect {
	return uaq.Select().Aggregate(fns...)
}

func (uaq *UserAuthQuery) prepareQuery(ctx context.Context) error {
	for _, f := range uaq.fields {
		if !user_auth.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uaq.path != nil {
		prev, err := uaq.path(ctx)
		if err != nil {
			return err
		}
		uaq.sql = prev
	}
	return nil
}

func (uaq *UserAuthQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*User_auth, error) {
	var (
		nodes       = []*User_auth{}
		_spec       = uaq.querySpec()
		loadedTypes = [1]bool{
			uaq.withOwner != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*User_auth).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &User_auth{config: uaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uaq.withOwner; query != nil {
		if err := uaq.loadOwner(ctx, query, nodes, nil,
			func(n *User_auth, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uaq *UserAuthQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*User_auth, init func(*User_auth), assign func(*User_auth, *User)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*User_auth)
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

func (uaq *UserAuthQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uaq.querySpec()
	_spec.Node.Columns = uaq.fields
	if len(uaq.fields) > 0 {
		_spec.Unique = uaq.unique != nil && *uaq.unique
	}
	return sqlgraph.CountNodes(ctx, uaq.driver, _spec)
}

func (uaq *UserAuthQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := uaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (uaq *UserAuthQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user_auth.Table,
			Columns: user_auth.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user_auth.FieldID,
			},
		},
		From:   uaq.sql,
		Unique: true,
	}
	if unique := uaq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := uaq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user_auth.FieldID)
		for i := range fields {
			if fields[i] != user_auth.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uaq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uaq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uaq *UserAuthQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uaq.driver.Dialect())
	t1 := builder.Table(user_auth.Table)
	columns := uaq.fields
	if len(columns) == 0 {
		columns = user_auth.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uaq.sql != nil {
		selector = uaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uaq.unique != nil && *uaq.unique {
		selector.Distinct()
	}
	for _, p := range uaq.predicates {
		p(selector)
	}
	for _, p := range uaq.order {
		p(selector)
	}
	if offset := uaq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uaq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserAuthGroupBy is the group-by builder for User_auth entities.
type UserAuthGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uagb *UserAuthGroupBy) Aggregate(fns ...AggregateFunc) *UserAuthGroupBy {
	uagb.fns = append(uagb.fns, fns...)
	return uagb
}

// Scan applies the group-by query and scans the result into the given value.
func (uagb *UserAuthGroupBy) Scan(ctx context.Context, v any) error {
	query, err := uagb.path(ctx)
	if err != nil {
		return err
	}
	uagb.sql = query
	return uagb.sqlScan(ctx, v)
}

func (uagb *UserAuthGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range uagb.fields {
		if !user_auth.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := uagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uagb *UserAuthGroupBy) sqlQuery() *sql.Selector {
	selector := uagb.sql.Select()
	aggregation := make([]string, 0, len(uagb.fns))
	for _, fn := range uagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(uagb.fields)+len(uagb.fns))
		for _, f := range uagb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(uagb.fields...)...)
}

// UserAuthSelect is the builder for selecting fields of UserAuth entities.
type UserAuthSelect struct {
	*UserAuthQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (uas *UserAuthSelect) Aggregate(fns ...AggregateFunc) *UserAuthSelect {
	uas.fns = append(uas.fns, fns...)
	return uas
}

// Scan applies the selector query and scans the result into the given value.
func (uas *UserAuthSelect) Scan(ctx context.Context, v any) error {
	if err := uas.prepareQuery(ctx); err != nil {
		return err
	}
	uas.sql = uas.UserAuthQuery.sqlQuery(ctx)
	return uas.sqlScan(ctx, v)
}

func (uas *UserAuthSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(uas.fns))
	for _, fn := range uas.fns {
		aggregation = append(aggregation, fn(uas.sql))
	}
	switch n := len(*uas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		uas.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		uas.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := uas.sql.Query()
	if err := uas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
