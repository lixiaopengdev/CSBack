// Code generated by ent, DO NOT EDIT.

package ent

import (
	"CSBackendTmp/ent/user"
	"CSBackendTmp/ent/user_history"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// User_history is the model entity for the User_history schema.
type User_history struct {
	config `json:"-"`
	// ID of the ent.
	// id
	ID uint64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// 用户历史类型
	Type user_history.Type `json:"type,omitempty"`
	// 设备特征码
	Name string `json:"name,omitempty"`
	// 资源url
	ResourceURL string `json:"resource_url,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uint64 `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the User_historyQuery when eager-loading is set.
	Edges User_historyEdges `json:"edges"`
}

// User_historyEdges holds the relations/edges for other nodes in the graph.
type User_historyEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e User_historyEdges) OwnerOrErr() (*User, error) {
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
func (*User_history) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user_history.FieldID, user_history.FieldUserID:
			values[i] = new(sql.NullInt64)
		case user_history.FieldType, user_history.FieldName, user_history.FieldResourceURL:
			values[i] = new(sql.NullString)
		case user_history.FieldCreateTime, user_history.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User_history", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User_history fields.
func (uh *User_history) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user_history.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			uh.ID = uint64(value.Int64)
		case user_history.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				uh.CreateTime = value.Time
			}
		case user_history.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				uh.UpdateTime = value.Time
			}
		case user_history.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				uh.Type = user_history.Type(value.String)
			}
		case user_history.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				uh.Name = value.String
			}
		case user_history.FieldResourceURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field resource_url", values[i])
			} else if value.Valid {
				uh.ResourceURL = value.String
			}
		case user_history.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				uh.UserID = uint64(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the User_history entity.
func (uh *User_history) QueryOwner() *UserQuery {
	return (&User_historyClient{config: uh.config}).QueryOwner(uh)
}

// Update returns a builder for updating this User_history.
// Note that you need to call User_history.Unwrap() before calling this method if this User_history
// was returned from a transaction, and the transaction was committed or rolled back.
func (uh *User_history) Update() *UserHistoryUpdateOne {
	return (&User_historyClient{config: uh.config}).UpdateOne(uh)
}

// Unwrap unwraps the User_history entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uh *User_history) Unwrap() *User_history {
	_tx, ok := uh.config.driver.(*txDriver)
	if !ok {
		panic("ent: User_history is not a transactional entity")
	}
	uh.config.driver = _tx.drv
	return uh
}

// String implements the fmt.Stringer.
func (uh *User_history) String() string {
	var builder strings.Builder
	builder.WriteString("User_history(")
	builder.WriteString(fmt.Sprintf("id=%v, ", uh.ID))
	builder.WriteString("create_time=")
	builder.WriteString(uh.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(uh.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", uh.Type))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(uh.Name)
	builder.WriteString(", ")
	builder.WriteString("resource_url=")
	builder.WriteString(uh.ResourceURL)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", uh.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// User_histories is a parsable slice of User_history.
type User_histories []*User_history

func (uh User_histories) config(cfg config) {
	for _i := range uh {
		uh[_i].config = cfg
	}
}