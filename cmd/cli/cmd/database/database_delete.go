package dbx

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	dbx "github.com/theopenlane/dbx/cmd/cli/cmd"
)

var databaseDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing dbx database",
	RunE: func(cmd *cobra.Command, _ []string) error {
		return deleteDatabase(cmd.Context())
	},
}

func init() {
	databaseCmd.AddCommand(databaseDeleteCmd)

	databaseDeleteCmd.Flags().StringP("name", "n", "", "database name to delete")
	dbx.ViperBindFlag("database.delete.name", databaseDeleteCmd.Flags().Lookup("name"))
}

func deleteDatabase(ctx context.Context) error {
	// setup dbx http client
	cli, err := dbx.GetGraphClient()
	if err != nil {
		return err
	}

	dName := viper.GetString("database.delete.name")
	if dName == "" {
		return dbx.NewRequiredFieldMissingError("name")
	}

	d, err := cli.Client.DeleteDatabase(ctx, dName, cli.Interceptor)
	if err != nil {
		return err
	}

	if viper.GetString("output.format") == "json" {
		s, err := json.Marshal(d.DeleteDatabase)
		if err != nil {
			return err
		}

		return dbx.JSONPrint(s)
	}

	return dbx.SingleRowTablePrint(d.DeleteDatabase)
}
