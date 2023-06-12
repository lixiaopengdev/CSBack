// Code generated by ent, DO NOT EDIT.

package ent

import (
	"CSBackendTmp/ent/collection"
	"CSBackendTmp/ent/predicate"
	"CSBackendTmp/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CollectionUpdate is the builder for updating Collection entities.
type CollectionUpdate struct {
	config
	hooks    []Hook
	mutation *CollectionMutation
}

// Where appends a list predicates to the CollectionUpdate builder.
func (cu *CollectionUpdate) Where(ps ...predicate.Collection) *CollectionUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdateTime sets the "update_time" field.
func (cu *CollectionUpdate) SetUpdateTime(t time.Time) *CollectionUpdate {
	cu.mutation.SetUpdateTime(t)
	return cu
}

// SetType sets the "type" field.
func (cu *CollectionUpdate) SetType(c collection.Type) *CollectionUpdate {
	cu.mutation.SetType(c)
	return cu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cu *CollectionUpdate) SetNillableType(c *collection.Type) *CollectionUpdate {
	if c != nil {
		cu.SetType(*c)
	}
	return cu
}

// SetItemID sets the "item_id" field.
func (cu *CollectionUpdate) SetItemID(u uint64) *CollectionUpdate {
	cu.mutation.ResetItemID()
	cu.mutation.SetItemID(u)
	return cu
}

// AddItemID adds u to the "item_id" field.
func (cu *CollectionUpdate) AddItemID(u int64) *CollectionUpdate {
	cu.mutation.AddItemID(u)
	return cu
}

// SetUserID sets the "user_id" field.
func (cu *CollectionUpdate) SetUserID(u uint64) *CollectionUpdate {
	cu.mutation.SetUserID(u)
	return cu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cu *CollectionUpdate) SetOwnerID(id uint64) *CollectionUpdate {
	cu.mutation.SetOwnerID(id)
	return cu
}

// SetOwner sets the "owner" edge to the User entity.
func (cu *CollectionUpdate) SetOwner(u *User) *CollectionUpdate {
	return cu.SetOwnerID(u.ID)
}

// Mutation returns the CollectionMutation object of the builder.
func (cu *CollectionUpdate) Mutation() *CollectionMutation {
	return cu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (cu *CollectionUpdate) ClearOwner() *CollectionUpdate {
	cu.mutation.ClearOwner()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CollectionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CollectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CollectionUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CollectionUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CollectionUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CollectionUpdate) defaults() {
	if _, ok := cu.mutation.UpdateTime(); !ok {
		v := collection.UpdateDefaultUpdateTime()
		cu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CollectionUpdate) check() error {
	if v, ok := cu.mutation.GetType(); ok {
		if err := collection.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Collection.type": %w`, err)}
		}
	}
	if _, ok := cu.mutation.OwnerID(); cu.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Collection.owner"`)
	}
	return nil
}

func (cu *CollectionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   collection.Table,
			Columns: collection.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: collection.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdateTime(); ok {
		_spec.SetField(collection.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := cu.mutation.GetType(); ok {
		_spec.SetField(collection.FieldType, field.TypeEnum, value)
	}
	if value, ok := cu.mutation.ItemID(); ok {
		_spec.SetField(collection.FieldItemID, field.TypeUint64, value)
	}
	if value, ok := cu.mutation.AddedItemID(); ok {
		_spec.AddField(collection.FieldItemID, field.TypeUint64, value)
	}
	if cu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   collection.OwnerTable,
			Columns: []string{collection.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   collection.OwnerTable,
			Columns: []string{collection.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{collection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CollectionUpdateOne is the builder for updating a single Collection entity.
type CollectionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CollectionMutation
}

// SetUpdateTime sets the "update_time" field.
func (cuo *CollectionUpdateOne) SetUpdateTime(t time.Time) *CollectionUpdateOne {
	cuo.mutation.SetUpdateTime(t)
	return cuo
}

// SetType sets the "type" field.
func (cuo *CollectionUpdateOne) SetType(c collection.Type) *CollectionUpdateOne {
	cuo.mutation.SetType(c)
	return cuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cuo *CollectionUpdateOne) SetNillableType(c *collection.Type) *CollectionUpdateOne {
	if c != nil {
		cuo.SetType(*c)
	}
	return cuo
}

// SetItemID sets the "item_id" field.
func (cuo *CollectionUpdateOne) SetItemID(u uint64) *CollectionUpdateOne {
	cuo.mutation.ResetItemID()
	cuo.mutation.SetItemID(u)
	return cuo
}

// AddItemID adds u to the "item_id" field.
func (cuo *CollectionUpdateOne) AddItemID(u int64) *CollectionUpdateOne {
	cuo.mutation.AddItemID(u)
	return cuo
}

// SetUserID sets the "user_id" field.
func (cuo *CollectionUpdateOne) SetUserID(u uint64) *CollectionUpdateOne {
	cuo.mutation.SetUserID(u)
	return cuo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cuo *CollectionUpdateOne) SetOwnerID(id uint64) *CollectionUpdateOne {
	cuo.mutation.SetOwnerID(id)
	return cuo
}

// SetOwner sets the "owner" edge to the User entity.
func (cuo *CollectionUpdateOne) SetOwner(u *User) *CollectionUpdateOne {
	return cuo.SetOwnerID(u.ID)
}

// Mutation returns the CollectionMutation object of the builder.
func (cuo *CollectionUpdateOne) Mutation() *CollectionMutation {
	return cuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (cuo *CollectionUpdateOne) ClearOwner() *CollectionUpdateOne {
	cuo.mutation.ClearOwner()
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CollectionUpdateOne) Select(field string, fields ...string) *CollectionUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Collection entity.
func (cuo *CollectionUpdateOne) Save(ctx context.Context) (*Collection, error) {
	var (
		err  error
		node *Collection
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CollectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Collection)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CollectionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CollectionUpdateOne) SaveX(ctx context.Context) *Collection {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CollectionUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CollectionUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CollectionUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdateTime(); !ok {
		v := collection.UpdateDefaultUpdateTime()
		cuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CollectionUpdateOne) check() error {
	if v, ok := cuo.mutation.GetType(); ok {
		if err := collection.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Collection.type": %w`, err)}
		}
	}
	if _, ok := cuo.mutation.OwnerID(); cuo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Collection.owner"`)
	}
	return nil
}

func (cuo *CollectionUpdateOne) sqlSave(ctx context.Context) (_node *Collection, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   collection.Table,
			Columns: collection.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: collection.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Collection.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, collection.FieldID)
		for _, f := range fields {
			if !collection.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != collection.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdateTime(); ok {
		_spec.SetField(collection.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.GetType(); ok {
		_spec.SetField(collection.FieldType, field.TypeEnum, value)
	}
	if value, ok := cuo.mutation.ItemID(); ok {
		_spec.SetField(collection.FieldItemID, field.TypeUint64, value)
	}
	if value, ok := cuo.mutation.AddedItemID(); ok {
		_spec.AddField(collection.FieldItemID, field.TypeUint64, value)
	}
	if cuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   collection.OwnerTable,
			Columns: []string{collection.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   collection.OwnerTable,
			Columns: []string{collection.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Collection{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{collection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
