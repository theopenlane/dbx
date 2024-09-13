package hooks

import (
	"context"

	"entgo.io/ent"

	"github.com/99designs/gqlgen/graphql"
	"github.com/rs/zerolog/log"
	"github.com/theopenlane/go-turso"
	"github.com/theopenlane/utils/rout"

	"github.com/theopenlane/dbx/internal/ent/generated"
	"github.com/theopenlane/dbx/internal/ent/generated/hook"
)

func HookGroupCreate() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.GroupFunc(func(ctx context.Context, mutation *generated.GroupMutation) (generated.Value, error) {
			name, _ := mutation.Name()
			loc, _ := mutation.PrimaryLocation()

			// create a turso group
			body := turso.CreateGroupRequest{
				Name:     name,
				Location: loc,
			}

			group, err := mutation.Turso.Group.CreateGroup(ctx, body)
			if err != nil {
				return nil, err
			}

			log.Info().Str("group", group.Group.Name).Strs("locations", group.Group.Locations).Msg("created turso group")

			// write things that we need to the database
			return next.Mutate(ctx, mutation)
		})
	}, ent.OpCreate)
}

func HookGroupUpdate() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.GroupFunc(func(ctx context.Context, mutation *generated.GroupMutation) (generated.Value, error) {
			name, _ := mutation.Name()
			locs, _ := mutation.Locations()

			// first get the group from Turso
			group, err := mutation.Turso.Group.GetGroup(ctx, name)
			if err != nil {
				return nil, err
			}

			// Add locations to the group that don't exist
			for _, loc := range locs {
				if !exists(loc, group.Group.Locations) {
					// add location to the group
					req := turso.GroupLocationRequest{
						GroupName: name,
						Location:  loc,
					}

					if _, err := mutation.Turso.Group.AddLocation(ctx, req); err != nil {
						log.Error().Str("group", name).Str("location", loc).Err(err).Msg("failed to add location to group")

						return nil, err
					}
				}
			}

			// write things that we need to the database
			return next.Mutate(ctx, mutation)
		})
	}, ent.OpUpdate|ent.OpUpdateOne)
}

func HookGroupDelete() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.GroupFunc(func(ctx context.Context, mutation *generated.GroupMutation) (generated.Value, error) {
			if ok := graphql.HasOperationContext(ctx); ok {
				gtx := graphql.GetOperationContext(ctx)
				name := gtx.Variables["name"].(string)

				if name == "" {
					log.Error().Msg("unable to delete group, no name provided")

					return nil, rout.InvalidField("name")
				}

				group, err := mutation.Turso.Group.DeleteGroup(ctx, name)
				if err != nil {
					return nil, err
				}

				log.Info().Interface("group", group.Group).Msg("deleted turso group")
			}

			// write things that we need to the database
			return next.Mutate(ctx, mutation)
		})
	}, ent.OpDelete|ent.OpDeleteOne)
}

// exists checks if a location exists in a list of locations
func exists(loc string, locs []string) bool {
	for _, l := range locs {
		if l == loc {
			return true
		}
	}

	return false
}
