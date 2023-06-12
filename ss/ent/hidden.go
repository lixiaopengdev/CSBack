// Code generated by ent, DO NOT EDIT.

package ent

import (
	"CSBackendTmp/ent/hidden"
	"CSBackendTmp/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Hidden is the model entity for the Hidden schema.
type Hidden struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uint64 `json:"user_id,omitempty"`
	// HiddenID holds the value of the "hidden_id" field.
	HiddenID uint64 `json:"hidden_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HiddenQuery when eager-loading is set.
	Edges HiddenEdges `json:"edges"`
}

// HiddenEdges holds the relations/edges for other nodes in the graph.
type HiddenEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HiddenEdges) OwnerOrErr() (*User, error) {
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
func (*Hidden) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case hidden.FieldID, hidden.FieldUserID, hidden.FieldHiddenID:
			values[i] = new(sql.NullInt64)
		case hidden.FieldCreateTime, hidden.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Hidden", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Hidden fields.
func (h *Hidden) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case hidden.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			h.ID = int(value.Int64)
		case hidden.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				h.CreateTime = value.Time
			}
		case hidden.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				h.UpdateTime = value.Time
			}
		case hidden.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				h.UserID = uint64(value.Int64)
			}
		case hidden.FieldHiddenID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field hidden_id", values[i])
			} else if value.Valid {
				h.HiddenID = uint64(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Hidden entity.
func (h *Hidden) QueryOwner() *UserQuery {
	return (&HiddenClient{config: h.config}).QueryOwner(h)
}

// Update returns a builder for updating this Hidden.
// Note that you need to call Hidden.Unwrap() before calling this method if this Hidden
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *Hidden) Update() *HiddenUpdateOne {
	return (&HiddenClient{config: h.config}).UpdateOne(h)
}

// Unwrap unwraps the Hidden entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *Hidden) Unwrap() *Hidden {
	_tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: Hidden is not a transactional entity")
	}
	h.config.driver = _tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *Hidden) String() string {
	var builder strings.Builder
	builder.WriteString("Hidden(")
	builder.WriteString(fmt.Sprintf("id=%v, ", h.ID))
	builder.WriteString("create_time=")
	builder.WriteString(h.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(h.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", h.UserID))
	builder.WriteString(", ")
	builder.WriteString("hidden_id=")
	builder.WriteString(fmt.Sprintf("%v", h.HiddenID))
	builder.WriteByte(')')
	return builder.String()
}

// Hiddens is a parsable slice of Hidden.
type Hiddens []*Hidden

func (h Hiddens) config(cfg config) {
	for _i := range h {
		h[_i].config = cfg
	}
}