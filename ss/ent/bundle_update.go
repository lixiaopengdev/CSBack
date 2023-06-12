// Code generated by ent, DO NOT EDIT.

package ent

import (
	"CSBackendTmp/ent/bundle"
	"CSBackendTmp/ent/mask"
	"CSBackendTmp/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BundleUpdate is the builder for updating Bundle entities.
type BundleUpdate struct {
	config
	hooks    []Hook
	mutation *BundleMutation
}

// Where appends a list predicates to the BundleUpdate builder.
func (bu *BundleUpdate) Where(ps ...predicate.Bundle) *BundleUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetUpdateTime sets the "update_time" field.
func (bu *BundleUpdate) SetUpdateTime(t time.Time) *BundleUpdate {
	bu.mutation.SetUpdateTime(t)
	return bu
}

// SetVerionID sets the "verionID" field.
func (bu *BundleUpdate) SetVerionID(u uint64) *BundleUpdate {
	bu.mutation.ResetVerionID()
	bu.mutation.SetVerionID(u)
	return bu
}

// SetNillableVerionID sets the "verionID" field if the given value is not nil.
func (bu *BundleUpdate) SetNillableVerionID(u *uint64) *BundleUpdate {
	if u != nil {
		bu.SetVerionID(*u)
	}
	return bu
}

// AddVerionID adds u to the "verionID" field.
func (bu *BundleUpdate) AddVerionID(u int64) *BundleUpdate {
	bu.mutation.AddVerionID(u)
	return bu
}

// ClearVerionID clears the value of the "verionID" field.
func (bu *BundleUpdate) ClearVerionID() *BundleUpdate {
	bu.mutation.ClearVerionID()
	return bu
}

// SetBundleURL sets the "bundle_url" field.
func (bu *BundleUpdate) SetBundleURL(s string) *BundleUpdate {
	bu.mutation.SetBundleURL(s)
	return bu
}

// SetNillableBundleURL sets the "bundle_url" field if the given value is not nil.
func (bu *BundleUpdate) SetNillableBundleURL(s *string) *BundleUpdate {
	if s != nil {
		bu.SetBundleURL(*s)
	}
	return bu
}

// ClearBundleURL clears the value of the "bundle_url" field.
func (bu *BundleUpdate) ClearBundleURL() *BundleUpdate {
	bu.mutation.ClearBundleURL()
	return bu
}

// SetStatus sets the "status" field.
func (bu *BundleUpdate) SetStatus(b bundle.Status) *BundleUpdate {
	bu.mutation.SetStatus(b)
	return bu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (bu *BundleUpdate) SetNillableStatus(b *bundle.Status) *BundleUpdate {
	if b != nil {
		bu.SetStatus(*b)
	}
	return bu
}

// SetPlatform sets the "platform" field.
func (bu *BundleUpdate) SetPlatform(b bundle.Platform) *BundleUpdate {
	bu.mutation.SetPlatform(b)
	return bu
}

// SetNillablePlatform sets the "platform" field if the given value is not nil.
func (bu *BundleUpdate) SetNillablePlatform(b *bundle.Platform) *BundleUpdate {
	if b != nil {
		bu.SetPlatform(*b)
	}
	return bu
}

// SetMaskID sets the "mask_id" field.
func (bu *BundleUpdate) SetMaskID(u uint64) *BundleUpdate {
	bu.mutation.SetMaskID(u)
	return bu
}

// SetNillableMaskID sets the "mask_id" field if the given value is not nil.
func (bu *BundleUpdate) SetNillableMaskID(u *uint64) *BundleUpdate {
	if u != nil {
		bu.SetMaskID(*u)
	}
	return bu
}

// ClearMaskID clears the value of the "mask_id" field.
func (bu *BundleUpdate) ClearMaskID() *BundleUpdate {
	bu.mutation.ClearMaskID()
	return bu
}

// SetOwnerID sets the "owner" edge to the Mask entity by ID.
func (bu *BundleUpdate) SetOwnerID(id uint64) *BundleUpdate {
	bu.mutation.SetOwnerID(id)
	return bu
}

// SetNillableOwnerID sets the "owner" edge to the Mask entity by ID if the given value is not nil.
func (bu *BundleUpdate) SetNillableOwnerID(id *uint64) *BundleUpdate {
	if id != nil {
		bu = bu.SetOwnerID(*id)
	}
	return bu
}

// SetOwner sets the "owner" edge to the Mask entity.
func (bu *BundleUpdate) SetOwner(m *Mask) *BundleUpdate {
	return bu.SetOwnerID(m.ID)
}

// Mutation returns the BundleMutation object of the builder.
func (bu *BundleUpdate) Mutation() *BundleMutation {
	return bu.mutation
}

// ClearOwner clears the "owner" edge to the Mask entity.
func (bu *BundleUpdate) ClearOwner() *BundleUpdate {
	bu.mutation.ClearOwner()
	return bu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BundleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	bu.defaults()
	if len(bu.hooks) == 0 {
		if err = bu.check(); err != nil {
			return 0, err
		}
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BundleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bu.check(); err != nil {
				return 0, err
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			if bu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BundleUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BundleUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BundleUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bu *BundleUpdate) defaults() {
	if _, ok := bu.mutation.UpdateTime(); !ok {
		v := bundle.UpdateDefaultUpdateTime()
		bu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BundleUpdate) check() error {
	if v, ok := bu.mutation.Status(); ok {
		if err := bundle.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Bundle.status": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Platform(); ok {
		if err := bundle.PlatformValidator(v); err != nil {
			return &ValidationError{Name: "platform", err: fmt.Errorf(`ent: validator failed for field "Bundle.platform": %w`, err)}
		}
	}
	return nil
}

func (bu *BundleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bundle.Table,
			Columns: bundle.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: bundle.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.UpdateTime(); ok {
		_spec.SetField(bundle.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := bu.mutation.VerionID(); ok {
		_spec.SetField(bundle.FieldVerionID, field.TypeUint64, value)
	}
	if value, ok := bu.mutation.AddedVerionID(); ok {
		_spec.AddField(bundle.FieldVerionID, field.TypeUint64, value)
	}
	if bu.mutation.VerionIDCleared() {
		_spec.ClearField(bundle.FieldVerionID, field.TypeUint64)
	}
	if value, ok := bu.mutation.BundleURL(); ok {
		_spec.SetField(bundle.FieldBundleURL, field.TypeString, value)
	}
	if bu.mutation.BundleURLCleared() {
		_spec.ClearField(bundle.FieldBundleURL, field.TypeString)
	}
	if value, ok := bu.mutation.Status(); ok {
		_spec.SetField(bundle.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := bu.mutation.Platform(); ok {
		_spec.SetField(bundle.FieldPlatform, field.TypeEnum, value)
	}
	if bu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bundle.OwnerTable,
			Columns: []string{bundle.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: mask.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bundle.OwnerTable,
			Columns: []string{bundle.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: mask.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bundle.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// BundleUpdateOne is the builder for updating a single Bundle entity.
type BundleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BundleMutation
}

// SetUpdateTime sets the "update_time" field.
func (buo *BundleUpdateOne) SetUpdateTime(t time.Time) *BundleUpdateOne {
	buo.mutation.SetUpdateTime(t)
	return buo
}

// SetVerionID sets the "verionID" field.
func (buo *BundleUpdateOne) SetVerionID(u uint64) *BundleUpdateOne {
	buo.mutation.ResetVerionID()
	buo.mutation.SetVerionID(u)
	return buo
}

// SetNillableVerionID sets the "verionID" field if the given value is not nil.
func (buo *BundleUpdateOne) SetNillableVerionID(u *uint64) *BundleUpdateOne {
	if u != nil {
		buo.SetVerionID(*u)
	}
	return buo
}

// AddVerionID adds u to the "verionID" field.
func (buo *BundleUpdateOne) AddVerionID(u int64) *BundleUpdateOne {
	buo.mutation.AddVerionID(u)
	return buo
}

// ClearVerionID clears the value of the "verionID" field.
func (buo *BundleUpdateOne) ClearVerionID() *BundleUpdateOne {
	buo.mutation.ClearVerionID()
	return buo
}

// SetBundleURL sets the "bundle_url" field.
func (buo *BundleUpdateOne) SetBundleURL(s string) *BundleUpdateOne {
	buo.mutation.SetBundleURL(s)
	return buo
}

// SetNillableBundleURL sets the "bundle_url" field if the given value is not nil.
func (buo *BundleUpdateOne) SetNillableBundleURL(s *string) *BundleUpdateOne {
	if s != nil {
		buo.SetBundleURL(*s)
	}
	return buo
}

// ClearBundleURL clears the value of the "bundle_url" field.
func (buo *BundleUpdateOne) ClearBundleURL() *BundleUpdateOne {
	buo.mutation.ClearBundleURL()
	return buo
}

// SetStatus sets the "status" field.
func (buo *BundleUpdateOne) SetStatus(b bundle.Status) *BundleUpdateOne {
	buo.mutation.SetStatus(b)
	return buo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (buo *BundleUpdateOne) SetNillableStatus(b *bundle.Status) *BundleUpdateOne {
	if b != nil {
		buo.SetStatus(*b)
	}
	return buo
}

// SetPlatform sets the "platform" field.
func (buo *BundleUpdateOne) SetPlatform(b bundle.Platform) *BundleUpdateOne {
	buo.mutation.SetPlatform(b)
	return buo
}

// SetNillablePlatform sets the "platform" field if the given value is not nil.
func (buo *BundleUpdateOne) SetNillablePlatform(b *bundle.Platform) *BundleUpdateOne {
	if b != nil {
		buo.SetPlatform(*b)
	}
	return buo
}

// SetMaskID sets the "mask_id" field.
func (buo *BundleUpdateOne) SetMaskID(u uint64) *BundleUpdateOne {
	buo.mutation.SetMaskID(u)
	return buo
}

// SetNillableMaskID sets the "mask_id" field if the given value is not nil.
func (buo *BundleUpdateOne) SetNillableMaskID(u *uint64) *BundleUpdateOne {
	if u != nil {
		buo.SetMaskID(*u)
	}
	return buo
}

// ClearMaskID clears the value of the "mask_id" field.
func (buo *BundleUpdateOne) ClearMaskID() *BundleUpdateOne {
	buo.mutation.ClearMaskID()
	return buo
}

// SetOwnerID sets the "owner" edge to the Mask entity by ID.
func (buo *BundleUpdateOne) SetOwnerID(id uint64) *BundleUpdateOne {
	buo.mutation.SetOwnerID(id)
	return buo
}

// SetNillableOwnerID sets the "owner" edge to the Mask entity by ID if the given value is not nil.
func (buo *BundleUpdateOne) SetNillableOwnerID(id *uint64) *BundleUpdateOne {
	if id != nil {
		buo = buo.SetOwnerID(*id)
	}
	return buo
}

// SetOwner sets the "owner" edge to the Mask entity.
func (buo *BundleUpdateOne) SetOwner(m *Mask) *BundleUpdateOne {
	return buo.SetOwnerID(m.ID)
}

// Mutation returns the BundleMutation object of the builder.
func (buo *BundleUpdateOne) Mutation() *BundleMutation {
	return buo.mutation
}

// ClearOwner clears the "owner" edge to the Mask entity.
func (buo *BundleUpdateOne) ClearOwner() *BundleUpdateOne {
	buo.mutation.ClearOwner()
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BundleUpdateOne) Select(field string, fields ...string) *BundleUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Bundle entity.
func (buo *BundleUpdateOne) Save(ctx context.Context) (*Bundle, error) {
	var (
		err  error
		node *Bundle
	)
	buo.defaults()
	if len(buo.hooks) == 0 {
		if err = buo.check(); err != nil {
			return nil, err
		}
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BundleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = buo.check(); err != nil {
				return nil, err
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			if buo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = buo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, buo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Bundle)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from BundleMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BundleUpdateOne) SaveX(ctx context.Context) *Bundle {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BundleUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BundleUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buo *BundleUpdateOne) defaults() {
	if _, ok := buo.mutation.UpdateTime(); !ok {
		v := bundle.UpdateDefaultUpdateTime()
		buo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BundleUpdateOne) check() error {
	if v, ok := buo.mutation.Status(); ok {
		if err := bundle.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Bundle.status": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Platform(); ok {
		if err := bundle.PlatformValidator(v); err != nil {
			return &ValidationError{Name: "platform", err: fmt.Errorf(`ent: validator failed for field "Bundle.platform": %w`, err)}
		}
	}
	return nil
}

func (buo *BundleUpdateOne) sqlSave(ctx context.Context) (_node *Bundle, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bundle.Table,
			Columns: bundle.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: bundle.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Bundle.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bundle.FieldID)
		for _, f := range fields {
			if !bundle.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != bundle.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.UpdateTime(); ok {
		_spec.SetField(bundle.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := buo.mutation.VerionID(); ok {
		_spec.SetField(bundle.FieldVerionID, field.TypeUint64, value)
	}
	if value, ok := buo.mutation.AddedVerionID(); ok {
		_spec.AddField(bundle.FieldVerionID, field.TypeUint64, value)
	}
	if buo.mutation.VerionIDCleared() {
		_spec.ClearField(bundle.FieldVerionID, field.TypeUint64)
	}
	if value, ok := buo.mutation.BundleURL(); ok {
		_spec.SetField(bundle.FieldBundleURL, field.TypeString, value)
	}
	if buo.mutation.BundleURLCleared() {
		_spec.ClearField(bundle.FieldBundleURL, field.TypeString)
	}
	if value, ok := buo.mutation.Status(); ok {
		_spec.SetField(bundle.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := buo.mutation.Platform(); ok {
		_spec.SetField(bundle.FieldPlatform, field.TypeEnum, value)
	}
	if buo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bundle.OwnerTable,
			Columns: []string{bundle.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: mask.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bundle.OwnerTable,
			Columns: []string{bundle.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: mask.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Bundle{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bundle.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}