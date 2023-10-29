// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/membership"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/google/uuid"
)

// Membership is the model entity for the Membership schema.
type Membership struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy uuid.UUID `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy uuid.UUID `json:"updated_by,omitempty"`
	// Current holds the value of the "current" field.
	Current bool `json:"current,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MembershipQuery when eager-loading is set.
	Edges                    MembershipEdges `json:"edges"`
	group_memberships        *uuid.UUID
	organization_memberships *uuid.UUID
	user_memberships         *uuid.UUID
	selectValues             sql.SelectValues
}

// MembershipEdges holds the relations/edges for other nodes in the graph.
type MembershipEdges struct {
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MembershipEdges) OrganizationOrErr() (*Organization, error) {
	if e.loadedTypes[0] {
		if e.Organization == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: organization.Label}
		}
		return e.Organization, nil
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MembershipEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MembershipEdges) GroupOrErr() (*Group, error) {
	if e.loadedTypes[2] {
		if e.Group == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Membership) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case membership.FieldCurrent:
			values[i] = new(sql.NullBool)
		case membership.FieldCreatedAt, membership.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case membership.FieldID, membership.FieldCreatedBy, membership.FieldUpdatedBy:
			values[i] = new(uuid.UUID)
		case membership.ForeignKeys[0]: // group_memberships
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case membership.ForeignKeys[1]: // organization_memberships
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case membership.ForeignKeys[2]: // user_memberships
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Membership fields.
func (m *Membership) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case membership.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				m.ID = *value
			}
		case membership.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case membership.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case membership.FieldCreatedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value != nil {
				m.CreatedBy = *value
			}
		case membership.FieldUpdatedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value != nil {
				m.UpdatedBy = *value
			}
		case membership.FieldCurrent:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field current", values[i])
			} else if value.Valid {
				m.Current = value.Bool
			}
		case membership.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field group_memberships", values[i])
			} else if value.Valid {
				m.group_memberships = new(uuid.UUID)
				*m.group_memberships = *value.S.(*uuid.UUID)
			}
		case membership.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field organization_memberships", values[i])
			} else if value.Valid {
				m.organization_memberships = new(uuid.UUID)
				*m.organization_memberships = *value.S.(*uuid.UUID)
			}
		case membership.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_memberships", values[i])
			} else if value.Valid {
				m.user_memberships = new(uuid.UUID)
				*m.user_memberships = *value.S.(*uuid.UUID)
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Membership.
// This includes values selected through modifiers, order, etc.
func (m *Membership) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryOrganization queries the "organization" edge of the Membership entity.
func (m *Membership) QueryOrganization() *OrganizationQuery {
	return NewMembershipClient(m.config).QueryOrganization(m)
}

// QueryUser queries the "user" edge of the Membership entity.
func (m *Membership) QueryUser() *UserQuery {
	return NewMembershipClient(m.config).QueryUser(m)
}

// QueryGroup queries the "group" edge of the Membership entity.
func (m *Membership) QueryGroup() *GroupQuery {
	return NewMembershipClient(m.config).QueryGroup(m)
}

// Update returns a builder for updating this Membership.
// Note that you need to call Membership.Unwrap() before calling this method if this Membership
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Membership) Update() *MembershipUpdateOne {
	return NewMembershipClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Membership entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Membership) Unwrap() *Membership {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("generated: Membership is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Membership) String() string {
	var builder strings.Builder
	builder.WriteString("Membership(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", m.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", m.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("current=")
	builder.WriteString(fmt.Sprintf("%v", m.Current))
	builder.WriteByte(')')
	return builder.String()
}

// Memberships is a parsable slice of Membership.
type Memberships []*Membership
