// Code generated by ent, DO NOT EDIT.

package device

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the device type in the database.
	Label = "device"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldPushToken holds the string denoting the push_token field in the database.
	FieldPushToken = "push_token"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the device in the database.
	Table = "devices"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "devices"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_id"
)

// Columns holds all SQL columns for device fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldType,
	FieldCode,
	FieldPushToken,
	FieldUserID,
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
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeIPhone  Type = "iPhone"
	TypeAndroid Type = "android"
	TypeUnknown Type = "unknown"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeIPhone, TypeAndroid, TypeUnknown:
		return nil
	default:
		return fmt.Errorf("device: invalid enum value for type field: %q", _type)
	}
}
