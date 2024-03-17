// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/subscriber"
)

// Subscriber is the model entity for the Subscriber schema.
type Subscriber struct {
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
	// email address of the subscriber
	Email string `json:"email,omitempty"`
	// phone number of the subscriber
	PhoneNumber string `json:"phone_number,omitempty"`
	// indicates if the email address has been verified
	VerifiedEmail bool `json:"verified_email,omitempty"`
	// indicates if the phone number has been verified
	VerifiedPhone bool `json:"verified_phone,omitempty"`
	// indicates if the subscriber is active or not, active users will have at least one verified contact method
	Active bool `json:"active,omitempty"`
	// the verification token sent to the user via email which should only be provided to the /subscribe endpoint + handler
	Token string `json:"token,omitempty"`
	// the ttl of the verification token which defaults to 7 days
	TTL *time.Time `json:"ttl,omitempty"`
	// the comparison secret to verify the token's signature
	Secret *[]byte `json:"secret,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubscriberQuery when eager-loading is set.
	Edges        SubscriberEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SubscriberEdges holds the relations/edges for other nodes in the graph.
type SubscriberEdges struct {
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
func (e SubscriberEdges) OwnerOrErr() (*Organization, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Subscriber) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case subscriber.FieldSecret:
			values[i] = new([]byte)
		case subscriber.FieldVerifiedEmail, subscriber.FieldVerifiedPhone, subscriber.FieldActive:
			values[i] = new(sql.NullBool)
		case subscriber.FieldID, subscriber.FieldCreatedBy, subscriber.FieldUpdatedBy, subscriber.FieldDeletedBy, subscriber.FieldOwnerID, subscriber.FieldEmail, subscriber.FieldPhoneNumber, subscriber.FieldToken:
			values[i] = new(sql.NullString)
		case subscriber.FieldCreatedAt, subscriber.FieldUpdatedAt, subscriber.FieldDeletedAt, subscriber.FieldTTL:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Subscriber fields.
func (s *Subscriber) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subscriber.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				s.ID = value.String
			}
		case subscriber.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case subscriber.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case subscriber.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				s.CreatedBy = value.String
			}
		case subscriber.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				s.UpdatedBy = value.String
			}
		case subscriber.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				s.DeletedAt = value.Time
			}
		case subscriber.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				s.DeletedBy = value.String
			}
		case subscriber.FieldOwnerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				s.OwnerID = value.String
			}
		case subscriber.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				s.Email = value.String
			}
		case subscriber.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number", values[i])
			} else if value.Valid {
				s.PhoneNumber = value.String
			}
		case subscriber.FieldVerifiedEmail:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field verified_email", values[i])
			} else if value.Valid {
				s.VerifiedEmail = value.Bool
			}
		case subscriber.FieldVerifiedPhone:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field verified_phone", values[i])
			} else if value.Valid {
				s.VerifiedPhone = value.Bool
			}
		case subscriber.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				s.Active = value.Bool
			}
		case subscriber.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				s.Token = value.String
			}
		case subscriber.FieldTTL:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ttl", values[i])
			} else if value.Valid {
				s.TTL = new(time.Time)
				*s.TTL = value.Time
			}
		case subscriber.FieldSecret:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field secret", values[i])
			} else if value != nil {
				s.Secret = value
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Subscriber.
// This includes values selected through modifiers, order, etc.
func (s *Subscriber) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Subscriber entity.
func (s *Subscriber) QueryOwner() *OrganizationQuery {
	return NewSubscriberClient(s.config).QueryOwner(s)
}

// Update returns a builder for updating this Subscriber.
// Note that you need to call Subscriber.Unwrap() before calling this method if this Subscriber
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Subscriber) Update() *SubscriberUpdateOne {
	return NewSubscriberClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Subscriber entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Subscriber) Unwrap() *Subscriber {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("generated: Subscriber is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Subscriber) String() string {
	var builder strings.Builder
	builder.WriteString("Subscriber(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(s.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(s.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(s.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(s.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(s.OwnerID)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(s.Email)
	builder.WriteString(", ")
	builder.WriteString("phone_number=")
	builder.WriteString(s.PhoneNumber)
	builder.WriteString(", ")
	builder.WriteString("verified_email=")
	builder.WriteString(fmt.Sprintf("%v", s.VerifiedEmail))
	builder.WriteString(", ")
	builder.WriteString("verified_phone=")
	builder.WriteString(fmt.Sprintf("%v", s.VerifiedPhone))
	builder.WriteString(", ")
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", s.Active))
	builder.WriteString(", ")
	builder.WriteString("token=")
	builder.WriteString(s.Token)
	builder.WriteString(", ")
	if v := s.TTL; v != nil {
		builder.WriteString("ttl=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := s.Secret; v != nil {
		builder.WriteString("secret=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Subscribers is a parsable slice of Subscriber.
type Subscribers []*Subscriber