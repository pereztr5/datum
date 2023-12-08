// Code generated by ent, DO NOT EDIT.

package ohauthtootoken

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the ohauthtootoken type in the database.
	Label = "oh_auth_too_token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldClientID holds the string denoting the client_id field in the database.
	FieldClientID = "client_id"
	// FieldScopes holds the string denoting the scopes field in the database.
	FieldScopes = "scopes"
	// FieldNonce holds the string denoting the nonce field in the database.
	FieldNonce = "nonce"
	// FieldClaimsUserID holds the string denoting the claims_user_id field in the database.
	FieldClaimsUserID = "claims_user_id"
	// FieldClaimsUsername holds the string denoting the claims_username field in the database.
	FieldClaimsUsername = "claims_username"
	// FieldClaimsEmail holds the string denoting the claims_email field in the database.
	FieldClaimsEmail = "claims_email"
	// FieldClaimsEmailVerified holds the string denoting the claims_email_verified field in the database.
	FieldClaimsEmailVerified = "claims_email_verified"
	// FieldClaimsGroups holds the string denoting the claims_groups field in the database.
	FieldClaimsGroups = "claims_groups"
	// FieldClaimsPreferredUsername holds the string denoting the claims_preferred_username field in the database.
	FieldClaimsPreferredUsername = "claims_preferred_username"
	// FieldConnectorID holds the string denoting the connector_id field in the database.
	FieldConnectorID = "connector_id"
	// FieldConnectorData holds the string denoting the connector_data field in the database.
	FieldConnectorData = "connector_data"
	// FieldLastUsed holds the string denoting the last_used field in the database.
	FieldLastUsed = "last_used"
	// Table holds the table name of the ohauthtootoken in the database.
	Table = "oh_auth_too_tokens"
)

// Columns holds all SQL columns for ohauthtootoken fields.
var Columns = []string{
	FieldID,
	FieldClientID,
	FieldScopes,
	FieldNonce,
	FieldClaimsUserID,
	FieldClaimsUsername,
	FieldClaimsEmail,
	FieldClaimsEmailVerified,
	FieldClaimsGroups,
	FieldClaimsPreferredUsername,
	FieldConnectorID,
	FieldConnectorData,
	FieldLastUsed,
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
	// ClientIDValidator is a validator for the "client_id" field. It is called by the builders before save.
	ClientIDValidator func(string) error
	// NonceValidator is a validator for the "nonce" field. It is called by the builders before save.
	NonceValidator func(string) error
	// ClaimsUserIDValidator is a validator for the "claims_user_id" field. It is called by the builders before save.
	ClaimsUserIDValidator func(string) error
	// ClaimsUsernameValidator is a validator for the "claims_username" field. It is called by the builders before save.
	ClaimsUsernameValidator func(string) error
	// ClaimsEmailValidator is a validator for the "claims_email" field. It is called by the builders before save.
	ClaimsEmailValidator func(string) error
	// ConnectorIDValidator is a validator for the "connector_id" field. It is called by the builders before save.
	ConnectorIDValidator func(string) error
	// DefaultLastUsed holds the default value on creation for the "last_used" field.
	DefaultLastUsed func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the OhAuthTooToken queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByClientID orders the results by the client_id field.
func ByClientID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientID, opts...).ToFunc()
}

// ByNonce orders the results by the nonce field.
func ByNonce(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNonce, opts...).ToFunc()
}

// ByClaimsUserID orders the results by the claims_user_id field.
func ByClaimsUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClaimsUserID, opts...).ToFunc()
}

// ByClaimsUsername orders the results by the claims_username field.
func ByClaimsUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClaimsUsername, opts...).ToFunc()
}

// ByClaimsEmail orders the results by the claims_email field.
func ByClaimsEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClaimsEmail, opts...).ToFunc()
}

// ByClaimsEmailVerified orders the results by the claims_email_verified field.
func ByClaimsEmailVerified(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClaimsEmailVerified, opts...).ToFunc()
}

// ByClaimsPreferredUsername orders the results by the claims_preferred_username field.
func ByClaimsPreferredUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClaimsPreferredUsername, opts...).ToFunc()
}

// ByConnectorID orders the results by the connector_id field.
func ByConnectorID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConnectorID, opts...).ToFunc()
}

// ByLastUsed orders the results by the last_used field.
func ByLastUsed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastUsed, opts...).ToFunc()
}