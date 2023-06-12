// Code generated by ent, DO NOT EDIT.

package ent

import (
	"CSBackendTmp/ent/csfield"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// CSField is the model entity for the CSField schema.
type CSField struct {
	config `json:"-"`
	// ID of the ent.
	// id
	ID uint64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// 场名称
	Name string `json:"name,omitempty"`
	// 场状态
	Status csfield.Status `json:"status,omitempty"`
	// 场类型
	Type csfield.Type `json:"type,omitempty"`
	// 场模式
	Mode csfield.Mode `json:"mode,omitempty"`
	// 场隐私等级
	PrivateLevel csfield.PrivateLevel `json:"private_level,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uint64 `json:"user_id,omitempty"`
	// MasterID holds the value of the "master_id" field.
	MasterID uint64 `json:"master_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CSFieldQuery when eager-loading is set.
	Edges CSFieldEdges `json:"edges"`
}

// CSFieldEdges holds the relations/edges for other nodes in the graph.
type CSFieldEdges struct {
	// JoinedUser holds the value of the joined_user edge.
	JoinedUser []*User `json:"joined_user,omitempty"`
	// Joins holds the value of the joins edge.
	Joins []*Join `json:"joins,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// JoinedUserOrErr returns the JoinedUser value or an error if the edge
// was not loaded in eager-loading.
func (e CSFieldEdges) JoinedUserOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.JoinedUser, nil
	}
	return nil, &NotLoadedError{edge: "joined_user"}
}

// JoinsOrErr returns the Joins value or an error if the edge
// was not loaded in eager-loading.
func (e CSFieldEdges) JoinsOrErr() ([]*Join, error) {
	if e.loadedTypes[1] {
		return e.Joins, nil
	}
	return nil, &NotLoadedError{edge: "joins"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CSField) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case csfield.FieldID, csfield.FieldUserID, csfield.FieldMasterID:
			values[i] = new(sql.NullInt64)
		case csfield.FieldName, csfield.FieldStatus, csfield.FieldType, csfield.FieldMode, csfield.FieldPrivateLevel:
			values[i] = new(sql.NullString)
		case csfield.FieldCreateTime, csfield.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CSField", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CSField fields.
func (cf *CSField) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case csfield.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cf.ID = uint64(value.Int64)
		case csfield.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				cf.CreateTime = value.Time
			}
		case csfield.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				cf.UpdateTime = value.Time
			}
		case csfield.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				cf.Name = value.String
			}
		case csfield.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				cf.Status = csfield.Status(value.String)
			}
		case csfield.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				cf.Type = csfield.Type(value.String)
			}
		case csfield.FieldMode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mode", values[i])
			} else if value.Valid {
				cf.Mode = csfield.Mode(value.String)
			}
		case csfield.FieldPrivateLevel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field private_level", values[i])
			} else if value.Valid {
				cf.PrivateLevel = csfield.PrivateLevel(value.String)
			}
		case csfield.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				cf.UserID = uint64(value.Int64)
			}
		case csfield.FieldMasterID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field master_id", values[i])
			} else if value.Valid {
				cf.MasterID = uint64(value.Int64)
			}
		}
	}
	return nil
}

// QueryJoinedUser queries the "joined_user" edge of the CSField entity.
func (cf *CSField) QueryJoinedUser() *UserQuery {
	return (&CSFieldClient{config: cf.config}).QueryJoinedUser(cf)
}

// QueryJoins queries the "joins" edge of the CSField entity.
func (cf *CSField) QueryJoins() *JoinQuery {
	return (&CSFieldClient{config: cf.config}).QueryJoins(cf)
}

// Update returns a builder for updating this CSField.
// Note that you need to call CSField.Unwrap() before calling this method if this CSField
// was returned from a transaction, and the transaction was committed or rolled back.
func (cf *CSField) Update() *CSFieldUpdateOne {
	return (&CSFieldClient{config: cf.config}).UpdateOne(cf)
}

// Unwrap unwraps the CSField entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cf *CSField) Unwrap() *CSField {
	_tx, ok := cf.config.driver.(*txDriver)
	if !ok {
		panic("ent: CSField is not a transactional entity")
	}
	cf.config.driver = _tx.drv
	return cf
}

// String implements the fmt.Stringer.
func (cf *CSField) String() string {
	var builder strings.Builder
	builder.WriteString("CSField(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cf.ID))
	builder.WriteString("create_time=")
	builder.WriteString(cf.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(cf.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(cf.Name)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", cf.Status))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", cf.Type))
	builder.WriteString(", ")
	builder.WriteString("mode=")
	builder.WriteString(fmt.Sprintf("%v", cf.Mode))
	builder.WriteString(", ")
	builder.WriteString("private_level=")
	builder.WriteString(fmt.Sprintf("%v", cf.PrivateLevel))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", cf.UserID))
	builder.WriteString(", ")
	builder.WriteString("master_id=")
	builder.WriteString(fmt.Sprintf("%v", cf.MasterID))
	builder.WriteByte(')')
	return builder.String()
}

// CSFields is a parsable slice of CSField.
type CSFields []*CSField

func (cf CSFields) config(cfg config) {
	for _i := range cf {
		cf[_i].config = cfg
	}
}
