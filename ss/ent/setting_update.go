// Code generated by ent, DO NOT EDIT.

package ent

import (
	"CSBackendTmp/ent/predicate"
	"CSBackendTmp/ent/setting"
	"CSBackendTmp/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SettingUpdate is the builder for updating Setting entities.
type SettingUpdate struct {
	config
	hooks    []Hook
	mutation *SettingMutation
}

// Where appends a list predicates to the SettingUpdate builder.
func (su *SettingUpdate) Where(ps ...predicate.Setting) *SettingUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdateTime sets the "update_time" field.
func (su *SettingUpdate) SetUpdateTime(t time.Time) *SettingUpdate {
	su.mutation.SetUpdateTime(t)
	return su
}

// SetFriendsOnline sets the "friends_online" field.
func (su *SettingUpdate) SetFriendsOnline(b bool) *SettingUpdate {
	su.mutation.SetFriendsOnline(b)
	return su
}

// SetNillableFriendsOnline sets the "friends_online" field if the given value is not nil.
func (su *SettingUpdate) SetNillableFriendsOnline(b *bool) *SettingUpdate {
	if b != nil {
		su.SetFriendsOnline(*b)
	}
	return su
}

// ClearFriendsOnline clears the value of the "friends_online" field.
func (su *SettingUpdate) ClearFriendsOnline() *SettingUpdate {
	su.mutation.ClearFriendsOnline()
	return su
}

// SetTimeDewFromFriends sets the "time_dew_from_friends" field.
func (su *SettingUpdate) SetTimeDewFromFriends(b bool) *SettingUpdate {
	su.mutation.SetTimeDewFromFriends(b)
	return su
}

// SetNillableTimeDewFromFriends sets the "time_dew_from_friends" field if the given value is not nil.
func (su *SettingUpdate) SetNillableTimeDewFromFriends(b *bool) *SettingUpdate {
	if b != nil {
		su.SetTimeDewFromFriends(*b)
	}
	return su
}

// ClearTimeDewFromFriends clears the value of the "time_dew_from_friends" field.
func (su *SettingUpdate) ClearTimeDewFromFriends() *SettingUpdate {
	su.mutation.ClearTimeDewFromFriends()
	return su
}

// SetDetailedNotification sets the "detailed_notification" field.
func (su *SettingUpdate) SetDetailedNotification(b bool) *SettingUpdate {
	su.mutation.SetDetailedNotification(b)
	return su
}

// SetNillableDetailedNotification sets the "detailed_notification" field if the given value is not nil.
func (su *SettingUpdate) SetNillableDetailedNotification(b *bool) *SettingUpdate {
	if b != nil {
		su.SetDetailedNotification(*b)
	}
	return su
}

// ClearDetailedNotification clears the value of the "detailed_notification" field.
func (su *SettingUpdate) ClearDetailedNotification() *SettingUpdate {
	su.mutation.ClearDetailedNotification()
	return su
}

// SetReceiveFieldInvitation sets the "receive_field_invitation" field.
func (su *SettingUpdate) SetReceiveFieldInvitation(b bool) *SettingUpdate {
	su.mutation.SetReceiveFieldInvitation(b)
	return su
}

// SetNillableReceiveFieldInvitation sets the "receive_field_invitation" field if the given value is not nil.
func (su *SettingUpdate) SetNillableReceiveFieldInvitation(b *bool) *SettingUpdate {
	if b != nil {
		su.SetReceiveFieldInvitation(*b)
	}
	return su
}

// ClearReceiveFieldInvitation clears the value of the "receive_field_invitation" field.
func (su *SettingUpdate) ClearReceiveFieldInvitation() *SettingUpdate {
	su.mutation.ClearReceiveFieldInvitation()
	return su
}

// SetSeeMyLocation sets the "see_my_location" field.
func (su *SettingUpdate) SetSeeMyLocation(b bool) *SettingUpdate {
	su.mutation.SetSeeMyLocation(b)
	return su
}

// SetNillableSeeMyLocation sets the "see_my_location" field if the given value is not nil.
func (su *SettingUpdate) SetNillableSeeMyLocation(b *bool) *SettingUpdate {
	if b != nil {
		su.SetSeeMyLocation(*b)
	}
	return su
}

// ClearSeeMyLocation clears the value of the "see_my_location" field.
func (su *SettingUpdate) ClearSeeMyLocation() *SettingUpdate {
	su.mutation.ClearSeeMyLocation()
	return su
}

// SetCamera sets the "camera" field.
func (su *SettingUpdate) SetCamera(b bool) *SettingUpdate {
	su.mutation.SetCamera(b)
	return su
}

// SetNillableCamera sets the "camera" field if the given value is not nil.
func (su *SettingUpdate) SetNillableCamera(b *bool) *SettingUpdate {
	if b != nil {
		su.SetCamera(*b)
	}
	return su
}

// ClearCamera clears the value of the "camera" field.
func (su *SettingUpdate) ClearCamera() *SettingUpdate {
	su.mutation.ClearCamera()
	return su
}

// SetMicrophone sets the "microphone" field.
func (su *SettingUpdate) SetMicrophone(b bool) *SettingUpdate {
	su.mutation.SetMicrophone(b)
	return su
}

// SetNillableMicrophone sets the "microphone" field if the given value is not nil.
func (su *SettingUpdate) SetNillableMicrophone(b *bool) *SettingUpdate {
	if b != nil {
		su.SetMicrophone(*b)
	}
	return su
}

// ClearMicrophone clears the value of the "microphone" field.
func (su *SettingUpdate) ClearMicrophone() *SettingUpdate {
	su.mutation.ClearMicrophone()
	return su
}

// SetHealthData sets the "health_data" field.
func (su *SettingUpdate) SetHealthData(b bool) *SettingUpdate {
	su.mutation.SetHealthData(b)
	return su
}

// SetNillableHealthData sets the "health_data" field if the given value is not nil.
func (su *SettingUpdate) SetNillableHealthData(b *bool) *SettingUpdate {
	if b != nil {
		su.SetHealthData(*b)
	}
	return su
}

// ClearHealthData clears the value of the "health_data" field.
func (su *SettingUpdate) ClearHealthData() *SettingUpdate {
	su.mutation.ClearHealthData()
	return su
}

// SetTimeDewLocation sets the "time_dew_location" field.
func (su *SettingUpdate) SetTimeDewLocation(b bool) *SettingUpdate {
	su.mutation.SetTimeDewLocation(b)
	return su
}

// SetNillableTimeDewLocation sets the "time_dew_location" field if the given value is not nil.
func (su *SettingUpdate) SetNillableTimeDewLocation(b *bool) *SettingUpdate {
	if b != nil {
		su.SetTimeDewLocation(*b)
	}
	return su
}

// ClearTimeDewLocation clears the value of the "time_dew_location" field.
func (su *SettingUpdate) ClearTimeDewLocation() *SettingUpdate {
	su.mutation.ClearTimeDewLocation()
	return su
}

// SetTimeDewMicrophone sets the "time_dew_microphone" field.
func (su *SettingUpdate) SetTimeDewMicrophone(b bool) *SettingUpdate {
	su.mutation.SetTimeDewMicrophone(b)
	return su
}

// SetNillableTimeDewMicrophone sets the "time_dew_microphone" field if the given value is not nil.
func (su *SettingUpdate) SetNillableTimeDewMicrophone(b *bool) *SettingUpdate {
	if b != nil {
		su.SetTimeDewMicrophone(*b)
	}
	return su
}

// ClearTimeDewMicrophone clears the value of the "time_dew_microphone" field.
func (su *SettingUpdate) ClearTimeDewMicrophone() *SettingUpdate {
	su.mutation.ClearTimeDewMicrophone()
	return su
}

// SetTimeDewLora sets the "time_dew_Lora" field.
func (su *SettingUpdate) SetTimeDewLora(b bool) *SettingUpdate {
	su.mutation.SetTimeDewLora(b)
	return su
}

// SetNillableTimeDewLora sets the "time_dew_Lora" field if the given value is not nil.
func (su *SettingUpdate) SetNillableTimeDewLora(b *bool) *SettingUpdate {
	if b != nil {
		su.SetTimeDewLora(*b)
	}
	return su
}

// ClearTimeDewLora clears the value of the "time_dew_Lora" field.
func (su *SettingUpdate) ClearTimeDewLora() *SettingUpdate {
	su.mutation.ClearTimeDewLora()
	return su
}

// SetPublicCollection sets the "public_collection" field.
func (su *SettingUpdate) SetPublicCollection(b bool) *SettingUpdate {
	su.mutation.SetPublicCollection(b)
	return su
}

// SetNillablePublicCollection sets the "public_collection" field if the given value is not nil.
func (su *SettingUpdate) SetNillablePublicCollection(b *bool) *SettingUpdate {
	if b != nil {
		su.SetPublicCollection(*b)
	}
	return su
}

// ClearPublicCollection clears the value of the "public_collection" field.
func (su *SettingUpdate) ClearPublicCollection() *SettingUpdate {
	su.mutation.ClearPublicCollection()
	return su
}

// SetUserID sets the "user_id" field.
func (su *SettingUpdate) SetUserID(u uint64) *SettingUpdate {
	su.mutation.SetUserID(u)
	return su
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (su *SettingUpdate) SetOwnerID(id uint64) *SettingUpdate {
	su.mutation.SetOwnerID(id)
	return su
}

// SetOwner sets the "owner" edge to the User entity.
func (su *SettingUpdate) SetOwner(u *User) *SettingUpdate {
	return su.SetOwnerID(u.ID)
}

// Mutation returns the SettingMutation object of the builder.
func (su *SettingUpdate) Mutation() *SettingMutation {
	return su.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (su *SettingUpdate) ClearOwner() *SettingUpdate {
	su.mutation.ClearOwner()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SettingUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	su.defaults()
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SettingUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SettingUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SettingUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SettingUpdate) defaults() {
	if _, ok := su.mutation.UpdateTime(); !ok {
		v := setting.UpdateDefaultUpdateTime()
		su.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SettingUpdate) check() error {
	if _, ok := su.mutation.OwnerID(); su.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Setting.owner"`)
	}
	return nil
}

func (su *SettingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   setting.Table,
			Columns: setting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: setting.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdateTime(); ok {
		_spec.SetField(setting.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := su.mutation.FriendsOnline(); ok {
		_spec.SetField(setting.FieldFriendsOnline, field.TypeBool, value)
	}
	if su.mutation.FriendsOnlineCleared() {
		_spec.ClearField(setting.FieldFriendsOnline, field.TypeBool)
	}
	if value, ok := su.mutation.TimeDewFromFriends(); ok {
		_spec.SetField(setting.FieldTimeDewFromFriends, field.TypeBool, value)
	}
	if su.mutation.TimeDewFromFriendsCleared() {
		_spec.ClearField(setting.FieldTimeDewFromFriends, field.TypeBool)
	}
	if value, ok := su.mutation.DetailedNotification(); ok {
		_spec.SetField(setting.FieldDetailedNotification, field.TypeBool, value)
	}
	if su.mutation.DetailedNotificationCleared() {
		_spec.ClearField(setting.FieldDetailedNotification, field.TypeBool)
	}
	if value, ok := su.mutation.ReceiveFieldInvitation(); ok {
		_spec.SetField(setting.FieldReceiveFieldInvitation, field.TypeBool, value)
	}
	if su.mutation.ReceiveFieldInvitationCleared() {
		_spec.ClearField(setting.FieldReceiveFieldInvitation, field.TypeBool)
	}
	if value, ok := su.mutation.SeeMyLocation(); ok {
		_spec.SetField(setting.FieldSeeMyLocation, field.TypeBool, value)
	}
	if su.mutation.SeeMyLocationCleared() {
		_spec.ClearField(setting.FieldSeeMyLocation, field.TypeBool)
	}
	if value, ok := su.mutation.Camera(); ok {
		_spec.SetField(setting.FieldCamera, field.TypeBool, value)
	}
	if su.mutation.CameraCleared() {
		_spec.ClearField(setting.FieldCamera, field.TypeBool)
	}
	if value, ok := su.mutation.Microphone(); ok {
		_spec.SetField(setting.FieldMicrophone, field.TypeBool, value)
	}
	if su.mutation.MicrophoneCleared() {
		_spec.ClearField(setting.FieldMicrophone, field.TypeBool)
	}
	if value, ok := su.mutation.HealthData(); ok {
		_spec.SetField(setting.FieldHealthData, field.TypeBool, value)
	}
	if su.mutation.HealthDataCleared() {
		_spec.ClearField(setting.FieldHealthData, field.TypeBool)
	}
	if value, ok := su.mutation.TimeDewLocation(); ok {
		_spec.SetField(setting.FieldTimeDewLocation, field.TypeBool, value)
	}
	if su.mutation.TimeDewLocationCleared() {
		_spec.ClearField(setting.FieldTimeDewLocation, field.TypeBool)
	}
	if value, ok := su.mutation.TimeDewMicrophone(); ok {
		_spec.SetField(setting.FieldTimeDewMicrophone, field.TypeBool, value)
	}
	if su.mutation.TimeDewMicrophoneCleared() {
		_spec.ClearField(setting.FieldTimeDewMicrophone, field.TypeBool)
	}
	if value, ok := su.mutation.TimeDewLora(); ok {
		_spec.SetField(setting.FieldTimeDewLora, field.TypeBool, value)
	}
	if su.mutation.TimeDewLoraCleared() {
		_spec.ClearField(setting.FieldTimeDewLora, field.TypeBool)
	}
	if value, ok := su.mutation.PublicCollection(); ok {
		_spec.SetField(setting.FieldPublicCollection, field.TypeBool, value)
	}
	if su.mutation.PublicCollectionCleared() {
		_spec.ClearField(setting.FieldPublicCollection, field.TypeBool)
	}
	if su.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   setting.OwnerTable,
			Columns: []string{setting.OwnerColumn},
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
	if nodes := su.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   setting.OwnerTable,
			Columns: []string{setting.OwnerColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{setting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SettingUpdateOne is the builder for updating a single Setting entity.
type SettingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SettingMutation
}

// SetUpdateTime sets the "update_time" field.
func (suo *SettingUpdateOne) SetUpdateTime(t time.Time) *SettingUpdateOne {
	suo.mutation.SetUpdateTime(t)
	return suo
}

// SetFriendsOnline sets the "friends_online" field.
func (suo *SettingUpdateOne) SetFriendsOnline(b bool) *SettingUpdateOne {
	suo.mutation.SetFriendsOnline(b)
	return suo
}

// SetNillableFriendsOnline sets the "friends_online" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableFriendsOnline(b *bool) *SettingUpdateOne {
	if b != nil {
		suo.SetFriendsOnline(*b)
	}
	return suo
}

// ClearFriendsOnline clears the value of the "friends_online" field.
func (suo *SettingUpdateOne) ClearFriendsOnline() *SettingUpdateOne {
	suo.mutation.ClearFriendsOnline()
	return suo
}

// SetTimeDewFromFriends sets the "time_dew_from_friends" field.
func (suo *SettingUpdateOne) SetTimeDewFromFriends(b bool) *SettingUpdateOne {
	suo.mutation.SetTimeDewFromFriends(b)
	return suo
}

// SetNillableTimeDewFromFriends sets the "time_dew_from_friends" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableTimeDewFromFriends(b *bool) *SettingUpdateOne {
	if b != nil {
		suo.SetTimeDewFromFriends(*b)
	}
	return suo
}

// ClearTimeDewFromFriends clears the value of the "time_dew_from_friends" field.
func (suo *SettingUpdateOne) ClearTimeDewFromFriends() *SettingUpdateOne {
	suo.mutation.ClearTimeDewFromFriends()
	return suo
}

// SetDetailedNotification sets the "detailed_notification" field.
func (suo *SettingUpdateOne) SetDetailedNotification(b bool) *SettingUpdateOne {
	suo.mutation.SetDetailedNotification(b)
	return suo
}

// SetNillableDetailedNotification sets the "detailed_notification" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableDetailedNotification(b *bool) *SettingUpdateOne {
	if b != nil {
		suo.SetDetailedNotification(*b)
	}
	return suo
}

// ClearDetailedNotification clears the value of the "detailed_notification" field.
func (suo *SettingUpdateOne) ClearDetailedNotification() *SettingUpdateOne {
	suo.mutation.ClearDetailedNotification()
	return suo
}

// SetReceiveFieldInvitation sets the "receive_field_invitation" field.
func (suo *SettingUpdateOne) SetReceiveFieldInvitation(b bool) *SettingUpdateOne {
	suo.mutation.SetReceiveFieldInvitation(b)
	return suo
}

// SetNillableReceiveFieldInvitation sets the "receive_field_invitation" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableReceiveFieldInvitation(b *bool) *SettingUpdateOne {
	if b != nil {
		suo.SetReceiveFieldInvitation(*b)
	}
	return suo
}

// ClearReceiveFieldInvitation clears the value of the "receive_field_invitation" field.
func (suo *SettingUpdateOne) ClearReceiveFieldInvitation() *SettingUpdateOne {
	suo.mutation.ClearReceiveFieldInvitation()
	return suo
}

// SetSeeMyLocation sets the "see_my_location" field.
func (suo *SettingUpdateOne) SetSeeMyLocation(b bool) *SettingUpdateOne {
	suo.mutation.SetSeeMyLocation(b)
	return suo
}

// SetNillableSeeMyLocation sets the "see_my_location" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableSeeMyLocation(b *bool) *SettingUpdateOne {
	if b != nil {
		suo.SetSeeMyLocation(*b)
	}
	return suo
}

// ClearSeeMyLocation clears the value of the "see_my_location" field.
func (suo *SettingUpdateOne) ClearSeeMyLocation() *SettingUpdateOne {
	suo.mutation.ClearSeeMyLocation()
	return suo
}

// SetCamera sets the "camera" field.
func (suo *SettingUpdateOne) SetCamera(b bool) *SettingUpdateOne {
	suo.mutation.SetCamera(b)
	return suo
}

// SetNillableCamera sets the "camera" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableCamera(b *bool) *SettingUpdateOne {
	if b != nil {
		suo.SetCamera(*b)
	}
	return suo
}

// ClearCamera clears the value of the "camera" field.
func (suo *SettingUpdateOne) ClearCamera() *SettingUpdateOne {
	suo.mutation.ClearCamera()
	return suo
}

// SetMicrophone sets the "microphone" field.
func (suo *SettingUpdateOne) SetMicrophone(b bool) *SettingUpdateOne {
	suo.mutation.SetMicrophone(b)
	return suo
}

// SetNillableMicrophone sets the "microphone" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableMicrophone(b *bool) *SettingUpdateOne {
	if b != nil {
		suo.SetMicrophone(*b)
	}
	return suo
}

