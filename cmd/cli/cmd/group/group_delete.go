package dbxgroup

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	dbx "github.com/theopenlane/dbx/cmd/cli/cmd"
)

var groupDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing dbx group",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteGroup(cmd.Context())
	},
}

func init() {
	groupCmd.AddCommand(groupDeleteCmd)

	groupDeleteCmd.Flags().StringP("name", "n", "", "group name to delete")
	dbx.ViperBindFlag("group.delete.name", groupDeleteCmd.Flags().Lookup("name"))
}

func deleteGroup(ctx context.Context) error {
	// setup dbx http client
	cli, err := dbx.GetGraphClient()
	if err != nil {
		return err
	}

	gName := viper.GetString("group.delete.name")
	if gName == "" {
		return dbx.NewRequiredFieldMissingError("name")
	}

	g, err := cli.Client.DeleteGroup(ctx, gName, cli.Interceptor)
	if err != nil {
		return err
	}

	if viper.GetString("output.format") == "json" {
		s, err := json.Marshal(g.DeleteGroup)
		if err != nil {
			return err
		}

		return dbx.JSONPrint(s)
	}

	return dbx.SingleRowTablePrint(g.DeleteGroup)
}
