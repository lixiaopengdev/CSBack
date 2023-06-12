// Code generated by ent, DO NOT EDIT.

package ent

import (
	"CSBackendTmp/ent/setting"
	"CSBackendTmp/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Setting is the model entity for the Setting schema.
type Setting struct {
	config `json:"-"`
	// ID of the ent.
	// id
	ID uint64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// FriendsOnline holds the value of the "friends_online" field.
	FriendsOnline bool `json:"friends_online,omitempty"`
	// TimeDewFromFriends holds the value of the "time_dew_from_friends" field.
	TimeDewFromFriends bool `json:"time_dew_from_friends,omitempty"`
	// DetailedNotification holds the value of the "detailed_notification" field.
	DetailedNotification bool `json:"detailed_notification,omitempty"`
	// ReceiveFieldInvitation holds the value of the "receive_field_invitation" field.
	ReceiveFieldInvitation bool `json:"receive_field_invitation,omitempty"`
	// SeeMyLocation holds the value of the "see_my_location" field.
	SeeMyLocation bool `json:"see_my_location,omitempty"`
	// Camera holds the value of the "camera" field.
	Camera bool `json:"camera,omitempty"`
	// Microphone holds the value of the "microphone" field.
	Microphone bool `json:"microphone,omitempty"`
	// HealthData holds the value of the "health_data" field.
	HealthData bool `json:"health_data,omitempty"`
	// TimeDewLocation holds the value of the "time_dew_location" field.
	TimeDewLocation bool `json:"time_dew_location,omitempty"`
	// TimeDewMicrophone holds the value of the "time_dew_microphone" field.
	TimeDewMicrophone bool `json:"time_dew_microphone,omitempty"`
	// TimeDewLora holds the value of the "time_dew_Lora" field.
	TimeDewLora bool `json:"time_dew_Lora,omitempty"`
	// PublicCollection holds the value of the "public_collection" field.
	PublicCollection bool `json:"public_collection,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uint64 `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SettingQuery when eager-loading is set.
	Edges SettingEdges `json:"edges"`
}

// SettingEdges holds the relations/edges for other nodes in the graph.
type SettingEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SettingEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Setting) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case setting.FieldFriendsOnline, setting.FieldTimeDewFromFriends, setting.FieldDetailedNotification, setting.FieldReceiveFieldInvitation, setting.FieldSeeMyLocation, setting.FieldCamera, setting.FieldMicrophone, setting.FieldHealthData, setting.FieldTimeDewLocation, setting.FieldTimeDewMicrophone, setting.FieldTimeDewLora, setting.FieldPublicCollection:
			values[i] = new(sql.NullBool)
		case setting.FieldID, setting.FieldUserID:
			values[i] = new(sql.NullInt64)
		case setting.FieldCreateTime, setting.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Setting", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Setting fields.
func (s *Setting) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case setting.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = uint64(value.Int64)
		case setting.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				s.CreateTime = value.Time
			}
		case setting.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				s.UpdateTime = value.Time
			}
		case setting.FieldFriendsOnline:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field friends_online", values[i])
			} else if value.Valid {
				s.FriendsOnline = value.Bool
			}
		case setting.FieldTimeDewFromFriends:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field time_dew_from_friends", values[i])
			} else if value.Valid {
				s.TimeDewFromFriends = value.Bool
			}
		case setting.FieldDetailedNotification:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field detailed_notification", values[i])
			} else if value.Valid {
				s.DetailedNotification = value.Bool
			}
		case setting.FieldReceiveFieldInvitation:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field receive_field_invitation", values[i])
			} else if value.Valid {
				s.ReceiveFieldInvitation = value.Bool
			}
		case setting.FieldSeeMyLocation:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field see_my_location", values[i])
			} else if value.Valid {
				s.SeeMyLocation = value.Bool
			}
		case setting.FieldCamera:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field camera", values[i])
			} else if value.Valid {
				s.Camera = value.Bool
			}
		case setting.FieldMicrophone:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field microphone", values[i])
			} else if value.Valid {
				s.Microphone = value.Bool
			}
		case setting.FieldHealthData:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field health_data", values[i])
			} else if value.Valid {
				s.HealthData = value.Bool
			}
		case setting.FieldTimeDewLocation:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field time_dew_location", values[i])
			} else if value.Valid {
				s.TimeDewLocation = value.Bool
			}
		case setting.FieldTimeDewMicrophone:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field time_dew_microphone", values[i])
			} else if value.Valid {
				s.TimeDewMicrophone = value.Bool
			}
		case setting.FieldTimeDewLora:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field time_dew_Lora", values[i])
			} else if value.Valid {
				s.TimeDewLora = value.Bool
			}
		case setting.FieldPublicCollection:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field public_collection", values[i])
			} else if value.Valid {
				s.PublicCollection = value.Bool
			}
		case setting.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				s.UserID = uint64(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Setting entity.
func (s *Setting) QueryOwner() *UserQuery {
	return (&SettingClient{config: s.config}).QueryOwner(s)
}

// Update returns a builder for updating this Setting.
// Note that you need to call Setting.Unwrap() before calling this method if this Setting
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Setting) Update() *SettingUpdateOne {
	return (&SettingClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Setting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Setting) Unwrap() *Setting {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Setting is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Setting) String() string {
	var builder strings.Builder
	builder.WriteString("Setting(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("create_time=")
	builder.WriteString(s.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(s.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("friends_online=")
	builder.WriteString(fmt.Sprintf("%v", s.FriendsOnline))
	builder.WriteString(", ")
	builder.WriteString("time_dew_from_friends=")
	builder.WriteString(fmt.Sprintf("%v", s.TimeDewFromFriends))
	builder.WriteString(", ")
	builder.WriteString("detailed_notification=")
	builder.WriteString(fmt.Sprintf("%v", s.DetailedNotification))
	builder.WriteString(", ")
	builder.WriteString("receive_field_invitation=")
	builder.WriteString(fmt.Sprintf("%v", s.ReceiveFieldInvitation))
	builder.WriteString(", ")
	builder.WriteString("see_my_location=")
	builder.WriteString(fmt.Sprintf("%v", s.SeeMyLocation))
	builder.WriteString(", ")
	builder.WriteString("camera=")
	builder.WriteString(fmt.Sprintf("%v", s.Camera))
	builder.WriteString(", ")
	builder.WriteString("microphone=")
	builder.WriteString(fmt.Sprintf("%v", s.Microphone))
	builder.WriteString(", ")
	builder.WriteString("health_data=")
	builder.WriteString(fmt.Sprintf("%v", s.HealthData))
	builder.WriteString(", ")
	builder.WriteString("time_dew_location=")
	builder.WriteString(fmt.Sprintf("%v", s.TimeDewLocation))
	builder.WriteString(", ")
	builder.WriteString("time_dew_microphone=")
	builder.WriteString(fmt.Sprintf("%v", s.TimeDewMicrophone))
	builder.WriteString(", ")
	builder.WriteString("time_dew_Lora=")
	builder.WriteString(fmt.Sprintf("%v", s.TimeDewLora))
	builder.WriteString(", ")
	builder.WriteString("public_collection=")
	builder.WriteString(fmt.Sprintf("%v", s.PublicCollection))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", s.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// Settings is a parsable slice of Setting.
type Settings []*Setting

func (s Settings) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}