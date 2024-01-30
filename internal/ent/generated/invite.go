// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/invite"
	"github.com/datumforge/datum/internal/ent/generated/organization"
)

// Invite is the model entity for the Invite schema.
type Invite struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy string `json:"deleted_by,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	OwnerID string `json:"owner_id,omitempty"`
	// the invitation token sent to the user via email which should only be provided to the /verify endpoint + handler
	Token string `json:"-"`
	// the expiration date of the invitation token which defaults to 14 days in the future from creation
	Expires *time.Time `json:"expires,omitempty"`
	// the email used as input to generate the invitation token and is the destination person the invitation is sent to who is required to accept to join the organization
	Recipient string `json:"recipient,omitempty"`
	// the status of the invitation
	Status enums.InviteStatus `json:"status,omitempty"`
	// Role holds the value of the "role" field.
	Role enums.Role `json:"role,omitempty"`
	// the number of attempts made to perform email send of the invitation, maximum of 5
	SendAttempts int `json:"send_attempts,omitempty"`
	// the user who initiated the invitation
	RequestorID string `json:"requestor_id,omitempty"`
	// the comparison secret to verify the token's signature
	Secret *[]byte `json:"-"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InviteQuery when eager-loading is set.
	Edges        InviteEdges `json:"edges"`
	selectValues sql.SelectValues
}

// InviteEdges holds the relations/edges for other nodes in the graph.
type InviteEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Organization `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InviteEdges) OwnerOrErr() (*Organization, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: organization.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Invite) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case invite.FieldSecret:
			values[i] = new([]byte)
		case invite.FieldSendAttempts:
			values[i] = new(sql.NullInt64)
		case invite.FieldID, invite.FieldCreatedBy, invite.FieldUpdatedBy, invite.FieldDeletedBy, invite.FieldOwnerID, invite.FieldToken, invite.FieldRecipient, invite.FieldStatus, invite.FieldRole, invite.FieldRequestorID:
			values[i] = new(sql.NullString)
		case invite.FieldCreatedAt, invite.FieldUpdatedAt, invite.FieldDeletedAt, invite.FieldExpires:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Invite fields.
func (i *Invite) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case invite.FieldID:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value.Valid {
				i.ID = value.String
			}
		case invite.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		case invite.FieldUpdatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[j])
			} else if value.Valid {
				i.UpdatedAt = value.Time
			}
		case invite.FieldCreatedBy:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[j])
			} else if value.Valid {
				i.CreatedBy = value.String
			}
		case invite.FieldUpdatedBy:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[j])
			} else if value.Valid {
				i.UpdatedBy = value.String
			}
		case invite.FieldDeletedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[j])
			} else if value.Valid {
				i.DeletedAt = value.Time
			}
		case invite.FieldDeletedBy:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[j])
			} else if value.Valid {
				i.DeletedBy = value.String
			}
		case invite.FieldOwnerID:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[j])
			} else if value.Valid {
				i.OwnerID = value.String
			}
		case invite.FieldToken:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[j])
			} else if value.Valid {
				i.Token = value.String
			}
		case invite.FieldExpires:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expires", values[j])
			} else if value.Valid {
				i.Expires = new(time.Time)
				*i.Expires = value.Time
			}
		case invite.FieldRecipient:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field recipient", values[j])
			} else if value.Valid {
				i.Recipient = value.String
			}
		case invite.FieldStatus:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[j])
			} else if value.Valid {
				i.Status = enums.InviteStatus(value.String)
			}
		case invite.FieldRole:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[j])
			} else if value.Valid {
				i.Role = enums.Role(value.String)
			}
		case invite.FieldSendAttempts:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field send_attempts", values[j])
			} else if value.Valid {
				i.SendAttempts = int(value.Int64)
			}
		case invite.FieldRequestorID:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field requestor_id", values[j])
			} else if value.Valid {
				i.RequestorID = value.String
			}
		case invite.FieldSecret:
			if value, ok := values[j].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field secret", values[j])
			} else if value != nil {
				i.Secret = value
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Invite.
// This includes values selected through modifiers, order, etc.
func (i *Invite) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Invite entity.
func (i *Invite) QueryOwner() *OrganizationQuery {
	return NewInviteClient(i.config).QueryOwner(i)
}

// Update returns a builder for updating this Invite.
// Note that you need to call Invite.Unwrap() before calling this method if this Invite
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Invite) Update() *InviteUpdateOne {
	return NewInviteClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Invite entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Invite) Unwrap() *Invite {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("generated: Invite is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Invite) String() string {
	var builder strings.Builder
	builder.WriteString("Invite(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(i.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(i.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(i.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(i.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(i.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(i.OwnerID)
	builder.WriteString(", ")
	builder.WriteString("token=<sensitive>")
	builder.WriteString(", ")
	if v := i.Expires; v != nil {
		builder.WriteString("expires=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("recipient=")
	builder.WriteString(i.Recipient)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", i.Status))
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", i.Role))
	builder.WriteString(", ")
	builder.WriteString("send_attempts=")
	builder.WriteString(fmt.Sprintf("%v", i.SendAttempts))
	builder.WriteString(", ")
	builder.WriteString("requestor_id=")
	builder.WriteString(i.RequestorID)
	builder.WriteString(", ")
	builder.WriteString("secret=<sensitive>")
	builder.WriteByte(')')
	return builder.String()
}

// Invites is a parsable slice of Invite.
type Invites []*Invite
