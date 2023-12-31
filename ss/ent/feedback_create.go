// Code generated by ent, DO NOT EDIT.

package ent

import (
	"CSBackendTmp/ent/feedback"
	"CSBackendTmp/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FeedbackCreate is the builder for creating a Feedback entity.
type FeedbackCreate struct {
	config
	mutation *FeedbackMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (fc *FeedbackCreate) SetCreateTime(t time.Time) *FeedbackCreate {
	fc.mutation.SetCreateTime(t)
	return fc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (fc *FeedbackCreate) SetNillableCreateTime(t *time.Time) *FeedbackCreate {
	if t != nil {
		fc.SetCreateTime(*t)
	}
	return fc
}

// SetUpdateTime sets the "update_time" field.
func (fc *FeedbackCreate) SetUpdateTime(t time.Time) *FeedbackCreate {
	fc.mutation.SetUpdateTime(t)
	return fc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (fc *FeedbackCreate) SetNillableUpdateTime(t *time.Time) *FeedbackCreate {
	if t != nil {
		fc.SetUpdateTime(*t)
	}
	return fc
}

// SetType sets the "type" field.
func (fc *FeedbackCreate) SetType(f feedback.Type) *FeedbackCreate {
	fc.mutation.SetType(f)
	return fc
}

// SetReason sets the "reason" field.
func (fc *FeedbackCreate) SetReason(s string) *FeedbackCreate {
	fc.mutation.SetReason(s)
	return fc
}

// SetItemID sets the "item_id" field.
func (fc *FeedbackCreate) SetItemID(u uint64) *FeedbackCreate {
	fc.mutation.SetItemID(u)
	return fc
}

// SetNillableItemID sets the "item_id" field if the given value is not nil.
func (fc *FeedbackCreate) SetNillableItemID(u *uint64) *FeedbackCreate {
	if u != nil {
		fc.SetItemID(*u)
	}
	return fc
}

// SetUserID sets the "user_id" field.
func (fc *FeedbackCreate) SetUserID(u uint64) *FeedbackCreate {
	fc.mutation.SetUserID(u)
	return fc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (fc *FeedbackCreate) SetOwnerID(id uint64) *FeedbackCreate {
	fc.mutation.SetOwnerID(id)
	return fc
}

// SetOwner sets the "owner" edge to the User entity.
func (fc *FeedbackCreate) SetOwner(u *User) *FeedbackCreate {
	return fc.SetOwnerID(u.ID)
}

// Mutation returns the FeedbackMutation object of the builder.
func (fc *FeedbackCreate) Mutation() *FeedbackMutation {
	return fc.mutation
}

// Save creates the Feedback in the database.
func (fc *FeedbackCreate) Save(ctx context.Context) (*Feedback, error) {
	var (
		err  error
		node *Feedback
	)
	fc.defaults()
	if len(fc.hooks) == 0 {
		if err = fc.check(); err != nil {
			return nil, err
		}
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FeedbackMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fc.check(); err != nil {
				return nil, err
			}
			fc.mutation = mutation
			if node, err = fc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			if fc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, fc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Feedback)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from FeedbackMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FeedbackCreate) SaveX(ctx context.Context) *Feedback {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FeedbackCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FeedbackCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FeedbackCreate) defaults() {
	if _, ok := fc.mutation.CreateTime(); !ok {
		v := feedback.DefaultCreateTime()
		fc.mutation.SetCreateTime(v)
	}
	if _, ok := fc.mutation.UpdateTime(); !ok {
		v := feedback.DefaultUpdateTime()
		fc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FeedbackCreate) check() error {
	if _, ok := fc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Feedback.create_time"`)}
	}
	if _, ok := fc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Feedback.update_time"`)}
	}
	if _, ok := fc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Feedback.type"`)}
	}
	if v, ok := fc.mutation.GetType(); ok {
		if err := feedback.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Feedback.type": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Reason(); !ok {
		return &ValidationError{Name: "reason", err: errors.New(`ent: missing required field "Feedback.reason"`)}
	}
	if _, ok := fc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Feedback.user_id"`)}
	}
	if _, ok := fc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Feedback.owner"`)}
	}
	return nil
}

func (fc *FeedbackCreate) sqlSave(ctx context.Context) (*Feedback, error) {
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (fc *FeedbackCreate) createSpec() (*Feedback, *sqlgraph.CreateSpec) {
	var (
		_node = &Feedback{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: feedback.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: feedback.FieldID,
			},
		}
	)
	_spec.OnConflict = fc.conflict
	if value, ok := fc.mutation.CreateTime(); ok {
		_spec.SetField(feedback.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := fc.mutation.UpdateTime(); ok {
		_spec.SetField(feedback.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := fc.mutation.GetType(); ok {
		_spec.SetField(feedback.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := fc.mutation.Reason(); ok {
		_spec.SetField(feedback.FieldReason, field.TypeString, value)
		_node.Reason = value
	}
	if value, ok := fc.mutation.ItemID(); ok {
		_spec.SetField(feedback.FieldItemID, field.TypeUint64, value)
		_node.ItemID = value
	}
	if nodes := fc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   feedback.OwnerTable,
			Columns: []string{feedback.OwnerColumn},
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
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Feedback.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FeedbackUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (fc *FeedbackCreate) OnConflict(opts ...sql.ConflictOption) *FeedbackUpsertOne {
	fc.conflict = opts
	return &FeedbackUpsertOne{
		create: fc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Feedback.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fc *FeedbackCreate) OnConflictColumns(columns ...string) *FeedbackUpsertOne {
	fc.conflict = append(fc.conflict, sql.ConflictColumns(columns...))
	return &FeedbackUpsertOne{
		create: fc,
	}
}

type (
	// FeedbackUpsertOne is the builder for "upsert"-ing
	//  one Feedback node.
	FeedbackUpsertOne struct {
		create *FeedbackCreate
	}

	// FeedbackUpsert is the "OnConflict" setter.
	FeedbackUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *FeedbackUpsert) SetUpdateTime(v time.Time) *FeedbackUpsert {
	u.Set(feedback.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *FeedbackUpsert) UpdateUpdateTime() *FeedbackUpsert {
	u.SetExcluded(feedback.FieldUpdateTime)
	return u
}

// SetType sets the "type" field.
func (u *FeedbackUpsert) SetType(v feedback.Type) *FeedbackUpsert {
	u.Set(feedback.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *FeedbackUpsert) UpdateType() *FeedbackUpsert {
	u.SetExcluded(feedback.FieldType)
	return u
}

// SetReason sets the "reason" field.
func (u *FeedbackUpsert) SetReason(v string) *FeedbackUpsert {
	u.Set(feedback.FieldReason, v)
	return u
}

// UpdateReason sets the "reason" field to the value that was provided on create.
func (u *FeedbackUpsert) UpdateReason() *FeedbackUpsert {
	u.SetExcluded(feedback.FieldReason)
	return u
}

// SetItemID sets the "item_id" field.
func (u *FeedbackUpsert) SetItemID(v uint64) *FeedbackUpsert {
	u.Set(feedback.FieldItemID, v)
	return u
}

// UpdateItemID sets the "item_id" field to the value that was provided on create.
func (u *FeedbackUpsert) UpdateItemID() *FeedbackUpsert {
	u.SetExcluded(feedback.FieldItemID)
	return u
}

// AddItemID adds v to the "item_id" field.
func (u *FeedbackUpsert) AddItemID(v uint64) *FeedbackUpsert {
	u.Add(feedback.FieldItemID, v)
	return u
}

// ClearItemID clears the value of the "item_id" field.
func (u *FeedbackUpsert) ClearItemID() *FeedbackUpsert {
	u.SetNull(feedback.FieldItemID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *FeedbackUpsert) SetUserID(v uint64) *FeedbackUpsert {
	u.Set(feedback.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *FeedbackUpsert) UpdateUserID() *FeedbackUpsert {
	u.SetExcluded(feedback.FieldUserID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Feedback.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *FeedbackUpsertOne) UpdateNewValues() *FeedbackUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(feedback.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Feedback.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FeedbackUpsertOne) Ignore() *FeedbackUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FeedbackUpsertOne) DoNothing() *FeedbackUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FeedbackCreate.OnConflict
// documentation for more info.
func (u *FeedbackUpsertOne) Update(set func(*FeedbackUpsert)) *FeedbackUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FeedbackUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *FeedbackUpsertOne) SetUpdateTime(v time.Time) *FeedbackUpsertOne {
	return u.Update(func(s *FeedbackUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *FeedbackUpsertOne) UpdateUpdateTime() *FeedbackUpsertOne {
	return u.Update(func(s *FeedbackUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetType sets the "type" field.
func (u *FeedbackUpsertOne) SetType(v feedback.Type) *FeedbackUpsertOne {
	return u.Update(func(s *FeedbackUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *FeedbackUpsertOne) UpdateType() *FeedbackUpsertOne {
	return u.Update(func(s *FeedbackUpsert) {
		s.UpdateType()
	})
}

// SetReason sets the "reason" field.
func (u *FeedbackUpsertOne) SetReason(v string) *FeedbackUpsertOne {
	return u.Update(func(s *FeedbackUpsert) {
		s.SetReason(v)
	})
}

// UpdateReason sets the "reason" field to the value that was provided on create.
func (u *FeedbackUpsertOne) UpdateReason() *FeedbackUpsertOne {
	return u.Update(func(s *FeedbackUpsert) {
		s.UpdateReason()
	})
}

// SetItemID sets the "item_id" field.
func (u *FeedbackUpsertOne) SetItemID(v uint64) *FeedbackUpsertOne {
	return u.Update(func(s *FeedbackUpsert) {
		s.SetItemID(v)
	})
}

// AddItemID adds v to the "item_id" field.
func (u *FeedbackUpsertOne) AddItemID(v uint64) *FeedbackUpsertOne {
	return u.Update(func(s *FeedbackUpsert) {
		s.AddItemID(v)
	})
}

// UpdateItemID sets the "item_id" field to the value that was provided on create.
func (u *FeedbackUpsertOne) UpdateItemID() *FeedbackUpsertOne {
	return u.Update(func(s *FeedbackUpsert) {
		s.UpdateItemID()
	})
}

// ClearItemID clears the value of the "item_id" field.
func (u *FeedbackUpsertOne) ClearItemID() *FeedbackUpsertOne {
	return u.Update(func(s *FeedbackUpsert) {
		s.ClearItemID()
	})
}

// SetUserID sets the "user_id" field.
func (u *FeedbackUpsertOne) SetUserID(v uint64) *FeedbackUpsertOne {
	return u.Update(func(s *FeedbackUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *FeedbackUpsertOne) UpdateUserID() *FeedbackUpsertOne {
	return u.Update(func(s *FeedbackUpsert) {
		s.UpdateUserID()
	})
}

// Exec executes the query.
func (u *FeedbackUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FeedbackCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FeedbackUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FeedbackUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FeedbackUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FeedbackCreateBulk is the builder for creating many Feedback entities in bulk.
type FeedbackCreateBulk struct {
	config
	builders []*FeedbackCreate
	conflict []sql.ConflictOption
}

// Save creates the Feedback entities in the database.
func (fcb *FeedbackCreateBulk) Save(ctx context.Context) ([]*Feedback, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Feedback, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FeedbackMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FeedbackCreateBulk) SaveX(ctx context.Context) []*Feedback {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FeedbackCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FeedbackCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Feedback.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FeedbackUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (fcb *FeedbackCreateBulk) OnConflict(opts ...sql.ConflictOption) *FeedbackUpsertBulk {
	fcb.conflict = opts
	return &FeedbackUpsertBulk{
		create: fcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Feedback.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fcb *FeedbackCreateBulk) OnConflictColumns(columns ...string) *FeedbackUpsertBulk {
	fcb.conflict = append(fcb.conflict, sql.ConflictColumns(columns...))
	return &FeedbackUpsertBulk{
		create: fcb,
	}
}

// FeedbackUpsertBulk is the builder for "upsert"-ing
// a bulk of Feedback nodes.
type FeedbackUpsertBulk struct {
	create *FeedbackCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Feedback.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *FeedbackUpsertBulk) UpdateNewValues() *FeedbackUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(feedback.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Feedback.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FeedbackUpsertBulk) Ignore() *FeedbackUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FeedbackUpsertBulk) DoNothing() *FeedbackUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FeedbackCreateBulk.OnConflict
// documentation for more info.
func (u *FeedbackUpsertBulk) Update(set func(*FeedbackUpsert)) *FeedbackUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FeedbackUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *FeedbackUpsertBulk) SetUpdateTime(v time.Time) *FeedbackUpsertBulk {
	return u.Update(func(s *FeedbackUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *FeedbackUpsertBulk) UpdateUpdateTime() *FeedbackUpsertBulk {
	return u.Update(func(s *FeedbackUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetType sets the "type" field.
func (u *FeedbackUpsertBulk) SetType(v feedback.Type) *FeedbackUpsertBulk {
	return u.Update(func(s *FeedbackUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *FeedbackUpsertBulk) UpdateType() *FeedbackUpsertBulk {
	return u.Update(func(s *FeedbackUpsert) {
		s.UpdateType()
	})
}

// SetReason sets the "reason" field.
func (u *FeedbackUpsertBulk) SetReason(v string) *FeedbackUpsertBulk {
	return u.Update(func(s *FeedbackUpsert) {
		s.SetReason(v)
	})
}

// UpdateReason sets the "reason" field to the value that was provided on create.
func (u *FeedbackUpsertBulk) UpdateReason() *FeedbackUpsertBulk {
	return u.Update(func(s *FeedbackUpsert) {
		s.UpdateReason()
	})
}

// SetItemID sets the "item_id" field.
func (u *FeedbackUpsertBulk) SetItemID(v uint64) *FeedbackUpsertBulk {
	return u.Update(func(s *FeedbackUpsert) {
		s.SetItemID(v)
	})
}

// AddItemID adds v to the "item_id" field.
func (u *FeedbackUpsertBulk) AddItemID(v uint64) *FeedbackUpsertBulk {
	return u.Update(func(s *FeedbackUpsert) {
		s.AddItemID(v)
	})
}

// UpdateItemID sets the "item_id" field to the value that was provided on create.
func (u *FeedbackUpsertBulk) UpdateItemID() *FeedbackUpsertBulk {
	return u.Update(func(s *FeedbackUpsert) {
		s.UpdateItemID()
	})
}

// ClearItemID clears the value of the "item_id" field.
func (u *FeedbackUpsertBulk) ClearItemID() *FeedbackUpsertBulk {
	return u.Update(func(s *FeedbackUpsert) {
		s.ClearItemID()
	})
}

// SetUserID sets the "user_id" field.
func (u *FeedbackUpsertBulk) SetUserID(v uint64) *FeedbackUpsertBulk {
	return u.Update(func(s *FeedbackUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *FeedbackUpsertBulk) UpdateUserID() *FeedbackUpsertBulk {
	return u.Update(func(s *FeedbackUpsert) {
		s.UpdateUserID()
	})
}

// Exec executes the query.
func (u *FeedbackUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FeedbackCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FeedbackCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FeedbackUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
