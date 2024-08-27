package dbxgroup

import (
	"context"
	"encoding/json"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	dbx "github.com/theopenlane/dbx/cmd/cli/cmd"
)

var groupGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get an existing dbx group",
	RunE: func(cmd *cobra.Command, args []string) error {
		return getGroup(cmd.Context())
	},
}

func init() {
	groupCmd.AddCommand(groupGetCmd)

	groupGetCmd.Flags().StringP("name", "n", "", "group name to query")
	dbx.ViperBindFlag("group.get.name", groupGetCmd.Flags().Lookup("name"))
}

func getGroup(ctx context.Context) error {
	// setup dbx http client
	cli, err := dbx.GetGraphClient()
	if err != nil {
		return err
	}

	if cli.Client == nil {
		log.Fatal("client is nil")
	}

	// filter options
	gName := viper.GetString("group.get.name")

	// if an group name is provided, filter on that group, otherwise get all
	if gName != "" {
		group, err := cli.Client.GetGroup(ctx, gName, cli.Interceptor)
		if err != nil {
			return err
		}

		if viper.GetString("output.format") == "json" {
			s, err := json.Marshal(group.Group)
			if err != nil {
				return err
			}

			return dbx.JSONPrint(s)
		}

		return dbx.SingleRowTablePrint(group.Group)
	}

	groups, err := cli.Client.GetAllGroups(ctx, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err := json.Marshal(groups.Groups)
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
