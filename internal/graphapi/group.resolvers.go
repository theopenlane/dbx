package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/theopenlane/dbx/internal/ent/generated"
	"github.com/theopenlane/dbx/internal/ent/generated/group"
	"github.com/theopenlane/utils/rout"
)

// CreateGroup is the resolver for the createGroup field.
func (r *mutationResolver) CreateGroup(ctx context.Context, input generated.CreateGroupInput) (*GroupCreatePayload, error) {
	group, err := withTransactionalMutation(ctx).Group.Create().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsConstraintError(err) {
			constraintError := err.(*generated.ConstraintError)

			log.Debug().Err(constraintError).Msg("constraint error")

			return nil, constraintError
		}

		if generated.IsValidationError(err) {
			ve := err.(*generated.ValidationError)

			return nil, rout.InvalidField(ve.Name)
		}

		log.Error().Err(err).Msg("failed to create group")

		return nil, err
	}

	return &GroupCreatePayload{Group: group}, err
}

// UpdateGroup is the resolver for the updateGroup field.
func (r *mutationResolver) UpdateGroup(ctx context.Context, name string, input generated.UpdateGroupInput) (*GroupUpdatePayload, error) {
	group, err := withTransactionalMutation(ctx).Group.
		Query().
		Where(group.NameEQ(name)).
		Only(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to get group")

		return nil, err
	}

	g, err := group.Update().
		SetInput(input).
		Save(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to update group")

		return nil, err
	}

	return &GroupUpdatePayload{Group: g}, nil
}

// DeleteGroup is the resolver for the deleteGroup field.
func (r *mutationResolver) DeleteGroup(ctx context.Context, name string) (*GroupDeletePayload, error) {
	group, err := withTransactionalMutation(ctx).Group.Query().Where(group.NameEQ(name)).Only(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to get group")

		return nil, err
	}

	if err := generated.GroupEdgeCleanup(ctx, group.ID); err != nil {
		return nil, newCascadeDeleteError(err)
	}

	if err := withTransactionalMutation(ctx).Group.DeleteOneID(group.ID).Exec(ctx); err != nil {
		log.Error().Err(err).Msg("failed to delete group")

		return nil, err
	}

	return &GroupDeletePayload{DeletedID: group.ID}, nil
}

// Group is the resolver for the group field.
func (r *queryResolver) Group(ctx context.Context, name string) (*generated.Group, error) {
	group, err := withTransactionalMutation(ctx).Group.Query().Where(group.NameEQ(name)).Only(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to get group")

		return nil, err
	}

	return group, err
}
