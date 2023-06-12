// Code generated by ent, DO NOT EDIT.

package user_auth

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the user_auth type in the database.
	Label = "user_auth"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldOauthSource holds the string denoting the oauth_source field in the database.
	FieldOauthSource = "oauth_source"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldMobileNo holds the string denoting the mobile_no field in the database.
	FieldMobileNo = "mobile_no"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldAccessToken holds the string denoting the access_token field in the database.
	FieldAccessToken = "access_token"
	// FieldOauthTokenType holds the string denoting the oauth_token_type field in the database.
	FieldOauthTokenType = "oauth_token_type"
	// FieldOauthRefreshToken holds the string denoting the oauth_refresh_token field in the database.
	FieldOauthRefreshToken = "oauth_refresh_token"
	// FieldOauthID holds the string denoting the oauth_id field in the database.
	FieldOauthID = "oauth_id"
	// FieldIsFinished holds the string denoting the is_finished field in the database.
	FieldIsFinished = "is_finished"
	// FieldOauthExpiry holds the string denoting the oauth_expiry field in the database.
	FieldOauthExpiry = "oauth_expiry"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the user_auth in the database.
	Table = "user_auths"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "user_auths"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_id"
)

// Columns holds all SQL columns for user_auth fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldType,
	FieldOauthSource,
	FieldEmail,
	FieldMobileNo,
	FieldPassword,
	FieldAccessToken,
	FieldOauthTokenType,
	FieldOauthRefreshToken,
	FieldOauthID,
	FieldIsFinished,
	FieldOauthExpiry,
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
	TypeOauth2 Type = "oauth2"
	TypeEmail  Type = "email"
	TypeMobile Type = "mobile"
	TypeJwt    Type = "jwt"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeOauth2, TypeEmail, TypeMobile, TypeJwt:
		return nil
	default:
		return fmt.Errorf("user_auth: invalid enum value for type field: %q", _type)
	}
}
