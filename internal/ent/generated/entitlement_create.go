// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/entitlement"
)

// EntitlementCreate is the builder for creating a Entitlement entity.
type EntitlementCreate struct {
	config
	mutation *EntitlementMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ec *EntitlementCreate) SetCreatedAt(t time.Time) *EntitlementCreate {
	ec.mutation.SetCreatedAt(t)
	return ec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ec *EntitlementCreate) SetNillableCreatedAt(t *time.Time) *EntitlementCreate {
	if t != nil {
		ec.SetCreatedAt(*t)
	}
	return ec
}

// SetUpdatedAt sets the "updated_at" field.
func (ec *EntitlementCreate) SetUpdatedAt(t time.Time) *EntitlementCreate {
	ec.mutation.SetUpdatedAt(t)
	return ec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ec *EntitlementCreate) SetNillableUpdatedAt(t *time.Time) *EntitlementCreate {
	if t != nil {
		ec.SetUpdatedAt(*t)
	}
	return ec
}

// SetCreatedBy sets the "created_by" field.
func (ec *EntitlementCreate) SetCreatedBy(s string) *EntitlementCreate {
	ec.mutation.SetCreatedBy(s)
	return ec
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ec *EntitlementCreate) SetNillableCreatedBy(s *string) *EntitlementCreate {
	if s != nil {
		ec.SetCreatedBy(*s)
	}
	return ec
}

// SetUpdatedBy sets the "updated_by" field.
func (ec *EntitlementCreate) SetUpdatedBy(s string) *EntitlementCreate {
	ec.mutation.SetUpdatedBy(s)
	return ec
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ec *EntitlementCreate) SetNillableUpdatedBy(s *string) *EntitlementCreate {
	if s != nil {
		ec.SetUpdatedBy(*s)
	}
	return ec
}

// SetTier sets the "tier" field.
func (ec *EntitlementCreate) SetTier(e entitlement.Tier) *EntitlementCreate {
	ec.mutation.SetTier(e)
	return ec
}

// SetNillableTier sets the "tier" field if the given value is not nil.
func (ec *EntitlementCreate) SetNillableTier(e *entitlement.Tier) *EntitlementCreate {
	if e != nil {
		ec.SetTier(*e)
	}
	return ec
}

// SetStripeCustomerID sets the "stripe_customer_id" field.
func (ec *EntitlementCreate) SetStripeCustomerID(s string) *EntitlementCreate {
	ec.mutation.SetStripeCustomerID(s)
	return ec
}

// SetNillableStripeCustomerID sets the "stripe_customer_id" field if the given value is not nil.
func (ec *EntitlementCreate) SetNillableStripeCustomerID(s *string) *EntitlementCreate {
	if s != nil {
		ec.SetStripeCustomerID(*s)
	}
	return ec
}

// SetStripeSubscriptionID sets the "stripe_subscription_id" field.
func (ec *EntitlementCreate) SetStripeSubscriptionID(s string) *EntitlementCreate {
	ec.mutation.SetStripeSubscriptionID(s)
	return ec
}

// SetNillableStripeSubscriptionID sets the "stripe_subscription_id" field if the given value is not nil.
func (ec *EntitlementCreate) SetNillableStripeSubscriptionID(s *string) *EntitlementCreate {
	if s != nil {
		ec.SetStripeSubscriptionID(*s)
	}
	return ec
}

