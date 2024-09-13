package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/theopenlane/dbx/internal/ent/generated"
	"github.com/theopenlane/dbx/internal/ent/generated/database"
	"github.com/theopenlane/utils/rout"
)

// CreateDatabase is the resolver for the createDatabase field.
func (r *mutationResolver) CreateDatabase(ctx context.Context, input generated.CreateDatabaseInput) (*DatabaseCreatePayload, error) {
	db, err := withTransactionalMutation(ctx).Database.Create().SetInput(input).Save(ctx)
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

		log.Error().Err(err).Msg("failed to create database")

		return nil, err
	}

	return &DatabaseCreatePayload{Database: db}, err
}

// UpdateDatabase is the resolver for the updateDatabase field.
func (r *mutationResolver) UpdateDatabase(ctx context.Context, name string, input generated.UpdateDatabaseInput) (*DatabaseUpdatePayload, error) {
	panic(fmt.Errorf("not implemented: UpdateDatabase - updateDatabase"))
}

// DeleteDatabase is the resolver for the deleteDatabase field.
func (r *mutationResolver) DeleteDatabase(ctx context.Context, name string) (*DatabaseDeletePayload, error) {
	db, err := withTransactionalMutation(ctx).Database.Query().Where(database.NameEQ(name)).Only(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to get database")

		return nil, err
	}

	if err := withTransactionalMutation(ctx).Database.DeleteOneID(db.ID).Exec(ctx); err != nil {
		log.Error().Err(err).Msg("failed to delete database")

		return nil, err
	}

	return &DatabaseDeletePayload{DeletedID: db.ID}, nil
}

// Database is the resolver for the database field.
func (r *queryResolver) Database(ctx context.Context, name string) (*generated.Database, error) {
	db, err := withTransactionalMutation(ctx).Database.Query().Where(database.NameEQ(name)).Only(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to get database")

		return nil, err
	}

	return db, err
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
