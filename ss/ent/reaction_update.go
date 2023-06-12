// Code generated by ent, DO NOT EDIT.

package ent

import (
	"CSBackendTmp/ent/predicate"
	"CSBackendTmp/ent/reaction"
	"CSBackendTmp/ent/timedew"
	"CSBackendTmp/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReactionUpdate is the builder for updating Reaction entities.
type ReactionUpdate struct {
	config
	hooks    []Hook
	mutation *ReactionMutation
}

// Where appends a list predicates to the ReactionUpdate builder.
func (ru *ReactionUpdate) Where(ps ...predicate.Reaction) *ReactionUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUpdateTime sets the "update_time" field.
func (ru *ReactionUpdate) SetUpdateTime(t time.Time) *ReactionUpdate {
	ru.mutation.SetUpdateTime(t)
	return ru
}

// SetIsLOL sets the "isLOL" field.
func (ru *ReactionUpdate) SetIsLOL(b bool) *ReactionUpdate {
	ru.mutation.SetIsLOL(b)
	return ru
}

// SetNillableIsLOL sets the "isLOL" field if the given value is not nil.
func (ru *ReactionUpdate) SetNillableIsLOL(b *bool) *ReactionUpdate {
	if b != nil {
		ru.SetIsLOL(*b)
	}
	return ru
}

// SetIsOMG sets the "isOMG" field.
func (ru *ReactionUpdate) SetIsOMG(b bool) *ReactionUpdate {
	ru.mutation.SetIsOMG(b)
	return ru
}

// SetNillableIsOMG sets the "isOMG" field if the given value is not nil.
func (ru *ReactionUpdate) SetNillableIsOMG(b *bool) *ReactionUpdate {
	if b != nil {
		ru.SetIsOMG(*b)
	}
	return ru
}

// SetIsCool sets the "isCool" field.
func (ru *ReactionUpdate) SetIsCool(b bool) *ReactionUpdate {
	ru.mutation.SetIsCool(b)
	return ru
}

// SetNillableIsCool sets the "isCool" field if the given value is not nil.
func (ru *ReactionUpdate) SetNillableIsCool(b *bool) *ReactionUpdate {
	if b != nil {
		ru.SetIsCool(*b)
	}
	return ru
}

// SetIsNooo sets the "isNooo" field.
func (ru *ReactionUpdate) SetIsNooo(b bool) *ReactionUpdate {
	ru.mutation.SetIsNooo(b)
	return ru
}

// SetNillableIsNooo sets the "isNooo" field if the given value is not nil.
func (ru *ReactionUpdate) SetNillableIsNooo(b *bool) *ReactionUpdate {
	if b != nil {
		ru.SetIsNooo(*b)
	}
	return ru
}

// SetIsDAMN sets the "isDAMN" field.
func (ru *ReactionUpdate) SetIsDAMN(b bool) *ReactionUpdate {
	ru.mutation.SetIsDAMN(b)
	return ru
}

// SetNillableIsDAMN sets the "isDAMN" field if the given value is not nil.
func (ru *ReactionUpdate) SetNillableIsDAMN(b *bool) *ReactionUpdate {
	if b != nil {
		ru.SetIsDAMN(*b)
	}
	return ru
}

// SetTimeDewID sets the "time_dew_id" field.
func (ru *ReactionUpdate) SetTimeDewID(u uint64) *ReactionUpdate {
	ru.mutation.SetTimeDewID(u)
	return ru
}

// SetUserID sets the "user_id" field.
func (ru *ReactionUpdate) SetUserID(u uint64) *ReactionUpdate {
	ru.mutation.SetUserID(u)
	return ru
}

// SetTimedewID sets the "timedew" edge to the TimeDew entity by ID.
func (ru *ReactionUpdate) SetTimedewID(id uint64) *ReactionUpdate {
	ru.mutation.SetTimedewID(id)
	return ru
}

// SetTimedew sets the "timedew" edge to the TimeDew entity.
func (ru *ReactionUpdate) SetTimedew(t *TimeDew) *ReactionUpdate {
	return ru.SetTimedewID(t.ID)
}

// SetUser sets the "user" edge to the User entity.
func (ru *ReactionUpdate) SetUser(u *User) *ReactionUpdate {
	return ru.SetUserID(u.ID)
}

// Mutation returns the ReactionMutation object of the builder.
func (ru *ReactionUpdate) Mutation() *ReactionMutation {
	return ru.mutation
}

// ClearTimedew clears the "timedew" edge to the TimeDew entity.
func (ru *ReactionUpdate) ClearTimedew() *ReactionUpdate {
	ru.mutation.ClearTimedew()
	return ru
}

// ClearUser clears the "user" edge to the User entity.
func (ru *ReactionUpdate) ClearUser() *ReactionUpdate {
	ru.mutation.ClearUser()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ReactionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ru.defaults()
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReactionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReactionUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReactionUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReactionUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *ReactionUpdate) defaults() {
	if _, ok := ru.mutation.UpdateTime(); !ok {
		v := reaction.UpdateDefaultUpdateTime()
		ru.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *ReactionUpdate) check() error {
	if _, ok := ru.mutation.TimedewID(); ru.mutation.TimedewCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Reaction.timedew"`)
	}
	if _, ok := ru.mutation.UserID(); ru.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Reaction.user"`)
	}
	return nil
}

