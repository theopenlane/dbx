package dbx

import (
	"context"
	"encoding/json"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	dbx "github.com/theopenlane/dbx/cmd/cli/cmd"
)

var databaseGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get an existing new dbx database",
	RunE: func(cmd *cobra.Command, args []string) error {
		return getDatabase(cmd.Context())
	},
}

func init() {
	databaseCmd.AddCommand(databaseGetCmd)

	databaseGetCmd.Flags().StringP("name", "n", "", "database name to query")
	dbx.ViperBindFlag("database.get.name", databaseGetCmd.Flags().Lookup("name"))
}

func getDatabase(ctx context.Context) error {
	// setup dbx http client
	cli, err := dbx.GetGraphClient()
	if err != nil {
		return err
	}

	if cli.Client == nil {
		log.Fatal("client is nil")
	}

	// filter options
	dName := viper.GetString("database.get.name")

	// if an db name is provided, filter on that db, otherwise get all
	if dName != "" {
		db, err := cli.Client.GetDatabase(ctx, dName, cli.Interceptor)
		if err != nil {
			return err
		}

		if viper.GetString("output.format") == "json" {
			s, err := json.Marshal(db.Database)
			if err != nil {
				return err
			}

			return dbx.JSONPrint(s)
		}

		return dbx.SingleRowTablePrint(db.Database)
	}

	dbs, err := cli.Client.GetAllDatabases(ctx, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err := json.Marshal(dbs.Databases)
	if err != nil {
		return err
	}

	// print json output
	if viper.GetString("output.format") == "json" {
		return dbx.JSONPrint(s)
	}

	// print table output
	var resp dbx.GraphResponse

	err = json.Unmarshal(s, &resp)
	if err != nil {
		return err
	}

	return dbx.RowsTablePrint(resp)
}
