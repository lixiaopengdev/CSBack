// Code generated by ent, DO NOT EDIT.

package ent

import (
	"CSBackendTmp/ent/bundle"
	"CSBackendTmp/ent/mask"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Bundle is the model entity for the Bundle schema.
type Bundle struct {
	config `json:"-"`
	// ID of the ent.
	// id
	ID uint64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// 版本ID
	VerionID uint64 `json:"verionID,omitempty"`
	// bundle资源地址
	BundleURL string `json:"bundle_url,omitempty"`
	// bundle状态
	Status bundle.Status `json:"status,omitempty"`
	// 平台
	Platform bundle.Platform `json:"platform,omitempty"`
	// MaskID holds the value of the "mask_id" field.
	MaskID uint64 `json:"mask_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BundleQuery when eager-loading is set.
	Edges BundleEdges `json:"edges"`
}

// BundleEdges holds the relations/edges for other nodes in the graph.
type BundleEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Mask `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BundleEdges) OwnerOrErr() (*Mask, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: mask.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Bundle) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bundle.FieldID, bundle.FieldVerionID, bundle.FieldMaskID:
			values[i] = new(sql.NullInt64)
		case bundle.FieldBundleURL, bundle.FieldStatus, bundle.FieldPlatform:
			values[i] = new(sql.NullString)
		case bundle.FieldCreateTime, bundle.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Bundle", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Bundle fields.
func (b *Bundle) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bundle.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = uint64(value.Int64)
		case bundle.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				b.CreateTime = value.Time
			}
		case bundle.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				b.UpdateTime = value.Time
			}
		case bundle.FieldVerionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field verionID", values[i])
			} else if value.Valid {
				b.VerionID = uint64(value.Int64)
			}
		case bundle.FieldBundleURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bundle_url", values[i])
			} else if value.Valid {
				b.BundleURL = value.String
			}
		case bundle.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				b.Status = bundle.Status(value.String)
			}
		case bundle.FieldPlatform:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field platform", values[i])
			} else if value.Valid {
				b.Platform = bundle.Platform(value.String)
			}
		case bundle.FieldMaskID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mask_id", values[i])
			} else if value.Valid {
				b.MaskID = uint64(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Bundle entity.
func (b *Bundle) QueryOwner() *MaskQuery {
	return (&BundleClient{config: b.config}).QueryOwner(b)
}

// Update returns a builder for updating this Bundle.
// Note that you need to call Bundle.Unwrap() before calling this method if this Bundle
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Bundle) Update() *BundleUpdateOne {
	return (&BundleClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the Bundle entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Bundle) Unwrap() *Bundle {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Bundle is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Bundle) String() string {
	var builder strings.Builder
	builder.WriteString("Bundle(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("create_time=")
	builder.WriteString(b.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(b.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("verionID=")
	builder.WriteString(fmt.Sprintf("%v", b.VerionID))
	builder.WriteString(", ")
	builder.WriteString("bundle_url=")
	builder.WriteString(b.BundleURL)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", b.Status))
	builder.WriteString(", ")
	builder.WriteString("platform=")
	builder.WriteString(fmt.Sprintf("%v", b.Platform))
	builder.WriteString(", ")
	builder.WriteString("mask_id=")
	builder.WriteString(fmt.Sprintf("%v", b.MaskID))
	builder.WriteByte(')')
	return builder.String()
}

// Bundles is a parsable slice of Bundle.
type Bundles []*Bundle

func (b Bundles) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