// ClearMicrophone clears the value of the "microphone" field.
func (suo *SettingUpdateOne) ClearMicrophone() *SettingUpdateOne {
	suo.mutation.ClearMicrophone()
	return suo
}

// SetHealthData sets the "health_data" field.
func (suo *SettingUpdateOne) SetHealthData(b bool) *SettingUpdateOne {
	suo.mutation.SetHealthData(b)
	return suo
}

// SetNillableHealthData sets the "health_data" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableHealthData(b *bool) *SettingUpdateOne {
	if b != nil {
		suo.SetHealthData(*b)
	}
	return suo
}

// ClearHealthData clears the value of the "health_data" field.
func (suo *SettingUpdateOne) ClearHealthData() *SettingUpdateOne {
	suo.mutation.ClearHealthData()
	return suo
}

// SetTimeDewLocation sets the "time_dew_location" field.
func (suo *SettingUpdateOne) SetTimeDewLocation(b bool) *SettingUpdateOne {
	suo.mutation.SetTimeDewLocation(b)
	return suo
}

// SetNillableTimeDewLocation sets the "time_dew_location" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableTimeDewLocation(b *bool) *SettingUpdateOne {
	if b != nil {
		suo.SetTimeDewLocation(*b)
	}
	return suo
}

// ClearTimeDewLocation clears the value of the "time_dew_location" field.
func (suo *SettingUpdateOne) ClearTimeDewLocation() *SettingUpdateOne {
	suo.mutation.ClearTimeDewLocation()
	return suo
}

