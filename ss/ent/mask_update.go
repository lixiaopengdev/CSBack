// Code generated by ent, DO NOT EDIT.

package ent

import (
	"CSBackendTmp/ent/bundle"
	"CSBackendTmp/ent/mask"
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

// MaskUpdate is the builder for updating Mask entities.
type MaskUpdate struct {
	config
	hooks    []Hook
	mutation *MaskMutation
}

// Where appends a list predicates to the MaskUpdate builder.
func (mu *MaskUpdate) Where(ps ...predicate.Mask) *MaskUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetUpdateTime sets the "update_time" field.
func (mu *MaskUpdate) SetUpdateTime(t time.Time) *MaskUpdate {
	mu.mutation.SetUpdateTime(t)
	return mu
}

// SetName sets the "name" field.
func (mu *MaskUpdate) SetName(s string) *MaskUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetDesc sets the "desc" field.
func (mu *MaskUpdate) SetDesc(s string) *MaskUpdate {
	mu.mutation.SetDesc(s)
	return mu
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (mu *MaskUpdate) SetNillableDesc(s *string) *MaskUpdate {
	if s != nil {
		mu.SetDesc(*s)
	}
	return mu
}

// SetGUID sets the "GUID" field.
func (mu *MaskUpdate) SetGUID(s string) *MaskUpdate {
	mu.mutation.SetGUID(s)
	return mu
}

// SetThumbnailURL sets the "thumbnail_url" field.
func (mu *MaskUpdate) SetThumbnailURL(s string) *MaskUpdate {
	mu.mutation.SetThumbnailURL(s)
	return mu
}

// SetNillableThumbnailURL sets the "thumbnail_url" field if the given value is not nil.
func (mu *MaskUpdate) SetNillableThumbnailURL(s *string) *MaskUpdate {
	if s != nil {
		mu.SetThumbnailURL(*s)
	}
	return mu
}

// SetStatus sets the "status" field.
func (mu *MaskUpdate) SetStatus(m mask.Status) *MaskUpdate {
	mu.mutation.SetStatus(m)
	return mu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mu *MaskUpdate) SetNillableStatus(m *mask.Status) *MaskUpdate {
	if m != nil {
		mu.SetStatus(*m)
	}
	return mu
}

// SetType sets the "type" field.
func (mu *MaskUpdate) SetType(m mask.Type) *MaskUpdate {
	mu.mutation.SetType(m)
	return mu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (mu *MaskUpdate) SetNillableType(m *mask.Type) *MaskUpdate {
	if m != nil {
		mu.SetType(*m)
	}
	return mu
}

// SetUserID sets the "user_id" field.
func (mu *MaskUpdate) SetUserID(u uint64) *MaskUpdate {
	mu.mutation.SetUserID(u)
	return mu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (mu *MaskUpdate) SetNillableUserID(u *uint64) *MaskUpdate {
	if u != nil {
		mu.SetUserID(*u)
	}
	return mu
}

// ClearUserID clears the value of the "user_id" field.
func (mu *MaskUpdate) ClearUserID() *MaskUpdate {
	mu.mutation.ClearUserID()
	return mu
}

// AddBundleIDs adds the "bundle" edge to the Bundle entity by IDs.
func (mu *MaskUpdate) AddBundleIDs(ids ...uint64) *MaskUpdate {
	mu.mutation.AddBundleIDs(ids...)
	return mu
}

// AddBundle adds the "bundle" edges to the Bundle entity.
func (mu *MaskUpdate) AddBundle(b ...*Bundle) *MaskUpdate {
	ids := make([]uint64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return mu.AddBundleIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (mu *MaskUpdate) SetOwnerID(id uint64) *MaskUpdate {
	mu.mutation.SetOwnerID(id)
	return mu
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (mu *MaskUpdate) SetNillableOwnerID(id *uint64) *MaskUpdate {
	if id != nil {
		mu = mu.SetOwnerID(*id)
	}
	return mu
}

// SetOwner sets the "owner" edge to the User entity.
func (mu *MaskUpdate) SetOwner(u *User) *MaskUpdate {
	return mu.SetOwnerID(u.ID)
}

// Mutation returns the MaskMutation object of the builder.
func (mu *MaskUpdate) Mutation() *MaskMutation {
	return mu.mutation
}

// ClearBundle clears all "bundle" edges to the Bundle entity.
func (mu *MaskUpdate) ClearBundle() *MaskUpdate {
	mu.mutation.ClearBundle()
	return mu
}

// RemoveBundleIDs removes the "bundle" edge to Bundle entities by IDs.
func (mu *MaskUpdate) RemoveBundleIDs(ids ...uint64) *MaskUpdate {
	mu.mutation.RemoveBundleIDs(ids...)
	return mu
}

// RemoveBundle removes "bundle" edges to Bundle entities.
func (mu *MaskUpdate) RemoveBundle(b ...*Bundle) *MaskUpdate {
	ids := make([]uint64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return mu.RemoveBundleIDs(ids...)
}

// ClearOwner clears the "owner" edge to the User entity.
func (mu *MaskUpdate) ClearOwner() *MaskUpdate {
	mu.mutation.ClearOwner()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MaskUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	mu.defaults()
	if len(mu.hooks) == 0 {
		if err = mu.check(); err != nil {
			return 0, err
		}
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mu.check(); err != nil {
				return 0, err
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MaskUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MaskUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MaskUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MaskUpdate) defaults() {
	if _, ok := mu.mutation.UpdateTime(); !ok {
		v := mask.UpdateDefaultUpdateTime()
		mu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MaskUpdate) check() error {
	if v, ok := mu.mutation.Status(); ok {
		if err := mask.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Mask.status": %w`, err)}
		}
	}
	if v, ok := mu.mutation.GetType(); ok {
		if err := mask.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Mask.type": %w`, err)}
		}
	}
	return nil
}

func (mu *MaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mask.Table,
			Columns: mask.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: mask.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.UpdateTime(); ok {
		_spec.SetField(mask.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.SetField(mask.FieldName, field.TypeString, value)
	}
	if value, ok := mu.mutation.Desc(); ok {
		_spec.SetField(mask.FieldDesc, field.TypeString, value)
	}
	if value, ok := mu.mutation.GUID(); ok {
		_spec.SetField(mask.FieldGUID, field.TypeString, value)
	}
	if value, ok := mu.mutation.ThumbnailURL(); ok {
		_spec.SetField(mask.FieldThumbnailURL, field.TypeString, value)
	}
	if value, ok := mu.mutation.Status(); ok {
		_spec.SetField(mask.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := mu.mutation.GetType(); ok {
		_spec.SetField(mask.FieldType, field.TypeEnum, value)
	}
	if mu.mutation.BundleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mask.BundleTable,
			Columns: []string{mask.BundleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: bundle.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedBundleIDs(); len(nodes) > 0 && !mu.mutation.BundleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mask.BundleTable,
			Columns: []string{mask.BundleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: bundle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.BundleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mask.BundleTable,
			Columns: []string{mask.BundleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: bundle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mask.OwnerTable,
			Columns: []string{mask.OwnerColumn},
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
	if nodes := mu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mask.OwnerTable,
			Columns: []string{mask.OwnerColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MaskUpdateOne is the builder for updating a single Mask entity.
type MaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MaskMutation
}

// SetUpdateTime sets the "update_time" field.
func (muo *MaskUpdateOne) SetUpdateTime(t time.Time) *MaskUpdateOne {
	muo.mutation.SetUpdateTime(t)
	return muo
}

// SetName sets the "name" field.
func (muo *MaskUpdateOne) SetName(s string) *MaskUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetDesc sets the "desc" field.
func (muo *MaskUpdateOne) SetDesc(s string) *MaskUpdateOne {
	muo.mutation.SetDesc(s)
	return muo
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (muo *MaskUpdateOne) SetNillableDesc(s *string) *MaskUpdateOne {
	if s != nil {
		muo.SetDesc(*s)
	}
	return muo
}

// SetGUID sets the "GUID" field.
func (muo *MaskUpdateOne) SetGUID(s string) *MaskUpdateOne {
	muo.mutation.SetGUID(s)
	return muo
}

// SetThumbnailURL sets the "thumbnail_url" field.
func (muo *MaskUpdateOne) SetThumbnailURL(s string) *MaskUpdateOne {
	muo.mutation.SetThumbnailURL(s)
	return muo
}

// SetNillableThumbnailURL sets the "thumbnail_url" field if the given value is not nil.
func (muo *MaskUpdateOne) SetNillableThumbnailURL(s *string) *MaskUpdateOne {
	if s != nil {
		muo.SetThumbnailURL(*s)
	}
	return muo
}

// SetStatus sets the "status" field.
func (muo *MaskUpdateOne) SetStatus(m mask.Status) *MaskUpdateOne {
	muo.mutation.SetStatus(m)
	return muo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (muo *MaskUpdateOne) SetNillableStatus(m *mask.Status) *MaskUpdateOne {
	if m != nil {
		muo.SetStatus(*m)
	}
	return muo
}

// SetType sets the "type" field.
func (muo *MaskUpdateOne) SetType(m mask.Type) *MaskUpdateOne {
	muo.mutation.SetType(m)
	return muo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (muo *MaskUpdateOne) SetNillableType(m *mask.Type) *MaskUpdateOne {
	if m != nil {
		muo.SetType(*m)
	}
	return muo
}

// SetUserID sets the "user_id" field.
func (muo *MaskUpdateOne) SetUserID(u uint64) *MaskUpdateOne {
	muo.mutation.SetUserID(u)
	return muo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (muo *MaskUpdateOne) SetNillableUserID(u *uint64) *MaskUpdateOne {
	if u != nil {
		muo.SetUserID(*u)
	}
	return muo
}

// ClearUserID clears the value of the "user_id" field.
func (muo *MaskUpdateOne) ClearUserID() *MaskUpdateOne {
	muo.mutation.ClearUserID()
	return muo
}

// AddBundleIDs adds the "bundle" edge to the Bundle entity by IDs.
func (muo *MaskUpdateOne) AddBundleIDs(ids ...uint64) *MaskUpdateOne {
	muo.mutation.AddBundleIDs(ids...)
	return muo
}

// AddBundle adds the "bundle" edges to the Bundle entity.
func (muo *MaskUpdateOne) AddBundle(b ...*Bundle) *MaskUpdateOne {
	ids := make([]uint64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return muo.AddBundleIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (muo *MaskUpdateOne) SetOwnerID(id uint64) *MaskUpdateOne {
	muo.mutation.SetOwnerID(id)
	return muo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (muo *MaskUpdateOne) SetNillableOwnerID(id *uint64) *MaskUpdateOne {
	if id != nil {
		muo = muo.SetOwnerID(*id)
	}
	return muo
}

// SetOwner sets the "owner" edge to the User entity.
func (muo *MaskUpdateOne) SetOwner(u *User) *MaskUpdateOne {
	return muo.SetOwnerID(u.ID)
}

// Mutation returns the MaskMutation object of the builder.
func (muo *MaskUpdateOne) Mutation() *MaskMutation {
	return muo.mutation
}

// ClearBundle clears all "bundle" edges to the Bundle entity.
func (muo *MaskUpdateOne) ClearBundle() *MaskUpdateOne {
	muo.mutation.ClearBundle()
	return muo
}

// RemoveBundleIDs removes the "bundle" edge to Bundle entities by IDs.
func (muo *MaskUpdateOne) RemoveBundleIDs(ids ...uint64) *MaskUpdateOne {
	muo.mutation.RemoveBundleIDs(ids...)
	return muo
}

// RemoveBundle removes "bundle" edges to Bundle entities.
func (muo *MaskUpdateOne) RemoveBundle(b ...*Bundle) *MaskUpdateOne {
	ids := make([]uint64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return muo.RemoveBundleIDs(ids...)
}

// ClearOwner clears the "owner" edge to the User entity.
func (muo *MaskUpdateOne) ClearOwner() *MaskUpdateOne {
	muo.mutation.ClearOwner()
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MaskUpdateOne) Select(field string, fields ...string) *MaskUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Mask entity.
func (muo *MaskUpdateOne) Save(ctx context.Context) (*Mask, error) {
	var (
		err  error
		node *Mask
	)
	muo.defaults()
	if len(muo.hooks) == 0 {
		if err = muo.check(); err != nil {
			return nil, err
		}
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = muo.check(); err != nil {
				return nil, err
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, muo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Mask)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MaskMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MaskUpdateOne) SaveX(ctx context.Context) *Mask {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MaskUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MaskUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MaskUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdateTime(); !ok {
		v := mask.UpdateDefaultUpdateTime()
		muo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MaskUpdateOne) check() error {
	if v, ok := muo.mutation.Status(); ok {
		if err := mask.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Mask.status": %w`, err)}
		}
	}
	if v, ok := muo.mutation.GetType(); ok {
		if err := mask.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Mask.type": %w`, err)}
		}
	}
	return nil
}

func (muo *MaskUpdateOne) sqlSave(ctx context.Context) (_node *Mask, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mask.Table,
			Columns: mask.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: mask.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Mask.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mask.FieldID)
		for _, f := range fields {
			if !mask.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.UpdateTime(); ok {
		_spec.SetField(mask.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.SetField(mask.FieldName, field.TypeString, value)
	}
	if value, ok := muo.mutation.Desc(); ok {
		_spec.SetField(mask.FieldDesc, field.TypeString, value)
	}
	if value, ok := muo.mutation.GUID(); ok {
		_spec.SetField(mask.FieldGUID, field.TypeString, value)
	}
	if value, ok := muo.mutation.ThumbnailURL(); ok {
		_spec.SetField(mask.FieldThumbnailURL, field.TypeString, value)
	}
	if value, ok := muo.mutation.Status(); ok {
		_spec.SetField(mask.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := muo.mutation.GetType(); ok {
		_spec.SetField(mask.FieldType, field.TypeEnum, value)
	}
	if muo.mutation.BundleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mask.BundleTable,
			Columns: []string{mask.BundleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: bundle.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedBundleIDs(); len(nodes) > 0 && !muo.mutation.BundleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mask.BundleTable,
			Columns: []string{mask.BundleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: bundle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.BundleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mask.BundleTable,
			Columns: []string{mask.BundleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: bundle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mask.OwnerTable,
			Columns: []string{mask.OwnerColumn},
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
	if nodes := muo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mask.OwnerTable,
			Columns: []string{mask.OwnerColumn},
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
	_node = &Mask{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
