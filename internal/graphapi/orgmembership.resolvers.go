package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"errors"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
)

// CreateOrgMembership is the resolver for the createOrgMembership field.
func (r *mutationResolver) CreateOrgMembership(ctx context.Context, input generated.CreateOrgMembershipInput) (*OrgMembershipCreatePayload, error) {
	// check permissions if authz is enabled
	// if auth is disabled, policy decisions will be skipped
	if r.authDisabled {
		ctx = privacy.DecisionContext(ctx, privacy.Allow)
	} else {
		// setup view context
		v := viewer.UserViewer{
			OrgID: input.OrgID,
		}

		ctx = viewer.NewContext(ctx, v)
	}

	om, err := withTransactionalMutation(ctx).OrgMembership.Create().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			validationError := err.(*generated.ValidationError)

			r.logger.Debugw("validation error", "field", validationError.Name, "error", validationError.Error())

			return nil, validationError
		}

		if generated.IsConstraintError(err) {
			constraintError := err.(*generated.ConstraintError)

			r.logger.Debugw("constraint error", "error", constraintError.Error())

			return nil, constraintError
		}

		if errors.Is(err, privacy.Deny) {
			return nil, newPermissionDeniedError(ActionCreate, "org member")
		}

		r.logger.Errorw("failed to create org member", "error", err)

		return nil, err
	}

	return &OrgMembershipCreatePayload{OrgMembership: om}, nil
}

// UpdateOrgMembership is the resolver for the updateOrgMembership field.
func (r *mutationResolver) UpdateOrgMembership(ctx context.Context, id string, input generated.UpdateOrgMembershipInput) (*OrgMembershipUpdatePayload, error) {
	// check permissions if authz is enabled
	// if auth is disabled, policy decisions will be skipped
	if r.authDisabled {
		ctx = privacy.DecisionContext(ctx, privacy.Allow)
	} else {
		// setup view context
		v := viewer.UserViewer{
			// TODO: this isn't right, fix it
			OrgID: id,
		}

		ctx = viewer.NewContext(ctx, v)
	}

	orgMember, err := withTransactionalMutation(ctx).OrgMembership.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			r.logger.Errorw("failed to get org member on update", "error", err)

			return nil, newPermissionDeniedError(ActionGet, "org member")
		}

		r.logger.Errorw("failed to get org member", "error", err)
		return nil, ErrInternalServerError
	}

	orgMember, err = orgMember.Update().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			r.logger.Errorw("failed to update org member", "error", err)

			return nil, newPermissionDeniedError(ActionUpdate, "org member")
		}

		r.logger.Errorw("failed to update org member role", "error", err)
		return nil, ErrInternalServerError
	}

	return &OrgMembershipUpdatePayload{OrgMembership: orgMember}, nil
}

// DeleteOrgMembership is the resolver for the deleteOrgMembership field.
func (r *mutationResolver) DeleteOrgMembership(ctx context.Context, id string) (*OrgMembershipDeletePayload, error) {
	// check permissions if authz is enabled
	// if auth is disabled, policy decisions will be skipped
	if r.authDisabled {
		ctx = privacy.DecisionContext(ctx, privacy.Allow)
	} else {
		// setup view context
		// TODO: make this the org in question?
		v := viewer.UserViewer{}

		ctx = viewer.NewContext(ctx, v)
	}

	if err := withTransactionalMutation(ctx).OrgMembership.DeleteOneID(id).Exec(ctx); err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			return nil, newPermissionDeniedError(ActionDelete, "org member")
		}

		r.logger.Errorw("failed to delete org member", "error", err)
		return nil, err
	}

	if err := generated.OrgMembershipEdgeCleanup(ctx, id); err != nil {
		return nil, newCascadeDeleteError(err)
	}

	return &OrgMembershipDeletePayload{DeletedID: id}, nil
}

// OrgMembership is the resolver for the orgMembership field.
func (r *queryResolver) OrgMembership(ctx context.Context, id string) (*generated.OrgMembership, error) {
	// check permissions if authz is enabled
	// if auth is disabled, policy decisions will be skipped
	if r.authDisabled {
		ctx = privacy.DecisionContext(ctx, privacy.Allow)
	} else {
		// setup view context
		// TODO: fix this
		v := viewer.UserViewer{}

		ctx = viewer.NewContext(ctx, v)
	}

	org, err := withTransactionalMutation(ctx).OrgMembership.Get(ctx, id)
	if err != nil {
		r.logger.Errorw("failed to get members of organization", "error", err)

		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			return nil, newPermissionDeniedError(ActionGet, "org members")
		}

		return nil, ErrInternalServerError
	}

	return org, nil
}