// SetTimeDewMicrophone sets the "time_dew_microphone" field.
func (suo *SettingUpdateOne) SetTimeDewMicrophone(b bool) *SettingUpdateOne {
	suo.mutation.SetTimeDewMicrophone(b)
	return suo
}

// SetNillableTimeDewMicrophone sets the "time_dew_microphone" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableTimeDewMicrophone(b *bool) *SettingUpdateOne {
	if b != nil {
		suo.SetTimeDewMicrophone(*b)
	}
	return suo
}

// ClearTimeDewMicrophone clears the value of the "time_dew_microphone" field.
func (suo *SettingUpdateOne) ClearTimeDewMicrophone() *SettingUpdateOne {
	suo.mutation.ClearTimeDewMicrophone()
	return suo
}

// SetTimeDewLora sets the "time_dew_Lora" field.
func (suo *SettingUpdateOne) SetTimeDewLora(b bool) *SettingUpdateOne {
	suo.mutation.SetTimeDewLora(b)
	return suo
}

// SetNillableTimeDewLora sets the "time_dew_Lora" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableTimeDewLora(b *bool) *SettingUpdateOne {
	if b != nil {
		suo.SetTimeDewLora(*b)
	}
	return suo
}

// ClearTimeDewLora clears the value of the "time_dew_Lora" field.
func (suo *SettingUpdateOne) ClearTimeDewLora() *SettingUpdateOne {
	suo.mutation.ClearTimeDewLora()
	return suo
}

