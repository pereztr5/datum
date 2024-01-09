// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/session"
	"github.com/datumforge/datum/internal/ent/generated/user"
)

// SessionCreate is the builder for creating a Session entity.
type SessionCreate struct {
	config
	mutation *SessionMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sc *SessionCreate) SetCreatedAt(t time.Time) *SessionCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SessionCreate) SetNillableCreatedAt(t *time.Time) *SessionCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SessionCreate) SetUpdatedAt(t time.Time) *SessionCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SessionCreate) SetNillableUpdatedAt(t *time.Time) *SessionCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetCreatedBy sets the "created_by" field.
func (sc *SessionCreate) SetCreatedBy(s string) *SessionCreate {
	sc.mutation.SetCreatedBy(s)
	return sc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (sc *SessionCreate) SetNillableCreatedBy(s *string) *SessionCreate {
	if s != nil {
		sc.SetCreatedBy(*s)
	}
	return sc
}

// SetUpdatedBy sets the "updated_by" field.
func (sc *SessionCreate) SetUpdatedBy(s string) *SessionCreate {
	sc.mutation.SetUpdatedBy(s)
	return sc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (sc *SessionCreate) SetNillableUpdatedBy(s *string) *SessionCreate {
	if s != nil {
		sc.SetUpdatedBy(*s)
	}
	return sc
}

// SetOwnerID sets the "owner_id" field.
func (sc *SessionCreate) SetOwnerID(s string) *SessionCreate {
	sc.mutation.SetOwnerID(s)
	return sc
}

// SetSessionToken sets the "session_token" field.
func (sc *SessionCreate) SetSessionToken(s string) *SessionCreate {
	sc.mutation.SetSessionToken(s)
	return sc
}

// SetIssuedAt sets the "issued_at" field.
func (sc *SessionCreate) SetIssuedAt(t time.Time) *SessionCreate {
	sc.mutation.SetIssuedAt(t)
	return sc
}

// SetExpiresAt sets the "expires_at" field.
func (sc *SessionCreate) SetExpiresAt(t time.Time) *SessionCreate {
	sc.mutation.SetExpiresAt(t)
	return sc
}

// SetOrganizationID sets the "organization_id" field.
func (sc *SessionCreate) SetOrganizationID(s string) *SessionCreate {
	sc.mutation.SetOrganizationID(s)
	return sc
}

// SetID sets the "id" field.
func (sc *SessionCreate) SetID(s string) *SessionCreate {
	sc.mutation.SetID(s)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *SessionCreate) SetNillableID(s *string) *SessionCreate {
	if s != nil {
		sc.SetID(*s)
	}
	return sc
}

// SetOwner sets the "owner" edge to the User entity.
func (sc *SessionCreate) SetOwner(u *User) *SessionCreate {
	return sc.SetOwnerID(u.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (sc *SessionCreate) Mutation() *SessionMutation {
	return sc.mutation
}

// Save creates the Session in the database.
func (sc *SessionCreate) Save(ctx context.Context) (*Session, error) {
	if err := sc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SessionCreate) SaveX(ctx context.Context) *Session {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SessionCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SessionCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SessionCreate) defaults() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		if session.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized session.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := session.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		if session.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized session.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := session.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		if session.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized session.DefaultID (forgotten import generated/runtime?)")
		}
		v := session.DefaultID()
		sc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sc *SessionCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "Session.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "Session.updated_at"`)}
	}
	if _, ok := sc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`generated: missing required field "Session.owner_id"`)}
	}
	if _, ok := sc.mutation.SessionToken(); !ok {
		return &ValidationError{Name: "session_token", err: errors.New(`generated: missing required field "Session.session_token"`)}
	}
	if _, ok := sc.mutation.IssuedAt(); !ok {
		return &ValidationError{Name: "issued_at", err: errors.New(`generated: missing required field "Session.issued_at"`)}
	}
	if _, ok := sc.mutation.ExpiresAt(); !ok {
		return &ValidationError{Name: "expires_at", err: errors.New(`generated: missing required field "Session.expires_at"`)}
	}
	if _, ok := sc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`generated: missing required field "Session.organization_id"`)}
	}
	if _, ok := sc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`generated: missing required edge "Session.owner"`)}
	}
	return nil
}

func (sc *SessionCreate) sqlSave(ctx context.Context) (*Session, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Session.ID type: %T", _spec.ID.Value)
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SessionCreate) createSpec() (*Session, *sqlgraph.CreateSpec) {
	var (
		_node = &Session{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(session.Table, sqlgraph.NewFieldSpec(session.FieldID, field.TypeString))
	)
	_spec.Schema = sc.schemaConfig.Session
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(session.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(session.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.CreatedBy(); ok {
		_spec.SetField(session.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := sc.mutation.UpdatedBy(); ok {
		_spec.SetField(session.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := sc.mutation.SessionToken(); ok {
		_spec.SetField(session.FieldSessionToken, field.TypeString, value)
		_node.SessionToken = value
	}
	if value, ok := sc.mutation.IssuedAt(); ok {
		_spec.SetField(session.FieldIssuedAt, field.TypeTime, value)
		_node.IssuedAt = value
	}
	if value, ok := sc.mutation.ExpiresAt(); ok {
		_spec.SetField(session.FieldExpiresAt, field.TypeTime, value)
		_node.ExpiresAt = value
	}
	if value, ok := sc.mutation.OrganizationID(); ok {
		_spec.SetField(session.FieldOrganizationID, field.TypeString, value)
		_node.OrganizationID = value
	}
	if nodes := sc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.OwnerTable,
			Columns: []string{session.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = sc.schemaConfig.Session
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SessionCreateBulk is the builder for creating many Session entities in bulk.
type SessionCreateBulk struct {
	config
	err      error
	builders []*SessionCreate
}

// Save creates the Session entities in the database.
func (scb *SessionCreateBulk) Save(ctx context.Context) ([]*Session, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Session, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SessionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SessionCreateBulk) SaveX(ctx context.Context) []*Session {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SessionCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SessionCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
