// Code generated by ent, DO NOT EDIT.

package ent

import (
	"CSBackendTmp/ent/invite_code"
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

// InviteCodeUpdate is the builder for updating Invite_Code entities.
type InviteCodeUpdate struct {
	config
	hooks    []Hook
	mutation *InviteCodeMutation
}

// Where appends a list predicates to the InviteCodeUpdate builder.
func (icu *InviteCodeUpdate) Where(ps ...predicate.Invite_Code) *InviteCodeUpdate {
	icu.mutation.Where(ps...)
	return icu
}

// SetUpdateTime sets the "update_time" field.
func (icu *InviteCodeUpdate) SetUpdateTime(t time.Time) *InviteCodeUpdate {
	icu.mutation.SetUpdateTime(t)
	return icu
}

// SetType sets the "type" field.
func (icu *InviteCodeUpdate) SetType(ic invite_code.Type) *InviteCodeUpdate {
	icu.mutation.SetType(ic)
	return icu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (icu *InviteCodeUpdate) SetNillableType(ic *invite_code.Type) *InviteCodeUpdate {
	if ic != nil {
		icu.SetType(*ic)
	}
	return icu
}

// SetStatus sets the "status" field.
func (icu *InviteCodeUpdate) SetStatus(ic invite_code.Status) *InviteCodeUpdate {
	icu.mutation.SetStatus(ic)
	return icu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (icu *InviteCodeUpdate) SetNillableStatus(ic *invite_code.Status) *InviteCodeUpdate {
	if ic != nil {
		icu.SetStatus(*ic)
	}
	return icu
}

// SetCode sets the "code" field.
func (icu *InviteCodeUpdate) SetCode(s string) *InviteCodeUpdate {
	icu.mutation.SetCode(s)
	return icu
}

// SetConsumerID sets the "consumer_id" field.
func (icu *InviteCodeUpdate) SetConsumerID(u uint64) *InviteCodeUpdate {
	icu.mutation.ResetConsumerID()
	icu.mutation.SetConsumerID(u)
	return icu
}

// SetNillableConsumerID sets the "consumer_id" field if the given value is not nil.
func (icu *InviteCodeUpdate) SetNillableConsumerID(u *uint64) *InviteCodeUpdate {
	if u != nil {
		icu.SetConsumerID(*u)
	}
	return icu
}

// AddConsumerID adds u to the "consumer_id" field.
func (icu *InviteCodeUpdate) AddConsumerID(u int64) *InviteCodeUpdate {
	icu.mutation.AddConsumerID(u)
	return icu
}

// ClearConsumerID clears the value of the "consumer_id" field.
func (icu *InviteCodeUpdate) ClearConsumerID() *InviteCodeUpdate {
	icu.mutation.ClearConsumerID()
	return icu
}

// SetUserID sets the "user_id" field.
func (icu *InviteCodeUpdate) SetUserID(u uint64) *InviteCodeUpdate {
	icu.mutation.SetUserID(u)
	return icu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (icu *InviteCodeUpdate) SetOwnerID(id uint64) *InviteCodeUpdate {
	icu.mutation.SetOwnerID(id)
	return icu
}

// SetOwner sets the "owner" edge to the User entity.
func (icu *InviteCodeUpdate) SetOwner(u *User) *InviteCodeUpdate {
	return icu.SetOwnerID(u.ID)
}

// Mutation returns the InviteCodeMutation object of the builder.
func (icu *InviteCodeUpdate) Mutation() *InviteCodeMutation {
	return icu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (icu *InviteCodeUpdate) ClearOwner() *InviteCodeUpdate {
	icu.mutation.ClearOwner()
	return icu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (icu *InviteCodeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	icu.defaults()
	if len(icu.hooks) == 0 {
		if err = icu.check(); err != nil {
			return 0, err
		}
		affected, err = icu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InviteCodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = icu.check(); err != nil {
				return 0, err
			}
			icu.mutation = mutation
			affected, err = icu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(icu.hooks) - 1; i >= 0; i-- {
			if icu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = icu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, icu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (icu *InviteCodeUpdate) SaveX(ctx context.Context) int {
	affected, err := icu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (icu *InviteCodeUpdate) Exec(ctx context.Context) error {
	_, err := icu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icu *InviteCodeUpdate) ExecX(ctx context.Context) {
	if err := icu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (icu *InviteCodeUpdate) defaults() {
	if _, ok := icu.mutation.UpdateTime(); !ok {
		v := invite_code.UpdateDefaultUpdateTime()
		icu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (icu *InviteCodeUpdate) check() error {
	if v, ok := icu.mutation.GetType(); ok {
		if err := invite_code.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Invite_Code.type": %w`, err)}
		}
	}
	if v, ok := icu.mutation.Status(); ok {
		if err := invite_code.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Invite_Code.status": %w`, err)}
		}
	}
	if _, ok := icu.mutation.OwnerID(); icu.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Invite_Code.owner"`)
	}
	return nil
}

func (icu *InviteCodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   invite_code.Table,
			Columns: invite_code.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: invite_code.FieldID,
			},
		},
	}
	if ps := icu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := icu.mutation.UpdateTime(); ok {
		_spec.SetField(invite_code.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := icu.mutation.GetType(); ok {
		_spec.SetField(invite_code.FieldType, field.TypeEnum, value)
	}
	if value, ok := icu.mutation.Status(); ok {
		_spec.SetField(invite_code.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := icu.mutation.Code(); ok {
		_spec.SetField(invite_code.FieldCode, field.TypeString, value)
	}
	if value, ok := icu.mutation.ConsumerID(); ok {
		_spec.SetField(invite_code.FieldConsumerID, field.TypeUint64, value)
	}
	if value, ok := icu.mutation.AddedConsumerID(); ok {
		_spec.AddField(invite_code.FieldConsumerID, field.TypeUint64, value)
	}
	if icu.mutation.ConsumerIDCleared() {
		_spec.ClearField(invite_code.FieldConsumerID, field.TypeUint64)
	}
	if icu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invite_code.OwnerTable,
			Columns: []string{invite_code.OwnerColumn},
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
	if nodes := icu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invite_code.OwnerTable,
			Columns: []string{invite_code.OwnerColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, icu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invite_code.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// InviteCodeUpdateOne is the builder for updating a single Invite_Code entity.
type InviteCodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *InviteCodeMutation
}

// SetUpdateTime sets the "update_time" field.
func (icuo *InviteCodeUpdateOne) SetUpdateTime(t time.Time) *InviteCodeUpdateOne {
	icuo.mutation.SetUpdateTime(t)
	return icuo
}

// SetType sets the "type" field.
func (icuo *InviteCodeUpdateOne) SetType(ic invite_code.Type) *InviteCodeUpdateOne {
	icuo.mutation.SetType(ic)
	return icuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (icuo *InviteCodeUpdateOne) SetNillableType(ic *invite_code.Type) *InviteCodeUpdateOne {
	if ic != nil {
		icuo.SetType(*ic)
	}
	return icuo
}

// SetStatus sets the "status" field.
func (icuo *InviteCodeUpdateOne) SetStatus(ic invite_code.Status) *InviteCodeUpdateOne {
	icuo.mutation.SetStatus(ic)
	return icuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (icuo *InviteCodeUpdateOne) SetNillableStatus(ic *invite_code.Status) *InviteCodeUpdateOne {
	if ic != nil {
		icuo.SetStatus(*ic)
	}
	return icuo
}

// SetCode sets the "code" field.
func (icuo *InviteCodeUpdateOne) SetCode(s string) *InviteCodeUpdateOne {
	icuo.mutation.SetCode(s)
	return icuo
}

// SetConsumerID sets the "consumer_id" field.
func (icuo *InviteCodeUpdateOne) SetConsumerID(u uint64) *InviteCodeUpdateOne {
	icuo.mutation.ResetConsumerID()
	icuo.mutation.SetConsumerID(u)
	return icuo
}

// SetNillableConsumerID sets the "consumer_id" field if the given value is not nil.
func (icuo *InviteCodeUpdateOne) SetNillableConsumerID(u *uint64) *InviteCodeUpdateOne {
	if u != nil {
		icuo.SetConsumerID(*u)
	}
	return icuo
}

// AddConsumerID adds u to the "consumer_id" field.
func (icuo *InviteCodeUpdateOne) AddConsumerID(u int64) *InviteCodeUpdateOne {
	icuo.mutation.AddConsumerID(u)
	return icuo
}

// ClearConsumerID clears the value of the "consumer_id" field.
func (icuo *InviteCodeUpdateOne) ClearConsumerID() *InviteCodeUpdateOne {
	icuo.mutation.ClearConsumerID()
	return icuo
}

// SetUserID sets the "user_id" field.
func (icuo *InviteCodeUpdateOne) SetUserID(u uint64) *InviteCodeUpdateOne {
	icuo.mutation.SetUserID(u)
	return icuo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (icuo *InviteCodeUpdateOne) SetOwnerID(id uint64) *InviteCodeUpdateOne {
	icuo.mutation.SetOwnerID(id)
	return icuo
}

// SetOwner sets the "owner" edge to the User entity.
func (icuo *InviteCodeUpdateOne) SetOwner(u *User) *InviteCodeUpdateOne {
	return icuo.SetOwnerID(u.ID)
}

// Mutation returns the InviteCodeMutation object of the builder.
func (icuo *InviteCodeUpdateOne) Mutation() *InviteCodeMutation {
	return icuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (icuo *InviteCodeUpdateOne) ClearOwner() *InviteCodeUpdateOne {
	icuo.mutation.ClearOwner()
	return icuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (icuo *InviteCodeUpdateOne) Select(field string, fields ...string) *InviteCodeUpdateOne {
	icuo.fields = append([]string{field}, fields...)
	return icuo
}

// Save executes the query and returns the updated Invite_Code entity.
func (icuo *InviteCodeUpdateOne) Save(ctx context.Context) (*Invite_Code, error) {
	var (
		err  error
		node *Invite_Code
	)
	icuo.defaults()
	if len(icuo.hooks) == 0 {
		if err = icuo.check(); err != nil {
			return nil, err
		}
		node, err = icuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InviteCodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = icuo.check(); err != nil {
				return nil, err
			}
			icuo.mutation = mutation
			node, err = icuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(icuo.hooks) - 1; i >= 0; i-- {
			if icuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = icuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, icuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Invite_Code)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from InviteCodeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (icuo *InviteCodeUpdateOne) SaveX(ctx context.Context) *Invite_Code {
	node, err := icuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (icuo *InviteCodeUpdateOne) Exec(ctx context.Context) error {
	_, err := icuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icuo *InviteCodeUpdateOne) ExecX(ctx context.Context) {
	if err := icuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (icuo *InviteCodeUpdateOne) defaults() {
	if _, ok := icuo.mutation.UpdateTime(); !ok {
		v := invite_code.UpdateDefaultUpdateTime()
		icuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (icuo *InviteCodeUpdateOne) check() error {
	if v, ok := icuo.mutation.GetType(); ok {
		if err := invite_code.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Invite_Code.type": %w`, err)}
		}
	}
	if v, ok := icuo.mutation.Status(); ok {
		if err := invite_code.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Invite_Code.status": %w`, err)}
		}
	}
	if _, ok := icuo.mutation.OwnerID(); icuo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Invite_Code.owner"`)
	}
	return nil
}

func (icuo *InviteCodeUpdateOne) sqlSave(ctx context.Context) (_node *Invite_Code, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   invite_code.Table,
			Columns: invite_code.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: invite_code.FieldID,
			},
		},
	}
	id, ok := icuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Invite_Code.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := icuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, invite_code.FieldID)
		for _, f := range fields {
			if !invite_code.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != invite_code.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := icuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := icuo.mutation.UpdateTime(); ok {
		_spec.SetField(invite_code.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := icuo.mutation.GetType(); ok {
		_spec.SetField(invite_code.FieldType, field.TypeEnum, value)
	}
	if value, ok := icuo.mutation.Status(); ok {
		_spec.SetField(invite_code.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := icuo.mutation.Code(); ok {
		_spec.SetField(invite_code.FieldCode, field.TypeString, value)
	}
	if value, ok := icuo.mutation.ConsumerID(); ok {
		_spec.SetField(invite_code.FieldConsumerID, field.TypeUint64, value)
	}
	if value, ok := icuo.mutation.AddedConsumerID(); ok {
		_spec.AddField(invite_code.FieldConsumerID, field.TypeUint64, value)
	}
	if icuo.mutation.ConsumerIDCleared() {
		_spec.ClearField(invite_code.FieldConsumerID, field.TypeUint64)
	}
	if icuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invite_code.OwnerTable,
			Columns: []string{invite_code.OwnerColumn},
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
	if nodes := icuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invite_code.OwnerTable,
			Columns: []string{invite_code.OwnerColumn},
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
	_node = &Invite_Code{config: icuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, icuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invite_code.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}