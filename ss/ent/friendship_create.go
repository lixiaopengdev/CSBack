// Code generated by ent, DO NOT EDIT.

package ent

import (
	"CSBackendTmp/ent/friendship"
	"CSBackendTmp/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FriendshipCreate is the builder for creating a Friendship entity.
type FriendshipCreate struct {
	config
	mutation *FriendshipMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (fc *FriendshipCreate) SetCreateTime(t time.Time) *FriendshipCreate {
	fc.mutation.SetCreateTime(t)
	return fc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (fc *FriendshipCreate) SetNillableCreateTime(t *time.Time) *FriendshipCreate {
	if t != nil {
		fc.SetCreateTime(*t)
	}
	return fc
}

// SetUpdateTime sets the "update_time" field.
func (fc *FriendshipCreate) SetUpdateTime(t time.Time) *FriendshipCreate {
	fc.mutation.SetUpdateTime(t)
	return fc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (fc *FriendshipCreate) SetNillableUpdateTime(t *time.Time) *FriendshipCreate {
	if t != nil {
		fc.SetUpdateTime(*t)
	}
	return fc
}

// SetStatus sets the "status" field.
func (fc *FriendshipCreate) SetStatus(f friendship.Status) *FriendshipCreate {
	fc.mutation.SetStatus(f)
	return fc
}

// SetRequestType sets the "request_type" field.
func (fc *FriendshipCreate) SetRequestType(ft friendship.RequestType) *FriendshipCreate {
	fc.mutation.SetRequestType(ft)
	return fc
}

// SetNillableRequestType sets the "request_type" field if the given value is not nil.
func (fc *FriendshipCreate) SetNillableRequestType(ft *friendship.RequestType) *FriendshipCreate {
	if ft != nil {
		fc.SetRequestType(*ft)
	}
	return fc
}

// SetCurrType sets the "curr_type" field.
func (fc *FriendshipCreate) SetCurrType(ft friendship.CurrType) *FriendshipCreate {
	fc.mutation.SetCurrType(ft)
	return fc
}

// SetNillableCurrType sets the "curr_type" field if the given value is not nil.
func (fc *FriendshipCreate) SetNillableCurrType(ft *friendship.CurrType) *FriendshipCreate {
	if ft != nil {
		fc.SetCurrType(*ft)
	}
	return fc
}

// SetUserID sets the "user_id" field.
func (fc *FriendshipCreate) SetUserID(u uint64) *FriendshipCreate {
	fc.mutation.SetUserID(u)
	return fc
}

// SetFriendID sets the "friend_id" field.
func (fc *FriendshipCreate) SetFriendID(u uint64) *FriendshipCreate {
	fc.mutation.SetFriendID(u)
	return fc
}

// SetUser sets the "user" edge to the User entity.
func (fc *FriendshipCreate) SetUser(u *User) *FriendshipCreate {
	return fc.SetUserID(u.ID)
}

// SetFriend sets the "friend" edge to the User entity.
func (fc *FriendshipCreate) SetFriend(u *User) *FriendshipCreate {
	return fc.SetFriendID(u.ID)
}

// Mutation returns the FriendshipMutation object of the builder.
func (fc *FriendshipCreate) Mutation() *FriendshipMutation {
	return fc.mutation
}

// Save creates the Friendship in the database.
func (fc *FriendshipCreate) Save(ctx context.Context) (*Friendship, error) {
	var (
		err  error
		node *Friendship
	)
	fc.defaults()
	if len(fc.hooks) == 0 {
		if err = fc.check(); err != nil {
			return nil, err
		}
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FriendshipMutation)
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
		nv, ok := v.(*Friendship)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from FriendshipMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FriendshipCreate) SaveX(ctx context.Context) *Friendship {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FriendshipCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FriendshipCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FriendshipCreate) defaults() {
	if _, ok := fc.mutation.CreateTime(); !ok {
		v := friendship.DefaultCreateTime()
		fc.mutation.SetCreateTime(v)
	}
	if _, ok := fc.mutation.UpdateTime(); !ok {
		v := friendship.DefaultUpdateTime()
		fc.mutation.SetUpdateTime(v)
	}
	if _, ok := fc.mutation.RequestType(); !ok {
		v := friendship.DefaultRequestType
		fc.mutation.SetRequestType(v)
	}
	if _, ok := fc.mutation.CurrType(); !ok {
		v := friendship.DefaultCurrType
		fc.mutation.SetCurrType(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FriendshipCreate) check() error {
	if _, ok := fc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Friendship.create_time"`)}
	}
	if _, ok := fc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Friendship.update_time"`)}
	}
	if _, ok := fc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Friendship.status"`)}
	}
	if v, ok := fc.mutation.Status(); ok {
		if err := friendship.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Friendship.status": %w`, err)}
		}
	}
	if _, ok := fc.mutation.RequestType(); !ok {
		return &ValidationError{Name: "request_type", err: errors.New(`ent: missing required field "Friendship.request_type"`)}
	}
	if v, ok := fc.mutation.RequestType(); ok {
		if err := friendship.RequestTypeValidator(v); err != nil {
			return &ValidationError{Name: "request_type", err: fmt.Errorf(`ent: validator failed for field "Friendship.request_type": %w`, err)}
		}
	}
	if _, ok := fc.mutation.CurrType(); !ok {
		return &ValidationError{Name: "curr_type", err: errors.New(`ent: missing required field "Friendship.curr_type"`)}
	}
	if v, ok := fc.mutation.CurrType(); ok {
		if err := friendship.CurrTypeValidator(v); err != nil {
			return &ValidationError{Name: "curr_type", err: fmt.Errorf(`ent: validator failed for field "Friendship.curr_type": %w`, err)}
		}
	}
	if _, ok := fc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Friendship.user_id"`)}
	}
	if _, ok := fc.mutation.FriendID(); !ok {
		return &ValidationError{Name: "friend_id", err: errors.New(`ent: missing required field "Friendship.friend_id"`)}
	}
	if _, ok := fc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Friendship.user"`)}
	}
	if _, ok := fc.mutation.FriendID(); !ok {
		return &ValidationError{Name: "friend", err: errors.New(`ent: missing required edge "Friendship.friend"`)}
	}
	return nil
}

