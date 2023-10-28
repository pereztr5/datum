// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/membership"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/google/uuid"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// MembershipUpdate is the builder for updating Membership entities.
type MembershipUpdate struct {
	config
	hooks    []Hook
	mutation *MembershipMutation
}

// Where appends a list predicates to the MembershipUpdate builder.
func (mu *MembershipUpdate) Where(ps ...predicate.Membership) *MembershipUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MembershipUpdate) SetUpdatedAt(t time.Time) *MembershipUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetCreatedBy sets the "created_by" field.
func (mu *MembershipUpdate) SetCreatedBy(i int) *MembershipUpdate {
	mu.mutation.ResetCreatedBy()
	mu.mutation.SetCreatedBy(i)
	return mu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mu *MembershipUpdate) SetNillableCreatedBy(i *int) *MembershipUpdate {
	if i != nil {
		mu.SetCreatedBy(*i)
	}
	return mu
}

// AddCreatedBy adds i to the "created_by" field.
func (mu *MembershipUpdate) AddCreatedBy(i int) *MembershipUpdate {
	mu.mutation.AddCreatedBy(i)
	return mu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (mu *MembershipUpdate) ClearCreatedBy() *MembershipUpdate {
	mu.mutation.ClearCreatedBy()
	return mu
}

// SetUpdatedBy sets the "updated_by" field.
func (mu *MembershipUpdate) SetUpdatedBy(i int) *MembershipUpdate {
	mu.mutation.ResetUpdatedBy()
	mu.mutation.SetUpdatedBy(i)
	return mu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mu *MembershipUpdate) SetNillableUpdatedBy(i *int) *MembershipUpdate {
	if i != nil {
		mu.SetUpdatedBy(*i)
	}
	return mu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (mu *MembershipUpdate) AddUpdatedBy(i int) *MembershipUpdate {
	mu.mutation.AddUpdatedBy(i)
	return mu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (mu *MembershipUpdate) ClearUpdatedBy() *MembershipUpdate {
	mu.mutation.ClearUpdatedBy()
	return mu
}

// SetCurrent sets the "current" field.
func (mu *MembershipUpdate) SetCurrent(b bool) *MembershipUpdate {
	mu.mutation.SetCurrent(b)
	return mu
}

// SetNillableCurrent sets the "current" field if the given value is not nil.
func (mu *MembershipUpdate) SetNillableCurrent(b *bool) *MembershipUpdate {
	if b != nil {
		mu.SetCurrent(*b)
	}
	return mu
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (mu *MembershipUpdate) SetOrganizationID(id uuid.UUID) *MembershipUpdate {
	mu.mutation.SetOrganizationID(id)
	return mu
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (mu *MembershipUpdate) SetOrganization(o *Organization) *MembershipUpdate {
	return mu.SetOrganizationID(o.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (mu *MembershipUpdate) SetUserID(id uuid.UUID) *MembershipUpdate {
	mu.mutation.SetUserID(id)
	return mu
}

// SetUser sets the "user" edge to the User entity.
func (mu *MembershipUpdate) SetUser(u *User) *MembershipUpdate {
	return mu.SetUserID(u.ID)
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (mu *MembershipUpdate) SetGroupID(id uuid.UUID) *MembershipUpdate {
	mu.mutation.SetGroupID(id)
	return mu
}

// SetGroup sets the "group" edge to the Group entity.
func (mu *MembershipUpdate) SetGroup(g *Group) *MembershipUpdate {
	return mu.SetGroupID(g.ID)
}

// Mutation returns the MembershipMutation object of the builder.
func (mu *MembershipUpdate) Mutation() *MembershipMutation {
	return mu.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (mu *MembershipUpdate) ClearOrganization() *MembershipUpdate {
	mu.mutation.ClearOrganization()
	return mu
}

// ClearUser clears the "user" edge to the User entity.
func (mu *MembershipUpdate) ClearUser() *MembershipUpdate {
	mu.mutation.ClearUser()
	return mu
}

// ClearGroup clears the "group" edge to the Group entity.
func (mu *MembershipUpdate) ClearGroup() *MembershipUpdate {
	mu.mutation.ClearGroup()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MembershipUpdate) Save(ctx context.Context) (int, error) {
	if err := mu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MembershipUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MembershipUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MembershipUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MembershipUpdate) defaults() error {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		if membership.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized membership.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := membership.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (mu *MembershipUpdate) check() error {
	if _, ok := mu.mutation.OrganizationID(); mu.mutation.OrganizationCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Membership.organization"`)
	}
	if _, ok := mu.mutation.UserID(); mu.mutation.UserCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Membership.user"`)
	}
	if _, ok := mu.mutation.GroupID(); mu.mutation.GroupCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Membership.group"`)
	}
	return nil
}

func (mu *MembershipUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(membership.Table, membership.Columns, sqlgraph.NewFieldSpec(membership.FieldID, field.TypeUUID))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.SetField(membership.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.CreatedBy(); ok {
		_spec.SetField(membership.FieldCreatedBy, field.TypeInt, value)
	}
	if value, ok := mu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(membership.FieldCreatedBy, field.TypeInt, value)
	}
	if mu.mutation.CreatedByCleared() {
		_spec.ClearField(membership.FieldCreatedBy, field.TypeInt)
	}
	if value, ok := mu.mutation.UpdatedBy(); ok {
		_spec.SetField(membership.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := mu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(membership.FieldUpdatedBy, field.TypeInt, value)
	}
	if mu.mutation.UpdatedByCleared() {
		_spec.ClearField(membership.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := mu.mutation.Current(); ok {
		_spec.SetField(membership.FieldCurrent, field.TypeBool, value)
	}
	if mu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.OrganizationTable,
			Columns: []string{membership.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = mu.schemaConfig.Membership
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.OrganizationTable,
			Columns: []string{membership.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = mu.schemaConfig.Membership
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.UserTable,
			Columns: []string{membership.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = mu.schemaConfig.Membership
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.UserTable,
			Columns: []string{membership.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = mu.schemaConfig.Membership
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.GroupTable,
			Columns: []string{membership.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = mu.schemaConfig.Membership
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.GroupTable,
			Columns: []string{membership.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = mu.schemaConfig.Membership
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = mu.schemaConfig.Membership
	ctx = internal.NewSchemaConfigContext(ctx, mu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{membership.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MembershipUpdateOne is the builder for updating a single Membership entity.
type MembershipUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MembershipMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MembershipUpdateOne) SetUpdatedAt(t time.Time) *MembershipUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetCreatedBy sets the "created_by" field.
func (muo *MembershipUpdateOne) SetCreatedBy(i int) *MembershipUpdateOne {
	muo.mutation.ResetCreatedBy()
	muo.mutation.SetCreatedBy(i)
	return muo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (muo *MembershipUpdateOne) SetNillableCreatedBy(i *int) *MembershipUpdateOne {
	if i != nil {
		muo.SetCreatedBy(*i)
	}
	return muo
}

// AddCreatedBy adds i to the "created_by" field.
func (muo *MembershipUpdateOne) AddCreatedBy(i int) *MembershipUpdateOne {
	muo.mutation.AddCreatedBy(i)
	return muo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (muo *MembershipUpdateOne) ClearCreatedBy() *MembershipUpdateOne {
	muo.mutation.ClearCreatedBy()
	return muo
}

// SetUpdatedBy sets the "updated_by" field.
func (muo *MembershipUpdateOne) SetUpdatedBy(i int) *MembershipUpdateOne {
	muo.mutation.ResetUpdatedBy()
	muo.mutation.SetUpdatedBy(i)
	return muo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (muo *MembershipUpdateOne) SetNillableUpdatedBy(i *int) *MembershipUpdateOne {
	if i != nil {
		muo.SetUpdatedBy(*i)
	}
	return muo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (muo *MembershipUpdateOne) AddUpdatedBy(i int) *MembershipUpdateOne {
	muo.mutation.AddUpdatedBy(i)
	return muo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (muo *MembershipUpdateOne) ClearUpdatedBy() *MembershipUpdateOne {
	muo.mutation.ClearUpdatedBy()
	return muo
}

// SetCurrent sets the "current" field.
func (muo *MembershipUpdateOne) SetCurrent(b bool) *MembershipUpdateOne {
	muo.mutation.SetCurrent(b)
	return muo
}

// SetNillableCurrent sets the "current" field if the given value is not nil.
func (muo *MembershipUpdateOne) SetNillableCurrent(b *bool) *MembershipUpdateOne {
	if b != nil {
		muo.SetCurrent(*b)
	}
	return muo
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (muo *MembershipUpdateOne) SetOrganizationID(id uuid.UUID) *MembershipUpdateOne {
	muo.mutation.SetOrganizationID(id)
	return muo
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (muo *MembershipUpdateOne) SetOrganization(o *Organization) *MembershipUpdateOne {
	return muo.SetOrganizationID(o.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (muo *MembershipUpdateOne) SetUserID(id uuid.UUID) *MembershipUpdateOne {
	muo.mutation.SetUserID(id)
	return muo
}

// SetUser sets the "user" edge to the User entity.
func (muo *MembershipUpdateOne) SetUser(u *User) *MembershipUpdateOne {
	return muo.SetUserID(u.ID)
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (muo *MembershipUpdateOne) SetGroupID(id uuid.UUID) *MembershipUpdateOne {
	muo.mutation.SetGroupID(id)
	return muo
}

// SetGroup sets the "group" edge to the Group entity.
func (muo *MembershipUpdateOne) SetGroup(g *Group) *MembershipUpdateOne {
	return muo.SetGroupID(g.ID)
}

// Mutation returns the MembershipMutation object of the builder.
func (muo *MembershipUpdateOne) Mutation() *MembershipMutation {
	return muo.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (muo *MembershipUpdateOne) ClearOrganization() *MembershipUpdateOne {
	muo.mutation.ClearOrganization()
	return muo
}

// ClearUser clears the "user" edge to the User entity.
func (muo *MembershipUpdateOne) ClearUser() *MembershipUpdateOne {
	muo.mutation.ClearUser()
	return muo
}

// ClearGroup clears the "group" edge to the Group entity.
func (muo *MembershipUpdateOne) ClearGroup() *MembershipUpdateOne {
	muo.mutation.ClearGroup()
	return muo
}

// Where appends a list predicates to the MembershipUpdate builder.
func (muo *MembershipUpdateOne) Where(ps ...predicate.Membership) *MembershipUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MembershipUpdateOne) Select(field string, fields ...string) *MembershipUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Membership entity.
func (muo *MembershipUpdateOne) Save(ctx context.Context) (*Membership, error) {
	if err := muo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MembershipUpdateOne) SaveX(ctx context.Context) *Membership {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MembershipUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MembershipUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MembershipUpdateOne) defaults() error {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		if membership.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized membership.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := membership.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (muo *MembershipUpdateOne) check() error {
	if _, ok := muo.mutation.OrganizationID(); muo.mutation.OrganizationCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Membership.organization"`)
	}
	if _, ok := muo.mutation.UserID(); muo.mutation.UserCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Membership.user"`)
	}
	if _, ok := muo.mutation.GroupID(); muo.mutation.GroupCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Membership.group"`)
	}
	return nil
}

func (muo *MembershipUpdateOne) sqlSave(ctx context.Context) (_node *Membership, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(membership.Table, membership.Columns, sqlgraph.NewFieldSpec(membership.FieldID, field.TypeUUID))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Membership.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, membership.FieldID)
		for _, f := range fields {
			if !membership.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != membership.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.SetField(membership.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.CreatedBy(); ok {
		_spec.SetField(membership.FieldCreatedBy, field.TypeInt, value)
	}
	if value, ok := muo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(membership.FieldCreatedBy, field.TypeInt, value)
	}
	if muo.mutation.CreatedByCleared() {
		_spec.ClearField(membership.FieldCreatedBy, field.TypeInt)
	}
	if value, ok := muo.mutation.UpdatedBy(); ok {
		_spec.SetField(membership.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := muo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(membership.FieldUpdatedBy, field.TypeInt, value)
	}
	if muo.mutation.UpdatedByCleared() {
		_spec.ClearField(membership.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := muo.mutation.Current(); ok {
		_spec.SetField(membership.FieldCurrent, field.TypeBool, value)
	}
	if muo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.OrganizationTable,
			Columns: []string{membership.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = muo.schemaConfig.Membership
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.OrganizationTable,
			Columns: []string{membership.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = muo.schemaConfig.Membership
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.UserTable,
			Columns: []string{membership.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = muo.schemaConfig.Membership
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.UserTable,
			Columns: []string{membership.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = muo.schemaConfig.Membership
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.GroupTable,
			Columns: []string{membership.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = muo.schemaConfig.Membership
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.GroupTable,
			Columns: []string{membership.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = muo.schemaConfig.Membership
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = muo.schemaConfig.Membership
	ctx = internal.NewSchemaConfigContext(ctx, muo.schemaConfig)
	_node = &Membership{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{membership.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
