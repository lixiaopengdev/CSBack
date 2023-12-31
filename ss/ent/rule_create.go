// Code generated by ent, DO NOT EDIT.

package ent

import (
	"CSBackendTmp/ent/card"
	"CSBackendTmp/ent/rule"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RuleCreate is the builder for creating a Rule entity.
type RuleCreate struct {
	config
	mutation *RuleMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (rc *RuleCreate) SetCreateTime(t time.Time) *RuleCreate {
	rc.mutation.SetCreateTime(t)
	return rc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (rc *RuleCreate) SetNillableCreateTime(t *time.Time) *RuleCreate {
	if t != nil {
		rc.SetCreateTime(*t)
	}
	return rc
}

// SetUpdateTime sets the "update_time" field.
func (rc *RuleCreate) SetUpdateTime(t time.Time) *RuleCreate {
	rc.mutation.SetUpdateTime(t)
	return rc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (rc *RuleCreate) SetNillableUpdateTime(t *time.Time) *RuleCreate {
	if t != nil {
		rc.SetUpdateTime(*t)
	}
	return rc
}

// SetName sets the "name" field.
func (rc *RuleCreate) SetName(s string) *RuleCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetID sets the "id" field.
func (rc *RuleCreate) SetID(u uint64) *RuleCreate {
	rc.mutation.SetID(u)
	return rc
}

// AddUsedIDs adds the "used" edge to the Card entity by IDs.
func (rc *RuleCreate) AddUsedIDs(ids ...uint64) *RuleCreate {
	rc.mutation.AddUsedIDs(ids...)
	return rc
}

// AddUsed adds the "used" edges to the Card entity.
func (rc *RuleCreate) AddUsed(c ...*Card) *RuleCreate {
	ids := make([]uint64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return rc.AddUsedIDs(ids...)
}

// Mutation returns the RuleMutation object of the builder.
func (rc *RuleCreate) Mutation() *RuleMutation {
	return rc.mutation
}

// Save creates the Rule in the database.
func (rc *RuleCreate) Save(ctx context.Context) (*Rule, error) {
	var (
		err  error
		node *Rule
	)
	rc.defaults()
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (rc *RuleCreate) SaveX(ctx context.Context) *Rule {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RuleCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RuleCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RuleCreate) defaults() {
	if _, ok := rc.mutation.CreateTime(); !ok {
		v := rule.DefaultCreateTime()
		rc.mutation.SetCreateTime(v)
	}
	if _, ok := rc.mutation.UpdateTime(); !ok {
		v := rule.DefaultUpdateTime()
		rc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RuleCreate) check() error {
	if _, ok := rc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Rule.create_time"`)}
	}
	if _, ok := rc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Rule.update_time"`)}
	}
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Rule.name"`)}
	}
	return nil
}

func (rc *RuleCreate) sqlSave(ctx context.Context) (*Rule, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	return _node, nil
}

func (rc *RuleCreate) createSpec() (*Rule, *sqlgraph.CreateSpec) {
	var (
		_node = &Rule{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: rule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: rule.FieldID,
			},
		}
	)
	_spec.OnConflict = rc.conflict
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.CreateTime(); ok {
		_spec.SetField(rule.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := rc.mutation.UpdateTime(); ok {
		_spec.SetField(rule.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.SetField(rule.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := rc.mutation.UsedIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Rule.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RuleUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (rc *RuleCreate) OnConflict(opts ...sql.ConflictOption) *RuleUpsertOne {
	rc.conflict = opts
	return &RuleUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Rule.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rc *RuleCreate) OnConflictColumns(columns ...string) *RuleUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &RuleUpsertOne{
		create: rc,
	}
}

type (
	// RuleUpsertOne is the builder for "upsert"-ing
	//  one Rule node.
	RuleUpsertOne struct {
		create *RuleCreate
	}

	// RuleUpsert is the "OnConflict" setter.
	RuleUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *RuleUpsert) SetUpdateTime(v time.Time) *RuleUpsert {
	u.Set(rule.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *RuleUpsert) UpdateUpdateTime() *RuleUpsert {
	u.SetExcluded(rule.FieldUpdateTime)
	return u
}

// SetName sets the "name" field.
func (u *RuleUpsert) SetName(v string) *RuleUpsert {
	u.Set(rule.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RuleUpsert) UpdateName() *RuleUpsert {
	u.SetExcluded(rule.FieldName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Rule.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(rule.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *RuleUpsertOne) UpdateNewValues() *RuleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(rule.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(rule.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Rule.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *RuleUpsertOne) Ignore() *RuleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RuleUpsertOne) DoNothing() *RuleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RuleCreate.OnConflict
// documentation for more info.
func (u *RuleUpsertOne) Update(set func(*RuleUpsert)) *RuleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RuleUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *RuleUpsertOne) SetUpdateTime(v time.Time) *RuleUpsertOne {
	return u.Update(func(s *RuleUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *RuleUpsertOne) UpdateUpdateTime() *RuleUpsertOne {
	return u.Update(func(s *RuleUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *RuleUpsertOne) SetName(v string) *RuleUpsertOne {
	return u.Update(func(s *RuleUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RuleUpsertOne) UpdateName() *RuleUpsertOne {
	return u.Update(func(s *RuleUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *RuleUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RuleCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RuleUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *RuleUpsertOne) ID(ctx context.Context) (id uint64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *RuleUpsertOne) IDX(ctx context.Context) uint64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// RuleCreateBulk is the builder for creating many Rule entities in bulk.
type RuleCreateBulk struct {
	config
	builders []*RuleCreate
	conflict []sql.ConflictOption
}

// Save creates the Rule entities in the database.
func (rcb *RuleCreateBulk) Save(ctx context.Context) ([]*Rule, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Rule, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RuleMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RuleCreateBulk) SaveX(ctx context.Context) []*Rule {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RuleCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RuleCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Rule.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RuleUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (rcb *RuleCreateBulk) OnConflict(opts ...sql.ConflictOption) *RuleUpsertBulk {
	rcb.conflict = opts
	return &RuleUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Rule.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rcb *RuleCreateBulk) OnConflictColumns(columns ...string) *RuleUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &RuleUpsertBulk{
		create: rcb,
	}
}

// RuleUpsertBulk is the builder for "upsert"-ing
// a bulk of Rule nodes.
type RuleUpsertBulk struct {
	create *RuleCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Rule.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(rule.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *RuleUpsertBulk) UpdateNewValues() *RuleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(rule.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(rule.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Rule.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *RuleUpsertBulk) Ignore() *RuleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RuleUpsertBulk) DoNothing() *RuleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RuleCreateBulk.OnConflict
// documentation for more info.
func (u *RuleUpsertBulk) Update(set func(*RuleUpsert)) *RuleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RuleUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *RuleUpsertBulk) SetUpdateTime(v time.Time) *RuleUpsertBulk {
	return u.Update(func(s *RuleUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *RuleUpsertBulk) UpdateUpdateTime() *RuleUpsertBulk {
	return u.Update(func(s *RuleUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *RuleUpsertBulk) SetName(v string) *RuleUpsertBulk {
	return u.Update(func(s *RuleUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RuleUpsertBulk) UpdateName() *RuleUpsertBulk {
	return u.Update(func(s *RuleUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *RuleUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the RuleCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RuleCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RuleUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