func (ru *ReactionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   reaction.Table,
			Columns: reaction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: reaction.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UpdateTime(); ok {
		_spec.SetField(reaction.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := ru.mutation.IsLOL(); ok {
		_spec.SetField(reaction.FieldIsLOL, field.TypeBool, value)
	}
	if value, ok := ru.mutation.IsOMG(); ok {
		_spec.SetField(reaction.FieldIsOMG, field.TypeBool, value)
	}
	if value, ok := ru.mutation.IsCool(); ok {
		_spec.SetField(reaction.FieldIsCool, field.TypeBool, value)
	}
	if value, ok := ru.mutation.IsNooo(); ok {
		_spec.SetField(reaction.FieldIsNooo, field.TypeBool, value)
	}
	if value, ok := ru.mutation.IsDAMN(); ok {
		_spec.SetField(reaction.FieldIsDAMN, field.TypeBool, value)
	}
	if ru.mutation.TimedewCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   reaction.TimedewTable,
			Columns: []string{reaction.TimedewColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: timedew.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.TimedewIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   reaction.TimedewTable,
			Columns: []string{reaction.TimedewColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: timedew.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   reaction.UserTable,
			Columns: []string{reaction.UserColumn},
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
	if nodes := ru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   reaction.UserTable,
			Columns: []string{reaction.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ReactionUpdateOne is the builder for updating a single Reaction entity.
type ReactionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReactionMutation
}

// SetUpdateTime sets the "update_time" field.
func (ruo *ReactionUpdateOne) SetUpdateTime(t time.Time) *ReactionUpdateOne {
	ruo.mutation.SetUpdateTime(t)
	return ruo
}

// SetIsLOL sets the "isLOL" field.
func (ruo *ReactionUpdateOne) SetIsLOL(b bool) *ReactionUpdateOne {
	ruo.mutation.SetIsLOL(b)
	return ruo
}

// SetNillableIsLOL sets the "isLOL" field if the given value is not nil.
func (ruo *ReactionUpdateOne) SetNillableIsLOL(b *bool) *ReactionUpdateOne {
	if b != nil {
		ruo.SetIsLOL(*b)
	}
	return ruo
}

// SetIsOMG sets the "isOMG" field.
func (ruo *ReactionUpdateOne) SetIsOMG(b bool) *ReactionUpdateOne {
	ruo.mutation.SetIsOMG(b)
	return ruo
}

// SetNillableIsOMG sets the "isOMG" field if the given value is not nil.
func (ruo *ReactionUpdateOne) SetNillableIsOMG(b *bool) *ReactionUpdateOne {
	if b != nil {
		ruo.SetIsOMG(*b)
	}
	return ruo
}

// SetIsCool sets the "isCool" field.
func (ruo *ReactionUpdateOne) SetIsCool(b bool) *ReactionUpdateOne {
	ruo.mutation.SetIsCool(b)
	return ruo
}

// SetNillableIsCool sets the "isCool" field if the given value is not nil.
func (ruo *ReactionUpdateOne) SetNillableIsCool(b *bool) *ReactionUpdateOne {
	if b != nil {
		ruo.SetIsCool(*b)
	}
	return ruo
}

// SetIsNooo sets the "isNooo" field.
func (ruo *ReactionUpdateOne) SetIsNooo(b bool) *ReactionUpdateOne {
	ruo.mutation.SetIsNooo(b)
	return ruo
}

// SetNillableIsNooo sets the "isNooo" field if the given value is not nil.
func (ruo *ReactionUpdateOne) SetNillableIsNooo(b *bool) *ReactionUpdateOne {
	if b != nil {
		ruo.SetIsNooo(*b)
	}
	return ruo
}

// SetIsDAMN sets the "isDAMN" field.
func (ruo *ReactionUpdateOne) SetIsDAMN(b bool) *ReactionUpdateOne {
	ruo.mutation.SetIsDAMN(b)
	return ruo
}

// SetNillableIsDAMN sets the "isDAMN" field if the given value is not nil.
func (ruo *ReactionUpdateOne) SetNillableIsDAMN(b *bool) *ReactionUpdateOne {
	if b != nil {
		ruo.SetIsDAMN(*b)
	}
	return ruo
}

// SetTimeDewID sets the "time_dew_id" field.
func (ruo *ReactionUpdateOne) SetTimeDewID(u uint64) *ReactionUpdateOne {
	ruo.mutation.SetTimeDewID(u)
	return ruo
}

// SetUserID sets the "user_id" field.
func (ruo *ReactionUpdateOne) SetUserID(u uint64) *ReactionUpdateOne {
	ruo.mutation.SetUserID(u)
	return ruo
}

// SetTimedewID sets the "timedew" edge to the TimeDew entity by ID.
func (ruo *ReactionUpdateOne) SetTimedewID(id uint64) *ReactionUpdateOne {
	ruo.mutation.SetTimedewID(id)
	return ruo
}

// SetTimedew sets the "timedew" edge to the TimeDew entity.
func (ruo *ReactionUpdateOne) SetTimedew(t *TimeDew) *ReactionUpdateOne {
	return ruo.SetTimedewID(t.ID)
}

// SetUser sets the "user" edge to the User entity.
func (ruo *ReactionUpdateOne) SetUser(u *User) *ReactionUpdateOne {
	return ruo.SetUserID(u.ID)
}

// Mutation returns the ReactionMutation object of the builder.
func (ruo *ReactionUpdateOne) Mutation() *ReactionMutation {
	return ruo.mutation
}

// ClearTimedew clears the "timedew" edge to the TimeDew entity.
func (ruo *ReactionUpdateOne) ClearTimedew() *ReactionUpdateOne {
	ruo.mutation.ClearTimedew()
	return ruo
}

// ClearUser clears the "user" edge to the User entity.
func (ruo *ReactionUpdateOne) ClearUser() *ReactionUpdateOne {
	ruo.mutation.ClearUser()
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ReactionUpdateOne) Select(field string, fields ...string) *ReactionUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Reaction entity.
func (ruo *ReactionUpdateOne) Save(ctx context.Context) (*Reaction, error) {
	var (
		err  error
		node *Reaction
	)
	ruo.defaults()
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReactionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Reaction)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ReactionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReactionUpdateOne) SaveX(ctx context.Context) *Reaction {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ReactionUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReactionUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *ReactionUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdateTime(); !ok {
		v := reaction.UpdateDefaultUpdateTime()
		ruo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *ReactionUpdateOne) check() error {
	if _, ok := ruo.mutation.TimedewID(); ruo.mutation.TimedewCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Reaction.timedew"`)
	}
	if _, ok := ruo.mutation.UserID(); ruo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Reaction.user"`)
	}
	return nil
}

func (ruo *ReactionUpdateOne) sqlSave(ctx context.Context) (_node *Reaction, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   reaction.Table,
			Columns: reaction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: reaction.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Reaction.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, reaction.FieldID)
		for _, f := range fields {
			if !reaction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != reaction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UpdateTime(); ok {
		_spec.SetField(reaction.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.IsLOL(); ok {
		_spec.SetField(reaction.FieldIsLOL, field.TypeBool, value)
	}
	if value, ok := ruo.mutation.IsOMG(); ok {
		_spec.SetField(reaction.FieldIsOMG, field.TypeBool, value)
	}
	if value, ok := ruo.mutation.IsCool(); ok {
		_spec.SetField(reaction.FieldIsCool, field.TypeBool, value)
	}
	if value, ok := ruo.mutation.IsNooo(); ok {
		_spec.SetField(reaction.FieldIsNooo, field.TypeBool, value)
	}
	if value, ok := ruo.mutation.IsDAMN(); ok {
		_spec.SetField(reaction.FieldIsDAMN, field.TypeBool, value)
	}
	if ruo.mutation.TimedewCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   reaction.TimedewTable,
			Columns: []string{reaction.TimedewColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: timedew.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.TimedewIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   reaction.TimedewTable,
			Columns: []string{reaction.TimedewColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: timedew.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   reaction.UserTable,
			Columns: []string{reaction.UserColumn},
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
	if nodes := ruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   reaction.UserTable,
			Columns: []string{reaction.UserColumn},
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
	_node = &Reaction{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
