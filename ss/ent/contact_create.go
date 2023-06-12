// Code generated by ent, DO NOT EDIT.

package ent

import (
	"CSBackendTmp/ent/contact"
	"CSBackendTmp/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ContactCreate is the builder for creating a Contact entity.
type ContactCreate struct {
	config
	mutation *ContactMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (cc *ContactCreate) SetCreateTime(t time.Time) *ContactCreate {
	cc.mutation.SetCreateTime(t)
	return cc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (cc *ContactCreate) SetNillableCreateTime(t *time.Time) *ContactCreate {
	if t != nil {
		cc.SetCreateTime(*t)
	}
	return cc
}

// SetUpdateTime sets the "update_time" field.
func (cc *ContactCreate) SetUpdateTime(t time.Time) *ContactCreate {
	cc.mutation.SetUpdateTime(t)
	return cc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (cc *ContactCreate) SetNillableUpdateTime(t *time.Time) *ContactCreate {
	if t != nil {
		cc.SetUpdateTime(*t)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *ContactCreate) SetName(s string) *ContactCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cc *ContactCreate) SetNillableName(s *string) *ContactCreate {
	if s != nil {
		cc.SetName(*s)
	}
	return cc
}

// SetMobileNo sets the "mobile_no" field.
func (cc *ContactCreate) SetMobileNo(s string) *ContactCreate {
	cc.mutation.SetMobileNo(s)
	return cc
}

// SetNillableMobileNo sets the "mobile_no" field if the given value is not nil.
func (cc *ContactCreate) SetNillableMobileNo(s *string) *ContactCreate {
	if s != nil {
		cc.SetMobileNo(*s)
	}
	return cc
}

// SetEmail sets the "email" field.
func (cc *ContactCreate) SetEmail(s string) *ContactCreate {
	cc.mutation.SetEmail(s)
	return cc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (cc *ContactCreate) SetNillableEmail(s *string) *ContactCreate {
	if s != nil {
		cc.SetEmail(*s)
	}
	return cc
}

// SetUserID sets the "user_id" field.
func (cc *ContactCreate) SetUserID(u uint64) *ContactCreate {
	cc.mutation.SetUserID(u)
	return cc
}

// SetID sets the "id" field.
func (cc *ContactCreate) SetID(u uint64) *ContactCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cc *ContactCreate) SetOwnerID(id uint64) *ContactCreate {
	cc.mutation.SetOwnerID(id)
	return cc
}

// SetOwner sets the "owner" edge to the User entity.
func (cc *ContactCreate) SetOwner(u *User) *ContactCreate {
	return cc.SetOwnerID(u.ID)
}

// Mutation returns the ContactMutation object of the builder.
func (cc *ContactCreate) Mutation() *ContactMutation {
	return cc.mutation
}

// Save creates the Contact in the database.
func (cc *ContactCreate) Save(ctx context.Context) (*Contact, error) {
	var (
		err  error
		node *Contact
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContactMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Contact)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ContactMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ContactCreate) SaveX(ctx context.Context) *Contact {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ContactCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ContactCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ContactCreate) defaults() {
	if _, ok := cc.mutation.CreateTime(); !ok {
		v := contact.DefaultCreateTime()
		cc.mutation.SetCreateTime(v)
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		v := contact.DefaultUpdateTime()
		cc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ContactCreate) check() error {
	if _, ok := cc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Contact.create_time"`)}
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Contact.update_time"`)}
	}
	if _, ok := cc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Contact.user_id"`)}
	}
	if _, ok := cc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Contact.owner"`)}
	}
	return nil
}

func (cc *ContactCreate) sqlSave(ctx context.Context) (*Contact, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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

func (cc *ContactCreate) createSpec() (*Contact, *sqlgraph.CreateSpec) {
	var (
		_node = &Contact{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: contact.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: contact.FieldID,
			},
		}
	)
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreateTime(); ok {
		_spec.SetField(contact.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := cc.mutation.UpdateTime(); ok {
		_spec.SetField(contact.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(contact.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.MobileNo(); ok {
		_spec.SetField(contact.FieldMobileNo, field.TypeString, value)
		_node.MobileNo = value
	}
	if value, ok := cc.mutation.Email(); ok {
		_spec.SetField(contact.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if nodes := cc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contact.OwnerTable,
			Columns: []string{contact.OwnerColumn},
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
//	client.Contact.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ContactUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (cc *ContactCreate) OnConflict(opts ...sql.ConflictOption) *ContactUpsertOne {
	cc.conflict = opts
	return &ContactUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Contact.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cc *ContactCreate) OnConflictColumns(columns ...string) *ContactUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &ContactUpsertOne{
		create: cc,
	}
}

type (
	// ContactUpsertOne is the builder for "upsert"-ing
	//  one Contact node.
	ContactUpsertOne struct {
		create *ContactCreate
	}

	// ContactUpsert is the "OnConflict" setter.
	ContactUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *ContactUpsert) SetUpdateTime(v time.Time) *ContactUpsert {
	u.Set(contact.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ContactUpsert) UpdateUpdateTime() *ContactUpsert {
	u.SetExcluded(contact.FieldUpdateTime)
	return u
}

// SetName sets the "name" field.
func (u *ContactUpsert) SetName(v string) *ContactUpsert {
	u.Set(contact.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ContactUpsert) UpdateName() *ContactUpsert {
	u.SetExcluded(contact.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *ContactUpsert) ClearName() *ContactUpsert {
	u.SetNull(contact.FieldName)
	return u
}

// SetMobileNo sets the "mobile_no" field.
func (u *ContactUpsert) SetMobileNo(v string) *ContactUpsert {
	u.Set(contact.FieldMobileNo, v)
	return u
}

// UpdateMobileNo sets the "mobile_no" field to the value that was provided on create.
func (u *ContactUpsert) UpdateMobileNo() *ContactUpsert {
	u.SetExcluded(contact.FieldMobileNo)
	return u
}

// ClearMobileNo clears the value of the "mobile_no" field.
func (u *ContactUpsert) ClearMobileNo() *ContactUpsert {
	u.SetNull(contact.FieldMobileNo)
	return u
}

// SetEmail sets the "email" field.
func (u *ContactUpsert) SetEmail(v string) *ContactUpsert {
	u.Set(contact.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *ContactUpsert) UpdateEmail() *ContactUpsert {
	u.SetExcluded(contact.FieldEmail)
	return u
}

// ClearEmail clears the value of the "email" field.
func (u *ContactUpsert) ClearEmail() *ContactUpsert {
	u.SetNull(contact.FieldEmail)
	return u
}

// SetUserID sets the "user_id" field.
func (u *ContactUpsert) SetUserID(v uint64) *ContactUpsert {
	u.Set(contact.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ContactUpsert) UpdateUserID() *ContactUpsert {
	u.SetExcluded(contact.FieldUserID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Contact.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(contact.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ContactUpsertOne) UpdateNewValues() *ContactUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(contact.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(contact.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Contact.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ContactUpsertOne) Ignore() *ContactUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ContactUpsertOne) DoNothing() *ContactUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ContactCreate.OnConflict
// documentation for more info.
func (u *ContactUpsertOne) Update(set func(*ContactUpsert)) *ContactUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ContactUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *ContactUpsertOne) SetUpdateTime(v time.Time) *ContactUpsertOne {
	return u.Update(func(s *ContactUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ContactUpsertOne) UpdateUpdateTime() *ContactUpsertOne {
	return u.Update(func(s *ContactUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *ContactUpsertOne) SetName(v string) *ContactUpsertOne {
	return u.Update(func(s *ContactUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ContactUpsertOne) UpdateName() *ContactUpsertOne {
	return u.Update(func(s *ContactUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *ContactUpsertOne) ClearName() *ContactUpsertOne {
	return u.Update(func(s *ContactUpsert) {
		s.ClearName()
	})
}

// SetMobileNo sets the "mobile_no" field.
func (u *ContactUpsertOne) SetMobileNo(v string) *ContactUpsertOne {
	return u.Update(func(s *ContactUpsert) {
		s.SetMobileNo(v)
	})
}

// UpdateMobileNo sets the "mobile_no" field to the value that was provided on create.
func (u *ContactUpsertOne) UpdateMobileNo() *ContactUpsertOne {
	return u.Update(func(s *ContactUpsert) {
		s.UpdateMobileNo()
	})
}

// ClearMobileNo clears the value of the "mobile_no" field.
func (u *ContactUpsertOne) ClearMobileNo() *ContactUpsertOne {
	return u.Update(func(s *ContactUpsert) {
		s.ClearMobileNo()
	})
}

// SetEmail sets the "email" field.
func (u *ContactUpsertOne) SetEmail(v string) *ContactUpsertOne {
	return u.Update(func(s *ContactUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *ContactUpsertOne) UpdateEmail() *ContactUpsertOne {
	return u.Update(func(s *ContactUpsert) {
		s.UpdateEmail()
	})
}

// ClearEmail clears the value of the "email" field.
func (u *ContactUpsertOne) ClearEmail() *ContactUpsertOne {
	return u.Update(func(s *ContactUpsert) {
		s.ClearEmail()
	})
}

// SetUserID sets the "user_id" field.
func (u *ContactUpsertOne) SetUserID(v uint64) *ContactUpsertOne {
	return u.Update(func(s *ContactUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ContactUpsertOne) UpdateUserID() *ContactUpsertOne {
	return u.Update(func(s *ContactUpsert) {
		s.UpdateUserID()
	})
}

// Exec executes the query.
func (u *ContactUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ContactCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ContactUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ContactUpsertOne) ID(ctx context.Context) (id uint64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ContactUpsertOne) IDX(ctx context.Context) uint64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ContactCreateBulk is the builder for creating many Contact entities in bulk.
type ContactCreateBulk struct {
	config
	builders []*ContactCreate
	conflict []sql.ConflictOption
}

// Save creates the Contact entities in the database.
func (ccb *ContactCreateBulk) Save(ctx context.Context) ([]*Contact, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Contact, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ContactMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ContactCreateBulk) SaveX(ctx context.Context) []*Contact {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ContactCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ContactCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Contact.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ContactUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (ccb *ContactCreateBulk) OnConflict(opts ...sql.ConflictOption) *ContactUpsertBulk {
	ccb.conflict = opts
	return &ContactUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Contact.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ccb *ContactCreateBulk) OnConflictColumns(columns ...string) *ContactUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &ContactUpsertBulk{
		create: ccb,
	}
}

// ContactUpsertBulk is the builder for "upsert"-ing
// a bulk of Contact nodes.
type ContactUpsertBulk struct {
	create *ContactCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Contact.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(contact.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ContactUpsertBulk) UpdateNewValues() *ContactUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(contact.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(contact.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Contact.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ContactUpsertBulk) Ignore() *ContactUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ContactUpsertBulk) DoNothing() *ContactUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ContactCreateBulk.OnConflict
// documentation for more info.
func (u *ContactUpsertBulk) Update(set func(*ContactUpsert)) *ContactUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ContactUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *ContactUpsertBulk) SetUpdateTime(v time.Time) *ContactUpsertBulk {
	return u.Update(func(s *ContactUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ContactUpsertBulk) UpdateUpdateTime() *ContactUpsertBulk {
	return u.Update(func(s *ContactUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *ContactUpsertBulk) SetName(v string) *ContactUpsertBulk {
	return u.Update(func(s *ContactUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ContactUpsertBulk) UpdateName() *ContactUpsertBulk {
	return u.Update(func(s *ContactUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *ContactUpsertBulk) ClearName() *ContactUpsertBulk {
	return u.Update(func(s *ContactUpsert) {
		s.ClearName()
	})
}

// SetMobileNo sets the "mobile_no" field.
func (u *ContactUpsertBulk) SetMobileNo(v string) *ContactUpsertBulk {
	return u.Update(func(s *ContactUpsert) {
		s.SetMobileNo(v)
	})
}

// UpdateMobileNo sets the "mobile_no" field to the value that was provided on create.
func (u *ContactUpsertBulk) UpdateMobileNo() *ContactUpsertBulk {
	return u.Update(func(s *ContactUpsert) {
		s.UpdateMobileNo()
	})
}

// ClearMobileNo clears the value of the "mobile_no" field.
func (u *ContactUpsertBulk) ClearMobileNo() *ContactUpsertBulk {
	return u.Update(func(s *ContactUpsert) {
		s.ClearMobileNo()
	})
}

// SetEmail sets the "email" field.
func (u *ContactUpsertBulk) SetEmail(v string) *ContactUpsertBulk {
	return u.Update(func(s *ContactUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *ContactUpsertBulk) UpdateEmail() *ContactUpsertBulk {
	return u.Update(func(s *ContactUpsert) {
		s.UpdateEmail()
	})
}

// ClearEmail clears the value of the "email" field.
func (u *ContactUpsertBulk) ClearEmail() *ContactUpsertBulk {
	return u.Update(func(s *ContactUpsert) {
		s.ClearEmail()
	})
}

// SetUserID sets the "user_id" field.
func (u *ContactUpsertBulk) SetUserID(v uint64) *ContactUpsertBulk {
	return u.Update(func(s *ContactUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ContactUpsertBulk) UpdateUserID() *ContactUpsertBulk {
	return u.Update(func(s *ContactUpsert) {
		s.UpdateUserID()
	})
}

// Exec executes the query.
func (u *ContactUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ContactCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ContactCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ContactUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}