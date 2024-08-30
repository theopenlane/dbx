package hooks

import (
	"context"
	"fmt"
	"strings"

	"entgo.io/ent"

	"github.com/99designs/gqlgen/graphql"
	"github.com/theopenlane/go-turso"
	"github.com/theopenlane/utils/rout"

	"github.com/theopenlane/dbx/internal/ent/generated"
	"github.com/theopenlane/dbx/internal/ent/generated/group"
	"github.com/theopenlane/dbx/internal/ent/generated/hook"
	"github.com/theopenlane/dbx/pkg/enums"
)

// HookCreateDatabase creates sets the name of the database and creates the database in turso, if the provider is turso
func HookCreateDatabase() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.DatabaseFunc(func(ctx context.Context, mutation *generated.DatabaseMutation) (generated.Value, error) {
			// get organization and provider from the request
			orgID, _ := mutation.OrganizationID()
			provider, _ := mutation.Provider()

			// create a name for the database
			name := strings.ToLower(fmt.Sprintf("org-%s", orgID))
			mutation.SetName(name)

			// if the provider is turso and turso is enabled, create a database
			if provider == enums.Turso && mutation.Turso != nil {
				// get the group to assign the database to
				groupName, err := getGroupName(ctx, mutation)
				if err != nil {
					return nil, err
				}

				// create a turso db
				body := turso.CreateDatabaseRequest{
					Group: groupName,
					Name:  name,
				}

				// create the database in turso
				db, err := mutation.Turso.Database.CreateDatabase(ctx, body)
				if err != nil {
					return nil, err
				}

				mutation.Logger.Infow("created turso db", "db", db.Database.DatabaseID, "hostname", db.Database.Hostname)

				mutation.SetDsn(db.Database.Hostname)
			} else {
				// set the dsn to the name
				mutation.SetDsn(fmt.Sprintf("file:%s.db", name))
			}

			// set the status of the database to active
			mutation.SetStatus(enums.Active)

			// write things that we need to the database
			return next.Mutate(ctx, mutation)
		})
	}, ent.OpCreate)
}

// HookDatabaseDelete deletes the database in turso
func HookDatabaseDelete() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.DatabaseFunc(func(ctx context.Context, mutation *generated.DatabaseMutation) (generated.Value, error) {
			if ok := graphql.HasOperationContext(ctx); ok {
				// TODO: this only works for a delete database and not on a cascade delete
				gtx := graphql.GetOperationContext(ctx)
				name := gtx.Variables["name"].(string)

				if name == "" {
					mutation.Logger.Errorw("unable to delete database, no name provided")

					return nil, rout.InvalidField("name")
				}

				db, err := mutation.Turso.Database.DeleteDatabase(ctx, name)
				if err != nil {
					return nil, err
				}

				mutation.Logger.Infow("deleted turso database", "database", db.Database)
			}

			// write things that we need to the database
			return next.Mutate(ctx, mutation)
		})
	}, ent.OpDelete|ent.OpDeleteOne)
}

// getGroupName gets the group name associated with the geo or group id
func getGroupName(ctx context.Context, mutation *generated.DatabaseMutation) (string, error) {
	groupID, ok := mutation.GroupID()

	// if the group id is set, get the group by the group id
	if ok && groupID != "" {
		g, err := mutation.Client().Group.Get(ctx, groupID)
		if err != nil {
			mutation.Logger.Errorw("unable to get group, invalid group ID", "error", err)

			return "", err
		}

		return g.Name, nil
	}

	// else get the group by the geo
	geo, ok := mutation.Geo()

	if !ok || geo == "" {
		mutation.Logger.Errorw("unable to get geo or group id, cannot create database")

		return "", rout.InvalidField("geo")
	}

	g, err := mutation.Client().Group.Query().Where(group.RegionEQ(enums.Region(geo))).Only(ctx)
	if err != nil {
		mutation.Logger.Errorw("unable to get associated group", "error", err)

		return "", err
	}

	if g == nil {
		mutation.Logger.Errorw("unable to get associated group", "geo", geo)

		return "", rout.InvalidField("geo")
	}

	// set the group id on the mutation
	mutation.SetGroupID(g.ID)

	return g.Name, nil
}
