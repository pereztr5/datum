// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/groupsettings"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/google/uuid"
)

// Group is the model entity for the Group schema.
type Group struct {
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
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// LogoURL holds the value of the "logo_url" field.
	LogoURL string `json:"logo_url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupQuery when eager-loading is set.
	Edges               GroupEdges `json:"edges"`
	organization_groups *uuid.UUID
	selectValues        sql.SelectValues
}

// GroupEdges holds the relations/edges for other nodes in the graph.
type GroupEdges struct {
	// Setting holds the value of the setting edge.
	Setting *GroupSettings `json:"setting,omitempty"`
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Owner holds the value of the owner edge.
	Owner *Organization `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedUsers map[string][]*User
}

// SettingOrErr returns the Setting value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupEdges) SettingOrErr() (*GroupSettings, error) {
	if e.loadedTypes[0] {
		if e.Setting == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: groupsettings.Label}
		}
		return e.Setting, nil
	}
	return nil, &NotLoadedError{edge: "setting"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupEdges) OwnerOrErr() (*Organization, error) {
	if e.loadedTypes[2] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: organization.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Group) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case group.FieldName, group.FieldDescription, group.FieldLogoURL:
			values[i] = new(sql.NullString)
		case group.FieldCreatedAt, group.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case group.FieldID, group.FieldCreatedBy, group.FieldUpdatedBy:
			values[i] = new(uuid.UUID)
		case group.ForeignKeys[0]: // organization_groups
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Group fields.
func (gr *Group) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case group.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				gr.ID = *value
			}
		case group.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gr.CreatedAt = value.Time
			}
		case group.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gr.UpdatedAt = value.Time
			}
		case group.FieldCreatedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value != nil {
				gr.CreatedBy = *value
			}
		case group.FieldUpdatedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value != nil {
				gr.UpdatedBy = *value
			}
		case group.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				gr.Name = value.String
			}
		case group.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				gr.Description = value.String
			}
		case group.FieldLogoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logo_url", values[i])
			} else if value.Valid {
				gr.LogoURL = value.String
			}
		case group.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field organization_groups", values[i])
			} else if value.Valid {
				gr.organization_groups = new(uuid.UUID)
				*gr.organization_groups = *value.S.(*uuid.UUID)
			}
		default:
			gr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Group.
// This includes values selected through modifiers, order, etc.
func (gr *Group) Value(name string) (ent.Value, error) {
	return gr.selectValues.Get(name)
}

// QuerySetting queries the "setting" edge of the Group entity.
func (gr *Group) QuerySetting() *GroupSettingsQuery {
	return NewGroupClient(gr.config).QuerySetting(gr)
}

// QueryUsers queries the "users" edge of the Group entity.
func (gr *Group) QueryUsers() *UserQuery {
	return NewGroupClient(gr.config).QueryUsers(gr)
}

// QueryOwner queries the "owner" edge of the Group entity.
func (gr *Group) QueryOwner() *OrganizationQuery {
	return NewGroupClient(gr.config).QueryOwner(gr)
}

// Update returns a builder for updating this Group.
// Note that you need to call Group.Unwrap() before calling this method if this Group
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *Group) Update() *GroupUpdateOne {
	return NewGroupClient(gr.config).UpdateOne(gr)
}

// Unwrap unwraps the Group entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gr *Group) Unwrap() *Group {
	_tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("generated: Group is not a transactional entity")
	}
	gr.config.driver = _tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *Group) String() string {
	var builder strings.Builder
	builder.WriteString("Group(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(gr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(gr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", gr.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", gr.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(gr.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(gr.Description)
	builder.WriteString(", ")
	builder.WriteString("logo_url=")
	builder.WriteString(gr.LogoURL)
	builder.WriteByte(')')
	return builder.String()
}

// NamedUsers returns the Users named value or an error if the edge was not
// loaded in eager-loading with this name.
func (gr *Group) NamedUsers(name string) ([]*User, error) {
	if gr.Edges.namedUsers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := gr.Edges.namedUsers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (gr *Group) appendNamedUsers(name string, edges ...*User) {
	if gr.Edges.namedUsers == nil {
		gr.Edges.namedUsers = make(map[string][]*User)
	}
	if len(edges) == 0 {
		gr.Edges.namedUsers[name] = []*User{}
	} else {
		gr.Edges.namedUsers[name] = append(gr.Edges.namedUsers[name], edges...)
	}
}

// Groups is a parsable slice of Group.
type Groups []*Group
