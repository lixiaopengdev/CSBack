// Code generated by ent, DO NOT EDIT.

package ent

import (
	"CSBackendTmp/ent/stream"
	"CSBackendTmp/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Stream is the model entity for the Stream schema.
type Stream struct {
	config `json:"-"`
	// ID of the ent.
	// id
	ID uint64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 流类型
	Type stream.Type `json:"type,omitempty"`
	// 流地址
	StreamURL string `json:"stream_url,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uint64 `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StreamQuery when eager-loading is set.
	Edges StreamEdges `json:"edges"`
}

// StreamEdges holds the relations/edges for other nodes in the graph.
type StreamEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StreamEdges) OwnerOrErr() (*User, error) {
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
func (*Stream) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case stream.FieldID, stream.FieldUserID:
			values[i] = new(sql.NullInt64)
		case stream.FieldName, stream.FieldType, stream.FieldStreamURL:
			values[i] = new(sql.NullString)
		case stream.FieldCreateTime, stream.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Stream", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Stream fields.
func (s *Stream) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case stream.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = uint64(value.Int64)
		case stream.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				s.CreateTime = value.Time
			}
		case stream.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				s.UpdateTime = value.Time
			}
		case stream.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case stream.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				s.Type = stream.Type(value.String)
			}
		case stream.FieldStreamURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field stream_url", values[i])
			} else if value.Valid {
				s.StreamURL = value.String
			}
		case stream.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				s.UserID = uint64(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Stream entity.
func (s *Stream) QueryOwner() *UserQuery {
	return (&StreamClient{config: s.config}).QueryOwner(s)
}

// Update returns a builder for updating this Stream.
// Note that you need to call Stream.Unwrap() before calling this method if this Stream
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Stream) Update() *StreamUpdateOne {
	return (&StreamClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Stream entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Stream) Unwrap() *Stream {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Stream is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Stream) String() string {
	var builder strings.Builder
	builder.WriteString("Stream(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("create_time=")
	builder.WriteString(s.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(s.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", s.Type))
	builder.WriteString(", ")
	builder.WriteString("stream_url=")
	builder.WriteString(s.StreamURL)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", s.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// Streams is a parsable slice of Stream.
type Streams []*Stream

func (s Streams) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
