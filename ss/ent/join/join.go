// Code generated by ent, DO NOT EDIT.

package join

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the join type in the database.
	Label = "join"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldJoinAt holds the string denoting the join_at field in the database.
	FieldJoinAt = "join_at"
	// FieldLeaveAt holds the string denoting the leave_at field in the database.
	FieldLeaveAt = "leave_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldCsFieldID holds the string denoting the cs_field_id field in the database.
	FieldCsFieldID = "cs_field_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeCsField holds the string denoting the cs_field edge name in mutations.
	EdgeCsField = "cs_field"
	// Table holds the table name of the join in the database.
	Table = "joins"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "joins"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// CsFieldTable is the table that holds the cs_field relation/edge.
	CsFieldTable = "joins"
	// CsFieldInverseTable is the table name for the CSField entity.
	// It exists in this package in order to avoid circular dependency with the "csfield" package.
	CsFieldInverseTable = "cs_fields"
	// CsFieldColumn is the table column denoting the cs_field relation/edge.
	CsFieldColumn = "cs_field_id"
)

// Columns holds all SQL columns for join fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldJoinAt,
	FieldLeaveAt,
	FieldStatus,
	FieldUserID,
	FieldCsFieldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultJoinAt holds the default value on creation for the "join_at" field.
	DefaultJoinAt func() time.Time
)

// Status defines the type for the "status" enum field.
type Status string

// StatusHost is the default value of the Status enum.
const DefaultStatus = StatusHost

// Status values.
const (
	StatusInfield     Status = "infield"
	StatusTempLeaving Status = "temp_leaving"
	StatusInvited     Status = "invited"
	StatusLeave       Status = "leave"
	StatusHost        Status = "host"
	StatusAdmin       Status = "admin"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusInfield, StatusTempLeaving, StatusInvited, StatusLeave, StatusHost, StatusAdmin:
		return nil
	default:
		return fmt.Errorf("join: invalid enum value for status field: %q", s)
	}
}