// SetPublicCollection sets the "public_collection" field.
func (suo *SettingUpdateOne) SetPublicCollection(b bool) *SettingUpdateOne {
	suo.mutation.SetPublicCollection(b)
	return suo
}

// SetNillablePublicCollection sets the "public_collection" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillablePublicCollection(b *bool) *SettingUpdateOne {
	if b != nil {
		suo.SetPublicCollection(*b)
	}
	return suo
}

// ClearPublicCollection clears the value of the "public_collection" field.
func (suo *SettingUpdateOne) ClearPublicCollection() *SettingUpdateOne {
	suo.mutation.ClearPublicCollection()
	return suo
}

// SetUserID sets the "user_id" field.
func (suo *SettingUpdateOne) SetUserID(u uint64) *SettingUpdateOne {
	suo.mutation.SetUserID(u)
	return suo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (suo *SettingUpdateOne) SetOwnerID(id uint64) *SettingUpdateOne {
	suo.mutation.SetOwnerID(id)
	return suo
}

// SetOwner sets the "owner" edge to the User entity.
func (suo *SettingUpdateOne) SetOwner(u *User) *SettingUpdateOne {
	return suo.SetOwnerID(u.ID)
}

// Mutation returns the SettingMutation object of the builder.
func (suo *SettingUpdateOne) Mutation() *SettingMutation {
	return suo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (suo *SettingUpdateOne) ClearOwner() *SettingUpdateOne {
	suo.mutation.ClearOwner()
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SettingUpdateOne) Select(field string, fields ...string) *SettingUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Setting entity.
func (suo *SettingUpdateOne) Save(ctx context.Context) (*Setting, error) {
	var (
		err  error
		node *Setting
	)
	suo.defaults()
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, suo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Setting)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SettingMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SettingUpdateOne) SaveX(ctx context.Context) *Setting {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SettingUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SettingUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SettingUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdateTime(); !ok {
		v := setting.UpdateDefaultUpdateTime()
		suo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SettingUpdateOne) check() error {
	if _, ok := suo.mutation.OwnerID(); suo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Setting.owner"`)
	}
	return nil
}

func (suo *SettingUpdateOne) sqlSave(ctx context.Context) (_node *Setting, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   setting.Table,
			Columns: setting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: setting.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Setting.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, setting.FieldID)
		for _, f := range fields {
			if !setting.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != setting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdateTime(); ok {
		_spec.SetField(setting.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := suo.mutation.FriendsOnline(); ok {
		_spec.SetField(setting.FieldFriendsOnline, field.TypeBool, value)
	}
	if suo.mutation.FriendsOnlineCleared() {
		_spec.ClearField(setting.FieldFriendsOnline, field.TypeBool)
	}
	if value, ok := suo.mutation.TimeDewFromFriends(); ok {
		_spec.SetField(setting.FieldTimeDewFromFriends, field.TypeBool, value)
	}
	if suo.mutation.TimeDewFromFriendsCleared() {
		_spec.ClearField(setting.FieldTimeDewFromFriends, field.TypeBool)
	}
	if value, ok := suo.mutation.DetailedNotification(); ok {
		_spec.SetField(setting.FieldDetailedNotification, field.TypeBool, value)
	}
	if suo.mutation.DetailedNotificationCleared() {
		_spec.ClearField(setting.FieldDetailedNotification, field.TypeBool)
	}
	if value, ok := suo.mutation.ReceiveFieldInvitation(); ok {
		_spec.SetField(setting.FieldReceiveFieldInvitation, field.TypeBool, value)
	}
	if suo.mutation.ReceiveFieldInvitationCleared() {
		_spec.ClearField(setting.FieldReceiveFieldInvitation, field.TypeBool)
	}
	if value, ok := suo.mutation.SeeMyLocation(); ok {
		_spec.SetField(setting.FieldSeeMyLocation, field.TypeBool, value)
	}
	if suo.mutation.SeeMyLocationCleared() {
		_spec.ClearField(setting.FieldSeeMyLocation, field.TypeBool)
	}
	if value, ok := suo.mutation.Camera(); ok {
		_spec.SetField(setting.FieldCamera, field.TypeBool, value)
	}
	if suo.mutation.CameraCleared() {
		_spec.ClearField(setting.FieldCamera, field.TypeBool)
	}
	if value, ok := suo.mutation.Microphone(); ok {
		_spec.SetField(setting.FieldMicrophone, field.TypeBool, value)
	}
	if suo.mutation.MicrophoneCleared() {
		_spec.ClearField(setting.FieldMicrophone, field.TypeBool)
	}
	if value, ok := suo.mutation.HealthData(); ok {
		_spec.SetField(setting.FieldHealthData, field.TypeBool, value)
	}
	if suo.mutation.HealthDataCleared() {
		_spec.ClearField(setting.FieldHealthData, field.TypeBool)
	}
	if value, ok := suo.mutation.TimeDewLocation(); ok {
		_spec.SetField(setting.FieldTimeDewLocation, field.TypeBool, value)
	}
	if suo.mutation.TimeDewLocationCleared() {
		_spec.ClearField(setting.FieldTimeDewLocation, field.TypeBool)
	}
	if value, ok := suo.mutation.TimeDewMicrophone(); ok {
		_spec.SetField(setting.FieldTimeDewMicrophone, field.TypeBool, value)
	}
	if suo.mutation.TimeDewMicrophoneCleared() {
		_spec.ClearField(setting.FieldTimeDewMicrophone, field.TypeBool)
	}
	if value, ok := suo.mutation.TimeDewLora(); ok {
		_spec.SetField(setting.FieldTimeDewLora, field.TypeBool, value)
	}
	if suo.mutation.TimeDewLoraCleared() {
		_spec.ClearField(setting.FieldTimeDewLora, field.TypeBool)
	}
	if value, ok := suo.mutation.PublicCollection(); ok {
		_spec.SetField(setting.FieldPublicCollection, field.TypeBool, value)
	}
	if suo.mutation.PublicCollectionCleared() {
		_spec.ClearField(setting.FieldPublicCollection, field.TypeBool)
	}
	if suo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   setting.OwnerTable,
			Columns: []string{setting.OwnerColumn},
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
	if nodes := suo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   setting.OwnerTable,
			Columns: []string{setting.OwnerColumn},
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
	_node = &Setting{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{setting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}