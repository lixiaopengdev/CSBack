// Code generated by ent, DO NOT EDIT.

package ent

import (
	"CSBackendTmp/ent/mask"
	"CSBackendTmp/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Mask is the model entity for the Mask schema.
type Mask struct {
	config `json:"-"`
	// ID of the ent.
	// id
	ID uint64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// 面具名称
	Name string `json:"name,omitempty"`
	// 面具描述
	Desc string `json:"desc,omitempty"`
	// GUID
	GUID string `json:"GUID,omitempty"`
	// 缩略图
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	// 面具状态
	Status mask.Status `json:"status,omitempty"`
	// 面具类型
	Type mask.Type `json:"type,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uint64 `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MaskQuery when eager-loading is set.
	Edges MaskEdges `json:"edges"`
}

// MaskEdges holds the relations/edges for other nodes in the graph.
type MaskEdges struct {
	// Bundle holds the value of the bundle edge.
	Bundle []*Bundle `json:"bundle,omitempty"`
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// BundleOrErr returns the Bundle value or an error if the edge
// was not loaded in eager-loading.
func (e MaskEdges) BundleOrErr() ([]*Bundle, error) {
	if e.loadedTypes[0] {
		return e.Bundle, nil
	}
	return nil, &NotLoadedError{edge: "bundle"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MaskEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Mask) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case mask.FieldID, mask.FieldUserID:
			values[i] = new(sql.NullInt64)
		case mask.FieldName, mask.FieldDesc, mask.FieldGUID, mask.FieldThumbnailURL, mask.FieldStatus, mask.FieldType:
			values[i] = new(sql.NullString)
		case mask.FieldCreateTime, mask.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Mask", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Mask fields.
func (m *Mask) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mask.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = uint64(value.Int64)
		case mask.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				m.CreateTime = value.Time
			}
		case mask.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				m.UpdateTime = value.Time
			}
		case mask.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case mask.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				m.Desc = value.String
			}
		case mask.FieldGUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field GUID", values[i])
			} else if value.Valid {
				m.GUID = value.String
			}
		case mask.FieldThumbnailURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field thumbnail_url", values[i])
			} else if value.Valid {
				m.ThumbnailURL = value.String
			}
		case mask.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				m.Status = mask.Status(value.String)
			}
		case mask.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				m.Type = mask.Type(value.String)
			}
		case mask.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				m.UserID = uint64(value.Int64)
			}
		}
	}
	return nil
}

// QueryBundle queries the "bundle" edge of the Mask entity.
func (m *Mask) QueryBundle() *BundleQuery {
	return (&MaskClient{config: m.config}).QueryBundle(m)
}

// QueryOwner queries the "owner" edge of the Mask entity.
func (m *Mask) QueryOwner() *UserQuery {
	return (&MaskClient{config: m.config}).QueryOwner(m)
}

// Update returns a builder for updating this Mask.
// Note that you need to call Mask.Unwrap() before calling this method if this Mask
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Mask) Update() *MaskUpdateOne {
	return (&MaskClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Mask entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Mask) Unwrap() *Mask {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Mask is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Mask) String() string {
	var builder strings.Builder
	builder.WriteString("Mask(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("create_time=")
	builder.WriteString(m.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(m.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(m.Desc)
	builder.WriteString(", ")
	builder.WriteString("GUID=")
	builder.WriteString(m.GUID)
	builder.WriteString(", ")
	builder.WriteString("thumbnail_url=")
	builder.WriteString(m.ThumbnailURL)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", m.Status))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", m.Type))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", m.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// Masks is a parsable slice of Mask.
type Masks []*Mask

func (m Masks) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