// SetExpiresAt sets the "expires_at" field.
func (ec *EntitlementCreate) SetExpiresAt(t time.Time) *EntitlementCreate {
	ec.mutation.SetExpiresAt(t)
	return ec
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (ec *EntitlementCreate) SetNillableExpiresAt(t *time.Time) *EntitlementCreate {
	if t != nil {
		ec.SetExpiresAt(*t)
	}
	return ec
}

// SetCancelled sets the "cancelled" field.
func (ec *EntitlementCreate) SetCancelled(b bool) *EntitlementCreate {
	ec.mutation.SetCancelled(b)
	return ec
}

// SetNillableCancelled sets the "cancelled" field if the given value is not nil.
func (ec *EntitlementCreate) SetNillableCancelled(b *bool) *EntitlementCreate {
	if b != nil {
		ec.SetCancelled(*b)
	}
	return ec
}

// SetID sets the "id" field.
func (ec *EntitlementCreate) SetID(s string) *EntitlementCreate {
	ec.mutation.SetID(s)
	return ec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ec *EntitlementCreate) SetNillableID(s *string) *EntitlementCreate {
	if s != nil {
		ec.SetID(*s)
	}
	return ec
}

// Mutation returns the EntitlementMutation object of the builder.
func (ec *EntitlementCreate) Mutation() *EntitlementMutation {
	return ec.mutation
}

// Save creates the Entitlement in the database.
func (ec *EntitlementCreate) Save(ctx context.Context) (*Entitlement, error) {
	if err := ec.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EntitlementCreate) SaveX(ctx context.Context) *Entitlement {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EntitlementCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EntitlementCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EntitlementCreate) defaults() error {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		if entitlement.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized entitlement.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := entitlement.DefaultCreatedAt()
		ec.mutation.SetCreatedAt(v)
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		if entitlement.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized entitlement.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := entitlement.DefaultUpdatedAt()
		ec.mutation.SetUpdatedAt(v)
	}
	if _, ok := ec.mutation.Tier(); !ok {
		v := entitlement.DefaultTier
		ec.mutation.SetTier(v)
	}
	if _, ok := ec.mutation.Cancelled(); !ok {
		v := entitlement.DefaultCancelled
		ec.mutation.SetCancelled(v)
	}
	if _, ok := ec.mutation.ID(); !ok {
		if entitlement.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized entitlement.DefaultID (forgotten import generated/runtime?)")
		}
		v := entitlement.DefaultID()
		ec.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ec *EntitlementCreate) check() error {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "Entitlement.created_at"`)}
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "Entitlement.updated_at"`)}
	}
	if _, ok := ec.mutation.Tier(); !ok {
		return &ValidationError{Name: "tier", err: errors.New(`generated: missing required field "Entitlement.tier"`)}
	}
	if v, ok := ec.mutation.Tier(); ok {
		if err := entitlement.TierValidator(v); err != nil {
			return &ValidationError{Name: "tier", err: fmt.Errorf(`generated: validator failed for field "Entitlement.tier": %w`, err)}
		}
	}
	if _, ok := ec.mutation.Cancelled(); !ok {
		return &ValidationError{Name: "cancelled", err: errors.New(`generated: missing required field "Entitlement.cancelled"`)}
	}
	return nil
}

func (ec *EntitlementCreate) sqlSave(ctx context.Context) (*Entitlement, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Entitlement.ID type: %T", _spec.ID.Value)
		}
	}
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *EntitlementCreate) createSpec() (*Entitlement, *sqlgraph.CreateSpec) {
	var (
		_node = &Entitlement{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(entitlement.Table, sqlgraph.NewFieldSpec(entitlement.FieldID, field.TypeString))
	)
	_spec.Schema = ec.schemaConfig.Entitlement
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.SetField(entitlement.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ec.mutation.UpdatedAt(); ok {
		_spec.SetField(entitlement.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ec.mutation.CreatedBy(); ok {
		_spec.SetField(entitlement.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := ec.mutation.UpdatedBy(); ok {
		_spec.SetField(entitlement.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := ec.mutation.Tier(); ok {
		_spec.SetField(entitlement.FieldTier, field.TypeEnum, value)
		_node.Tier = value
	}
	if value, ok := ec.mutation.StripeCustomerID(); ok {
		_spec.SetField(entitlement.FieldStripeCustomerID, field.TypeString, value)
		_node.StripeCustomerID = value
	}
	if value, ok := ec.mutation.StripeSubscriptionID(); ok {
		_spec.SetField(entitlement.FieldStripeSubscriptionID, field.TypeString, value)
		_node.StripeSubscriptionID = value
	}
	if value, ok := ec.mutation.ExpiresAt(); ok {
		_spec.SetField(entitlement.FieldExpiresAt, field.TypeTime, value)
		_node.ExpiresAt = value
	}
	if value, ok := ec.mutation.Cancelled(); ok {
		_spec.SetField(entitlement.FieldCancelled, field.TypeBool, value)
		_node.Cancelled = value
	}
	return _node, _spec
}

// EntitlementCreateBulk is the builder for creating many Entitlement entities in bulk.
type EntitlementCreateBulk struct {
	config
	err      error
	builders []*EntitlementCreate
}

// Save creates the Entitlement entities in the database.
func (ecb *EntitlementCreateBulk) Save(ctx context.Context) ([]*Entitlement, error) {
	if ecb.err != nil {
		return nil, ecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Entitlement, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EntitlementMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EntitlementCreateBulk) SaveX(ctx context.Context) []*Entitlement {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EntitlementCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EntitlementCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}