func (fc *FriendshipCreate) sqlSave(ctx context.Context) (*Friendship, error) {
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

func (fc *FriendshipCreate) createSpec() (*Friendship, *sqlgraph.CreateSpec) {
	var (
		_node = &Friendship{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: friendship.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: friendship.FieldID,
			},
		}
	)
	_spec.OnConflict = fc.conflict
	if value, ok := fc.mutation.CreateTime(); ok {
		_spec.SetField(friendship.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := fc.mutation.UpdateTime(); ok {
		_spec.SetField(friendship.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := fc.mutation.Status(); ok {
		_spec.SetField(friendship.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := fc.mutation.RequestType(); ok {
		_spec.SetField(friendship.FieldRequestType, field.TypeEnum, value)
		_node.RequestType = value
	}
	if value, ok := fc.mutation.CurrType(); ok {
		_spec.SetField(friendship.FieldCurrType, field.TypeEnum, value)
		_node.CurrType = value
	}
	if nodes := fc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.UserTable,
			Columns: []string{friendship.UserColumn},
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
	if nodes := fc.mutation.FriendIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.FriendTable,
			Columns: []string{friendship.FriendColumn},
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
		_node.FriendID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Friendship.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FriendshipUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (fc *FriendshipCreate) OnConflict(opts ...sql.ConflictOption) *FriendshipUpsertOne {
	fc.conflict = opts
	return &FriendshipUpsertOne{
		create: fc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Friendship.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fc *FriendshipCreate) OnConflictColumns(columns ...string) *FriendshipUpsertOne {
	fc.conflict = append(fc.conflict, sql.ConflictColumns(columns...))
	return &FriendshipUpsertOne{
		create: fc,
	}
}

type (
	// FriendshipUpsertOne is the builder for "upsert"-ing
	//  one Friendship node.
	FriendshipUpsertOne struct {
		create *FriendshipCreate
	}

	// FriendshipUpsert is the "OnConflict" setter.
	FriendshipUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *FriendshipUpsert) SetUpdateTime(v time.Time) *FriendshipUpsert {
	u.Set(friendship.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *FriendshipUpsert) UpdateUpdateTime() *FriendshipUpsert {
	u.SetExcluded(friendship.FieldUpdateTime)
	return u
}

// SetStatus sets the "status" field.
func (u *FriendshipUpsert) SetStatus(v friendship.Status) *FriendshipUpsert {
	u.Set(friendship.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *FriendshipUpsert) UpdateStatus() *FriendshipUpsert {
	u.SetExcluded(friendship.FieldStatus)
	return u
}

// SetRequestType sets the "request_type" field.
func (u *FriendshipUpsert) SetRequestType(v friendship.RequestType) *FriendshipUpsert {
	u.Set(friendship.FieldRequestType, v)
	return u
}

// UpdateRequestType sets the "request_type" field to the value that was provided on create.
func (u *FriendshipUpsert) UpdateRequestType() *FriendshipUpsert {
	u.SetExcluded(friendship.FieldRequestType)
	return u
}

// SetCurrType sets the "curr_type" field.
func (u *FriendshipUpsert) SetCurrType(v friendship.CurrType) *FriendshipUpsert {
	u.Set(friendship.FieldCurrType, v)
	return u
}

// UpdateCurrType sets the "curr_type" field to the value that was provided on create.
func (u *FriendshipUpsert) UpdateCurrType() *FriendshipUpsert {
	u.SetExcluded(friendship.FieldCurrType)
	return u
}

// SetUserID sets the "user_id" field.
func (u *FriendshipUpsert) SetUserID(v uint64) *FriendshipUpsert {
	u.Set(friendship.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *FriendshipUpsert) UpdateUserID() *FriendshipUpsert {
	u.SetExcluded(friendship.FieldUserID)
	return u
}

// SetFriendID sets the "friend_id" field.
func (u *FriendshipUpsert) SetFriendID(v uint64) *FriendshipUpsert {
	u.Set(friendship.FieldFriendID, v)
	return u
}

// UpdateFriendID sets the "friend_id" field to the value that was provided on create.
func (u *FriendshipUpsert) UpdateFriendID() *FriendshipUpsert {
	u.SetExcluded(friendship.FieldFriendID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Friendship.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *FriendshipUpsertOne) UpdateNewValues() *FriendshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(friendship.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Friendship.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FriendshipUpsertOne) Ignore() *FriendshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FriendshipUpsertOne) DoNothing() *FriendshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FriendshipCreate.OnConflict
// documentation for more info.
func (u *FriendshipUpsertOne) Update(set func(*FriendshipUpsert)) *FriendshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FriendshipUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *FriendshipUpsertOne) SetUpdateTime(v time.Time) *FriendshipUpsertOne {
	return u.Update(func(s *FriendshipUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *FriendshipUpsertOne) UpdateUpdateTime() *FriendshipUpsertOne {
	return u.Update(func(s *FriendshipUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetStatus sets the "status" field.
func (u *FriendshipUpsertOne) SetStatus(v friendship.Status) *FriendshipUpsertOne {
	return u.Update(func(s *FriendshipUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *FriendshipUpsertOne) UpdateStatus() *FriendshipUpsertOne {
	return u.Update(func(s *FriendshipUpsert) {
		s.UpdateStatus()
	})
}

// SetRequestType sets the "request_type" field.
func (u *FriendshipUpsertOne) SetRequestType(v friendship.RequestType) *FriendshipUpsertOne {
	return u.Update(func(s *FriendshipUpsert) {
		s.SetRequestType(v)
	})
}

// UpdateRequestType sets the "request_type" field to the value that was provided on create.
func (u *FriendshipUpsertOne) UpdateRequestType() *FriendshipUpsertOne {
	return u.Update(func(s *FriendshipUpsert) {
		s.UpdateRequestType()
	})
}

// SetCurrType sets the "curr_type" field.
func (u *FriendshipUpsertOne) SetCurrType(v friendship.CurrType) *FriendshipUpsertOne {
	return u.Update(func(s *FriendshipUpsert) {
		s.SetCurrType(v)
	})
}

// UpdateCurrType sets the "curr_type" field to the value that was provided on create.
func (u *FriendshipUpsertOne) UpdateCurrType() *FriendshipUpsertOne {
	return u.Update(func(s *FriendshipUpsert) {
		s.UpdateCurrType()
	})
}

// SetUserID sets the "user_id" field.
func (u *FriendshipUpsertOne) SetUserID(v uint64) *FriendshipUpsertOne {
	return u.Update(func(s *FriendshipUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *FriendshipUpsertOne) UpdateUserID() *FriendshipUpsertOne {
	return u.Update(func(s *FriendshipUpsert) {
		s.UpdateUserID()
	})
}

// SetFriendID sets the "friend_id" field.
func (u *FriendshipUpsertOne) SetFriendID(v uint64) *FriendshipUpsertOne {
	return u.Update(func(s *FriendshipUpsert) {
		s.SetFriendID(v)
	})
}

// UpdateFriendID sets the "friend_id" field to the value that was provided on create.
func (u *FriendshipUpsertOne) UpdateFriendID() *FriendshipUpsertOne {
	return u.Update(func(s *FriendshipUpsert) {
		s.UpdateFriendID()
	})
}

// Exec executes the query.
func (u *FriendshipUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FriendshipCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FriendshipUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FriendshipUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FriendshipUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FriendshipCreateBulk is the builder for creating many Friendship entities in bulk.
type FriendshipCreateBulk struct {
	config
	builders []*FriendshipCreate
	conflict []sql.ConflictOption
}

// Save creates the Friendship entities in the database.
func (fcb *FriendshipCreateBulk) Save(ctx context.Context) ([]*Friendship, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Friendship, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FriendshipMutation)
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
func (fcb *FriendshipCreateBulk) SaveX(ctx context.Context) []*Friendship {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FriendshipCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FriendshipCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Friendship.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FriendshipUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (fcb *FriendshipCreateBulk) OnConflict(opts ...sql.ConflictOption) *FriendshipUpsertBulk {
	fcb.conflict = opts
	return &FriendshipUpsertBulk{
		create: fcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Friendship.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fcb *FriendshipCreateBulk) OnConflictColumns(columns ...string) *FriendshipUpsertBulk {
	fcb.conflict = append(fcb.conflict, sql.ConflictColumns(columns...))
	return &FriendshipUpsertBulk{
		create: fcb,
	}
}

// FriendshipUpsertBulk is the builder for "upsert"-ing
// a bulk of Friendship nodes.
type FriendshipUpsertBulk struct {
	create *FriendshipCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Friendship.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *FriendshipUpsertBulk) UpdateNewValues() *FriendshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(friendship.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Friendship.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FriendshipUpsertBulk) Ignore() *FriendshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FriendshipUpsertBulk) DoNothing() *FriendshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FriendshipCreateBulk.OnConflict
// documentation for more info.
func (u *FriendshipUpsertBulk) Update(set func(*FriendshipUpsert)) *FriendshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FriendshipUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *FriendshipUpsertBulk) SetUpdateTime(v time.Time) *FriendshipUpsertBulk {
	return u.Update(func(s *FriendshipUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *FriendshipUpsertBulk) UpdateUpdateTime() *FriendshipUpsertBulk {
	return u.Update(func(s *FriendshipUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetStatus sets the "status" field.
func (u *FriendshipUpsertBulk) SetStatus(v friendship.Status) *FriendshipUpsertBulk {
	return u.Update(func(s *FriendshipUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *FriendshipUpsertBulk) UpdateStatus() *FriendshipUpsertBulk {
	return u.Update(func(s *FriendshipUpsert) {
		s.UpdateStatus()
	})
}

// SetRequestType sets the "request_type" field.
func (u *FriendshipUpsertBulk) SetRequestType(v friendship.RequestType) *FriendshipUpsertBulk {
	return u.Update(func(s *FriendshipUpsert) {
		s.SetRequestType(v)
	})
}

// UpdateRequestType sets the "request_type" field to the value that was provided on create.
func (u *FriendshipUpsertBulk) UpdateRequestType() *FriendshipUpsertBulk {
	return u.Update(func(s *FriendshipUpsert) {
		s.UpdateRequestType()
	})
}

// SetCurrType sets the "curr_type" field.
func (u *FriendshipUpsertBulk) SetCurrType(v friendship.CurrType) *FriendshipUpsertBulk {
	return u.Update(func(s *FriendshipUpsert) {
		s.SetCurrType(v)
	})
}

// UpdateCurrType sets the "curr_type" field to the value that was provided on create.
func (u *FriendshipUpsertBulk) UpdateCurrType() *FriendshipUpsertBulk {
	return u.Update(func(s *FriendshipUpsert) {
		s.UpdateCurrType()
	})
}

// SetUserID sets the "user_id" field.
func (u *FriendshipUpsertBulk) SetUserID(v uint64) *FriendshipUpsertBulk {
	return u.Update(func(s *FriendshipUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *FriendshipUpsertBulk) UpdateUserID() *FriendshipUpsertBulk {
	return u.Update(func(s *FriendshipUpsert) {
		s.UpdateUserID()
	})
}

// SetFriendID sets the "friend_id" field.
func (u *FriendshipUpsertBulk) SetFriendID(v uint64) *FriendshipUpsertBulk {
	return u.Update(func(s *FriendshipUpsert) {
		s.SetFriendID(v)
	})
}

// UpdateFriendID sets the "friend_id" field to the value that was provided on create.
func (u *FriendshipUpsertBulk) UpdateFriendID() *FriendshipUpsertBulk {
	return u.Update(func(s *FriendshipUpsert) {
		s.UpdateFriendID()
	})
}

// Exec executes the query.
func (u *FriendshipUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FriendshipCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FriendshipCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FriendshipUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
