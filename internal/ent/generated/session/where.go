// Code generated by ent, DO NOT EDIT.

package session

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/google/uuid"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldUpdatedBy, v))
}

// Disabled applies equality check predicate on the "disabled" field. It's identical to DisabledEQ.
func Disabled(v bool) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldDisabled, v))
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v string) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldToken, v))
}

// UserAgent applies equality check predicate on the "user_agent" field. It's identical to UserAgentEQ.
func UserAgent(v string) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldUserAgent, v))
}

// Ips applies equality check predicate on the "ips" field. It's identical to IpsEQ.
func Ips(v string) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldIps, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.Session {
	return predicate.Session(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.Session {
	return predicate.Session(sql.FieldNotNull(FieldCreatedBy))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.Session {
	return predicate.Session(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.Session {
	return predicate.Session(sql.FieldNotNull(FieldUpdatedBy))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldType, vs...))
}

// DisabledEQ applies the EQ predicate on the "disabled" field.
func DisabledEQ(v bool) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldDisabled, v))
}

// DisabledNEQ applies the NEQ predicate on the "disabled" field.
func DisabledNEQ(v bool) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldDisabled, v))
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v string) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldToken, v))
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v string) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldToken, v))
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...string) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldToken, vs...))
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...string) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldToken, vs...))
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v string) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldToken, v))
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v string) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldToken, v))
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v string) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldToken, v))
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v string) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldToken, v))
}

// TokenContains applies the Contains predicate on the "token" field.
func TokenContains(v string) predicate.Session {
	return predicate.Session(sql.FieldContains(FieldToken, v))
}

// TokenHasPrefix applies the HasPrefix predicate on the "token" field.
func TokenHasPrefix(v string) predicate.Session {
	return predicate.Session(sql.FieldHasPrefix(FieldToken, v))
}

// TokenHasSuffix applies the HasSuffix predicate on the "token" field.
func TokenHasSuffix(v string) predicate.Session {
	return predicate.Session(sql.FieldHasSuffix(FieldToken, v))
}

// TokenEqualFold applies the EqualFold predicate on the "token" field.
func TokenEqualFold(v string) predicate.Session {
	return predicate.Session(sql.FieldEqualFold(FieldToken, v))
}

// TokenContainsFold applies the ContainsFold predicate on the "token" field.
func TokenContainsFold(v string) predicate.Session {
	return predicate.Session(sql.FieldContainsFold(FieldToken, v))
}

// UserAgentEQ applies the EQ predicate on the "user_agent" field.
func UserAgentEQ(v string) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldUserAgent, v))
}

// UserAgentNEQ applies the NEQ predicate on the "user_agent" field.
func UserAgentNEQ(v string) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldUserAgent, v))
}

// UserAgentIn applies the In predicate on the "user_agent" field.
func UserAgentIn(vs ...string) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldUserAgent, vs...))
}

// UserAgentNotIn applies the NotIn predicate on the "user_agent" field.
func UserAgentNotIn(vs ...string) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldUserAgent, vs...))
}

// UserAgentGT applies the GT predicate on the "user_agent" field.
func UserAgentGT(v string) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldUserAgent, v))
}

// UserAgentGTE applies the GTE predicate on the "user_agent" field.
func UserAgentGTE(v string) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldUserAgent, v))
}

// UserAgentLT applies the LT predicate on the "user_agent" field.
func UserAgentLT(v string) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldUserAgent, v))
}

// UserAgentLTE applies the LTE predicate on the "user_agent" field.
func UserAgentLTE(v string) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldUserAgent, v))
}

// UserAgentContains applies the Contains predicate on the "user_agent" field.
func UserAgentContains(v string) predicate.Session {
	return predicate.Session(sql.FieldContains(FieldUserAgent, v))
}

// UserAgentHasPrefix applies the HasPrefix predicate on the "user_agent" field.
func UserAgentHasPrefix(v string) predicate.Session {
	return predicate.Session(sql.FieldHasPrefix(FieldUserAgent, v))
}

// UserAgentHasSuffix applies the HasSuffix predicate on the "user_agent" field.
func UserAgentHasSuffix(v string) predicate.Session {
	return predicate.Session(sql.FieldHasSuffix(FieldUserAgent, v))
}

// UserAgentIsNil applies the IsNil predicate on the "user_agent" field.
func UserAgentIsNil() predicate.Session {
	return predicate.Session(sql.FieldIsNull(FieldUserAgent))
}

// UserAgentNotNil applies the NotNil predicate on the "user_agent" field.
func UserAgentNotNil() predicate.Session {
	return predicate.Session(sql.FieldNotNull(FieldUserAgent))
}

// UserAgentEqualFold applies the EqualFold predicate on the "user_agent" field.
func UserAgentEqualFold(v string) predicate.Session {
	return predicate.Session(sql.FieldEqualFold(FieldUserAgent, v))
}

// UserAgentContainsFold applies the ContainsFold predicate on the "user_agent" field.
func UserAgentContainsFold(v string) predicate.Session {
	return predicate.Session(sql.FieldContainsFold(FieldUserAgent, v))
}

// IpsEQ applies the EQ predicate on the "ips" field.
func IpsEQ(v string) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldIps, v))
}

// IpsNEQ applies the NEQ predicate on the "ips" field.
func IpsNEQ(v string) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldIps, v))
}

// IpsIn applies the In predicate on the "ips" field.
func IpsIn(vs ...string) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldIps, vs...))
}

// IpsNotIn applies the NotIn predicate on the "ips" field.
func IpsNotIn(vs ...string) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldIps, vs...))
}

// IpsGT applies the GT predicate on the "ips" field.
func IpsGT(v string) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldIps, v))
}

// IpsGTE applies the GTE predicate on the "ips" field.
func IpsGTE(v string) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldIps, v))
}

// IpsLT applies the LT predicate on the "ips" field.
func IpsLT(v string) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldIps, v))
}

// IpsLTE applies the LTE predicate on the "ips" field.
func IpsLTE(v string) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldIps, v))
}

// IpsContains applies the Contains predicate on the "ips" field.
func IpsContains(v string) predicate.Session {
	return predicate.Session(sql.FieldContains(FieldIps, v))
}

// IpsHasPrefix applies the HasPrefix predicate on the "ips" field.
func IpsHasPrefix(v string) predicate.Session {
	return predicate.Session(sql.FieldHasPrefix(FieldIps, v))
}

// IpsHasSuffix applies the HasSuffix predicate on the "ips" field.
func IpsHasSuffix(v string) predicate.Session {
	return predicate.Session(sql.FieldHasSuffix(FieldIps, v))
}

// IpsEqualFold applies the EqualFold predicate on the "ips" field.
func IpsEqualFold(v string) predicate.Session {
	return predicate.Session(sql.FieldEqualFold(FieldIps, v))
}

// IpsContainsFold applies the ContainsFold predicate on the "ips" field.
func IpsContainsFold(v string) predicate.Session {
	return predicate.Session(sql.FieldContainsFold(FieldIps, v))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UsersTable, UsersColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.Session
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		step := newUsersStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.Session
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Session) predicate.Session {
	return predicate.Session(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Session) predicate.Session {
	return predicate.Session(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Session) predicate.Session {
	return predicate.Session(sql.NotPredicates(p))
}
