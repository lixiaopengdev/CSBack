// Code generated by ent, DO NOT EDIT.

package ent

import (
	"CSBackendTmp/ent/card"
	"CSBackendTmp/ent/predicate"
	"CSBackendTmp/ent/rule"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RuleUpdate is the builder for updating Rule entities.
type RuleUpdate struct {
	config
	hooks    []Hook
	mutation *RuleMutation
}

// Where appends a list predicates to the RuleUpdate builder.
func (ru *RuleUpdate) Where(ps ...predicate.Rule) *RuleUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUpdateTime sets the "update_time" field.
func (ru *RuleUpdate) SetUpdateTime(t time.Time) *RuleUpdate {
	ru.mutation.SetUpdateTime(t)
	return ru
}

// SetName sets the "name" field.
func (ru *RuleUpdate) SetName(s string) *RuleUpdate {
	ru.mutation.SetName(s)
	return ru
}

// AddUsedIDs adds the "used" edge to the Card entity by IDs.
func (ru *RuleUpdate) AddUsedIDs(ids ...uint64) *RuleUpdate {
	ru.mutation.AddUsedIDs(ids...)
	return ru
}

// AddUsed adds the "used" edges to the Card entity.
func (ru *RuleUpdate) AddUsed(c ...*Card) *RuleUpdate {
	ids := make([]uint64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ru.AddUsedIDs(ids...)
}

// Mutation returns the RuleMutation object of the builder.
func (ru *RuleUpdate) Mutation() *RuleMutation {
	return ru.mutation
}

// ClearUsed clears all "used" edges to the Card entity.
func (ru *RuleUpdate) ClearUsed() *RuleUpdate {
	ru.mutation.ClearUsed()
	return ru
}

// RemoveUsedIDs removes the "used" edge to Card entities by IDs.
func (ru *RuleUpdate) RemoveUsedIDs(ids ...uint64) *RuleUpdate {
	ru.mutation.RemoveUsedIDs(ids...)
	return ru
}

// RemoveUsed removes "used" edges to Card entities.
func (ru *RuleUpdate) RemoveUsed(c ...*Card) *RuleUpdate {
	ids := make([]uint64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ru.RemoveUsedIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RuleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ru.defaults()
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
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
func (ru *RuleUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RuleUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RuleUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RuleUpdate) defaults() {
	if _, ok := ru.mutation.UpdateTime(); !ok {
		v := rule.UpdateDefaultUpdateTime()
		ru.mutation.SetUpdateTime(v)
	}
}

func (ru *RuleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rule.Table,
			Columns: rule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: rule.FieldID,
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
		_spec.SetField(rule.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(rule.FieldName, field.TypeString, value)
	}
	if ru.mutation.UsedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   rule.UsedTable,
			Columns: rule.UsedPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: card.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedUsedIDs(); len(nodes) > 0 && !ru.mutation.UsedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   rule.UsedTable,
			Columns: rule.UsedPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.UsedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   rule.UsedTable,
			Columns: rule.UsedPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: card.FieldID,
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
			err = &NotFoundError{rule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// RuleUpdateOne is the builder for updating a single Rule entity.
type RuleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RuleMutation
}

// SetUpdateTime sets the "update_time" field.
func (ruo *RuleUpdateOne) SetUpdateTime(t time.Time) *RuleUpdateOne {
	ruo.mutation.SetUpdateTime(t)
	return ruo
}

// SetName sets the "name" field.
func (ruo *RuleUpdateOne) SetName(s string) *RuleUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// AddUsedIDs adds the "used" edge to the Card entity by IDs.
func (ruo *RuleUpdateOne) AddUsedIDs(ids ...uint64) *RuleUpdateOne {
	ruo.mutation.AddUsedIDs(ids...)
	return ruo
}

// AddUsed adds the "used" edges to the Card entity.
func (ruo *RuleUpdateOne) AddUsed(c ...*Card) *RuleUpdateOne {
	ids := make([]uint64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ruo.AddUsedIDs(ids...)
}

// Mutation returns the RuleMutation object of the builder.
func (ruo *RuleUpdateOne) Mutation() *RuleMutation {
	return ruo.mutation
}

// ClearUsed clears all "used" edges to the Card entity.
func (ruo *RuleUpdateOne) ClearUsed() *RuleUpdateOne {
	ruo.mutation.ClearUsed()
	return ruo
}

// RemoveUsedIDs removes the "used" edge to Card entities by IDs.
func (ruo *RuleUpdateOne) RemoveUsedIDs(ids ...uint64) *RuleUpdateOne {
	ruo.mutation.RemoveUsedIDs(ids...)
	return ruo
}

// RemoveUsed removes "used" edges to Card entities.
func (ruo *RuleUpdateOne) RemoveUsed(c ...*Card) *RuleUpdateOne {
	ids := make([]uint64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ruo.RemoveUsedIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RuleUpdateOne) Select(field string, fields ...string) *RuleUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Rule entity.
func (ruo *RuleUpdateOne) Save(ctx context.Context) (*Rule, error) {
	var (
		err  error
		node *Rule
	)
	ruo.defaults()
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
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
		nv, ok := v.(*Rule)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RuleMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RuleUpdateOne) SaveX(ctx context.Context) *Rule {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RuleUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RuleUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RuleUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdateTime(); !ok {
		v := rule.UpdateDefaultUpdateTime()
		ruo.mutation.SetUpdateTime(v)
	}
}

func (ruo *RuleUpdateOne) sqlSave(ctx context.Context) (_node *Rule, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rule.Table,
			Columns: rule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: rule.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Rule.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rule.FieldID)
		for _, f := range fields {
			if !rule.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != rule.FieldID {
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
		_spec.SetField(rule.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(rule.FieldName, field.TypeString, value)
	}
	if ruo.mutation.UsedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   rule.UsedTable,
			Columns: rule.UsedPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: card.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedUsedIDs(); len(nodes) > 0 && !ruo.mutation.UsedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   rule.UsedTable,
			Columns: rule.UsedPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.UsedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   rule.UsedTable,
			Columns: rule.UsedPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Rule{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}