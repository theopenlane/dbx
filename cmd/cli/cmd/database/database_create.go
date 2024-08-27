package dbx

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	dbx "github.com/theopenlane/dbx/cmd/cli/cmd"
	"github.com/theopenlane/dbx/pkg/dbxclient"
	"github.com/theopenlane/dbx/pkg/enums"
)

var databaseCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new dbx database",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createDatabase(cmd.Context())
	},
}

func init() {
	databaseCmd.AddCommand(databaseCreateCmd)

	databaseCreateCmd.Flags().StringP("org-id", "o", "", "owning organization id of the database")
	dbx.ViperBindFlag("database.create.orgid", databaseCreateCmd.Flags().Lookup("org-id"))

	databaseCreateCmd.Flags().StringP("provider", "p", "turso", "provider of the database (local, turso)")
	dbx.ViperBindFlag("database.create.provider", databaseCreateCmd.Flags().Lookup("provider"))

	databaseCreateCmd.Flags().StringP("group-id", "g", "", "group name to assign to the database")
	dbx.ViperBindFlag("database.create.groupid", databaseCreateCmd.Flags().Lookup("group-id"))
}

func createDatabase(ctx context.Context) error {
	cli, err := dbx.GetGraphClient()
	if err != nil {
		return err
	}

	orgID := viper.GetString("database.create.orgid")
	if orgID == "" {
		return dbx.NewRequiredFieldMissingError("organization_id")
	}

	provider := viper.GetString("database.create.provider")
	if provider == "" {
		return dbx.NewRequiredFieldMissingError("provider")
	}

	groupID := viper.GetString("database.create.groupid")
	if groupID == "" {
		return dbx.NewRequiredFieldMissingError("group_id")
	}

	input := dbxclient.CreateDatabaseInput{
		OrganizationID: orgID,
		Provider:       enums.ToDatabaseProvider(provider),
		GroupID:        groupID,
	}

	d, err := cli.Client.CreateDatabase(ctx, input, cli.Interceptor)
	if err != nil {
		return err
	}

	if viper.GetString("output.format") == "json" {
		s, err := json.Marshal(d.CreateDatabase.Database)
		if err != nil {
			return err
		}

		return dbx.JSONPrint(s)
	}

	return dbx.SingleRowTablePrint(d.CreateDatabase.Database)
}
