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
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/orgmembership"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// OrgMembershipUpdate is the builder for updating OrgMembership entities.
type OrgMembershipUpdate struct {
	config
	hooks    []Hook
	mutation *OrgMembershipMutation
}

// Where appends a list predicates to the OrgMembershipUpdate builder.
func (omu *OrgMembershipUpdate) Where(ps ...predicate.OrgMembership) *OrgMembershipUpdate {
	omu.mutation.Where(ps...)
	return omu
}

// SetUpdatedAt sets the "updated_at" field.
func (omu *OrgMembershipUpdate) SetUpdatedAt(t time.Time) *OrgMembershipUpdate {
	omu.mutation.SetUpdatedAt(t)
	return omu
}

// SetUpdatedBy sets the "updated_by" field.
func (omu *OrgMembershipUpdate) SetUpdatedBy(s string) *OrgMembershipUpdate {
	omu.mutation.SetUpdatedBy(s)
	return omu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (omu *OrgMembershipUpdate) SetNillableUpdatedBy(s *string) *OrgMembershipUpdate {
	if s != nil {
		omu.SetUpdatedBy(*s)
	}
	return omu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (omu *OrgMembershipUpdate) ClearUpdatedBy() *OrgMembershipUpdate {
	omu.mutation.ClearUpdatedBy()
	return omu
}

// SetDeletedAt sets the "deleted_at" field.
func (omu *OrgMembershipUpdate) SetDeletedAt(t time.Time) *OrgMembershipUpdate {
	omu.mutation.SetDeletedAt(t)
	return omu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (omu *OrgMembershipUpdate) SetNillableDeletedAt(t *time.Time) *OrgMembershipUpdate {
	if t != nil {
		omu.SetDeletedAt(*t)
	}
	return omu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (omu *OrgMembershipUpdate) ClearDeletedAt() *OrgMembershipUpdate {
	omu.mutation.ClearDeletedAt()
	return omu
}

// SetDeletedBy sets the "deleted_by" field.
func (omu *OrgMembershipUpdate) SetDeletedBy(s string) *OrgMembershipUpdate {
	omu.mutation.SetDeletedBy(s)
	return omu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (omu *OrgMembershipUpdate) SetNillableDeletedBy(s *string) *OrgMembershipUpdate {
	if s != nil {
		omu.SetDeletedBy(*s)
	}
	return omu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (omu *OrgMembershipUpdate) ClearDeletedBy() *OrgMembershipUpdate {
	omu.mutation.ClearDeletedBy()
	return omu
}

// SetRole sets the "role" field.
func (omu *OrgMembershipUpdate) SetRole(e enums.Role) *OrgMembershipUpdate {
	omu.mutation.SetRole(e)
	return omu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (omu *OrgMembershipUpdate) SetNillableRole(e *enums.Role) *OrgMembershipUpdate {
	if e != nil {
		omu.SetRole(*e)
	}
	return omu
}

// Mutation returns the OrgMembershipMutation object of the builder.
func (omu *OrgMembershipUpdate) Mutation() *OrgMembershipMutation {
	return omu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (omu *OrgMembershipUpdate) Save(ctx context.Context) (int, error) {
	if err := omu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, omu.sqlSave, omu.mutation, omu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (omu *OrgMembershipUpdate) SaveX(ctx context.Context) int {
	affected, err := omu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (omu *OrgMembershipUpdate) Exec(ctx context.Context) error {
	_, err := omu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (omu *OrgMembershipUpdate) ExecX(ctx context.Context) {
	if err := omu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (omu *OrgMembershipUpdate) defaults() error {
	if _, ok := omu.mutation.UpdatedAt(); !ok {
		if orgmembership.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized orgmembership.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := orgmembership.UpdateDefaultUpdatedAt()
		omu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (omu *OrgMembershipUpdate) check() error {
	if v, ok := omu.mutation.Role(); ok {
		if err := orgmembership.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`generated: validator failed for field "OrgMembership.role": %w`, err)}
		}
	}
	if _, ok := omu.mutation.OrgID(); omu.mutation.OrgCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "OrgMembership.org"`)
	}
	if _, ok := omu.mutation.UserID(); omu.mutation.UserCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "OrgMembership.user"`)
	}
	return nil
}

func (omu *OrgMembershipUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := omu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(orgmembership.Table, orgmembership.Columns, sqlgraph.NewFieldSpec(orgmembership.FieldID, field.TypeString))
	if ps := omu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := omu.mutation.UpdatedAt(); ok {
		_spec.SetField(orgmembership.FieldUpdatedAt, field.TypeTime, value)
	}
	if omu.mutation.CreatedByCleared() {
		_spec.ClearField(orgmembership.FieldCreatedBy, field.TypeString)
	}
	if value, ok := omu.mutation.UpdatedBy(); ok {
		_spec.SetField(orgmembership.FieldUpdatedBy, field.TypeString, value)
	}
	if omu.mutation.UpdatedByCleared() {
		_spec.ClearField(orgmembership.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := omu.mutation.DeletedAt(); ok {
		_spec.SetField(orgmembership.FieldDeletedAt, field.TypeTime, value)
	}
	if omu.mutation.DeletedAtCleared() {
		_spec.ClearField(orgmembership.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := omu.mutation.DeletedBy(); ok {
		_spec.SetField(orgmembership.FieldDeletedBy, field.TypeString, value)
	}
	if omu.mutation.DeletedByCleared() {
		_spec.ClearField(orgmembership.FieldDeletedBy, field.TypeString)
	}
	if value, ok := omu.mutation.Role(); ok {
		_spec.SetField(orgmembership.FieldRole, field.TypeEnum, value)
	}
	_spec.Node.Schema = omu.schemaConfig.OrgMembership
	ctx = internal.NewSchemaConfigContext(ctx, omu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, omu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orgmembership.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	omu.mutation.done = true
	return n, nil
}

// OrgMembershipUpdateOne is the builder for updating a single OrgMembership entity.
type OrgMembershipUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrgMembershipMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (omuo *OrgMembershipUpdateOne) SetUpdatedAt(t time.Time) *OrgMembershipUpdateOne {
	omuo.mutation.SetUpdatedAt(t)
	return omuo
}

// SetUpdatedBy sets the "updated_by" field.
func (omuo *OrgMembershipUpdateOne) SetUpdatedBy(s string) *OrgMembershipUpdateOne {
	omuo.mutation.SetUpdatedBy(s)
	return omuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (omuo *OrgMembershipUpdateOne) SetNillableUpdatedBy(s *string) *OrgMembershipUpdateOne {
	if s != nil {
		omuo.SetUpdatedBy(*s)
	}
	return omuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (omuo *OrgMembershipUpdateOne) ClearUpdatedBy() *OrgMembershipUpdateOne {
	omuo.mutation.ClearUpdatedBy()
	return omuo
}

// SetDeletedAt sets the "deleted_at" field.
func (omuo *OrgMembershipUpdateOne) SetDeletedAt(t time.Time) *OrgMembershipUpdateOne {
	omuo.mutation.SetDeletedAt(t)
	return omuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (omuo *OrgMembershipUpdateOne) SetNillableDeletedAt(t *time.Time) *OrgMembershipUpdateOne {
	if t != nil {
		omuo.SetDeletedAt(*t)
	}
	return omuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (omuo *OrgMembershipUpdateOne) ClearDeletedAt() *OrgMembershipUpdateOne {
	omuo.mutation.ClearDeletedAt()
	return omuo
}

// SetDeletedBy sets the "deleted_by" field.
func (omuo *OrgMembershipUpdateOne) SetDeletedBy(s string) *OrgMembershipUpdateOne {
	omuo.mutation.SetDeletedBy(s)
	return omuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (omuo *OrgMembershipUpdateOne) SetNillableDeletedBy(s *string) *OrgMembershipUpdateOne {
	if s != nil {
		omuo.SetDeletedBy(*s)
	}
	return omuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (omuo *OrgMembershipUpdateOne) ClearDeletedBy() *OrgMembershipUpdateOne {
	omuo.mutation.ClearDeletedBy()
	return omuo
}

// SetRole sets the "role" field.
func (omuo *OrgMembershipUpdateOne) SetRole(e enums.Role) *OrgMembershipUpdateOne {
	omuo.mutation.SetRole(e)
	return omuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (omuo *OrgMembershipUpdateOne) SetNillableRole(e *enums.Role) *OrgMembershipUpdateOne {
	if e != nil {
		omuo.SetRole(*e)
	}
	return omuo
}

// Mutation returns the OrgMembershipMutation object of the builder.
func (omuo *OrgMembershipUpdateOne) Mutation() *OrgMembershipMutation {
	return omuo.mutation
}

// Where appends a list predicates to the OrgMembershipUpdate builder.
func (omuo *OrgMembershipUpdateOne) Where(ps ...predicate.OrgMembership) *OrgMembershipUpdateOne {
	omuo.mutation.Where(ps...)
	return omuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (omuo *OrgMembershipUpdateOne) Select(field string, fields ...string) *OrgMembershipUpdateOne {
	omuo.fields = append([]string{field}, fields...)
	return omuo
}

// Save executes the query and returns the updated OrgMembership entity.
func (omuo *OrgMembershipUpdateOne) Save(ctx context.Context) (*OrgMembership, error) {
	if err := omuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, omuo.sqlSave, omuo.mutation, omuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (omuo *OrgMembershipUpdateOne) SaveX(ctx context.Context) *OrgMembership {
	node, err := omuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (omuo *OrgMembershipUpdateOne) Exec(ctx context.Context) error {
	_, err := omuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (omuo *OrgMembershipUpdateOne) ExecX(ctx context.Context) {
	if err := omuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (omuo *OrgMembershipUpdateOne) defaults() error {
	if _, ok := omuo.mutation.UpdatedAt(); !ok {
		if orgmembership.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized orgmembership.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := orgmembership.UpdateDefaultUpdatedAt()
		omuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (omuo *OrgMembershipUpdateOne) check() error {
	if v, ok := omuo.mutation.Role(); ok {
		if err := orgmembership.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`generated: validator failed for field "OrgMembership.role": %w`, err)}
		}
	}
	if _, ok := omuo.mutation.OrgID(); omuo.mutation.OrgCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "OrgMembership.org"`)
	}
	if _, ok := omuo.mutation.UserID(); omuo.mutation.UserCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "OrgMembership.user"`)
	}
	return nil
}

func (omuo *OrgMembershipUpdateOne) sqlSave(ctx context.Context) (_node *OrgMembership, err error) {
	if err := omuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(orgmembership.Table, orgmembership.Columns, sqlgraph.NewFieldSpec(orgmembership.FieldID, field.TypeString))
	id, ok := omuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "OrgMembership.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := omuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orgmembership.FieldID)
		for _, f := range fields {
			if !orgmembership.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != orgmembership.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := omuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := omuo.mutation.UpdatedAt(); ok {
		_spec.SetField(orgmembership.FieldUpdatedAt, field.TypeTime, value)
	}
	if omuo.mutation.CreatedByCleared() {
		_spec.ClearField(orgmembership.FieldCreatedBy, field.TypeString)
	}
	if value, ok := omuo.mutation.UpdatedBy(); ok {
		_spec.SetField(orgmembership.FieldUpdatedBy, field.TypeString, value)
	}
	if omuo.mutation.UpdatedByCleared() {
		_spec.ClearField(orgmembership.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := omuo.mutation.DeletedAt(); ok {
		_spec.SetField(orgmembership.FieldDeletedAt, field.TypeTime, value)
	}
	if omuo.mutation.DeletedAtCleared() {
		_spec.ClearField(orgmembership.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := omuo.mutation.DeletedBy(); ok {
		_spec.SetField(orgmembership.FieldDeletedBy, field.TypeString, value)
	}
	if omuo.mutation.DeletedByCleared() {
		_spec.ClearField(orgmembership.FieldDeletedBy, field.TypeString)
	}
	if value, ok := omuo.mutation.Role(); ok {
		_spec.SetField(orgmembership.FieldRole, field.TypeEnum, value)
	}
	_spec.Node.Schema = omuo.schemaConfig.OrgMembership
	ctx = internal.NewSchemaConfigContext(ctx, omuo.schemaConfig)
	_node = &OrgMembership{config: omuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, omuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orgmembership.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	omuo.mutation.done = true
	return _node, nil
}