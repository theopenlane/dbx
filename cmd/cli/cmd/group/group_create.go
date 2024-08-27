package dbxgroup

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	dbx "github.com/theopenlane/dbx/cmd/cli/cmd"
	"github.com/theopenlane/dbx/pkg/dbxclient"
	"github.com/theopenlane/dbx/pkg/enums"
)

var groupCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new dbx group",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createGroup(cmd.Context())
	},
}

func init() {
	groupCmd.AddCommand(groupCreateCmd)

	groupCreateCmd.Flags().StringP("name", "n", "", "name of the group")
	dbx.ViperBindFlag("group.create.name", groupCreateCmd.Flags().Lookup("name"))

	groupCreateCmd.Flags().StringP("description", "d", "", "description of the group")
	dbx.ViperBindFlag("group.create.description", groupCreateCmd.Flags().Lookup("description"))

	groupCreateCmd.Flags().StringP("region", "r", "", "region of the group (AMER, EMEA, APAC)")
	dbx.ViperBindFlag("group.create.region", groupCreateCmd.Flags().Lookup("region"))

	groupCreateCmd.Flags().StringP("primary-location", "l", "", "primary location of the group")
	dbx.ViperBindFlag("group.create.location", groupCreateCmd.Flags().Lookup("primary-location"))
}

func createGroup(ctx context.Context) error {
	cli, err := dbx.GetGraphClient()
	if err != nil {
		return err
	}

	name := viper.GetString("group.create.name")
	if name == "" {
		return dbx.NewRequiredFieldMissingError("name")
	}

	description := viper.GetString("group.create.description")
	location := viper.GetString("group.create.location")
	region := viper.GetString("group.create.region")

	input := dbxclient.CreateGroupInput{
		Name:            name,
		Description:     &description,
		PrimaryLocation: location,
		Region:          enums.ToRegion(region),
	}

	g, err := cli.Client.CreateGroup(ctx, input, cli.Interceptor)
	if err != nil {
		return err
	}

	if viper.GetString("output.format") == "json" {
		s, err := json.Marshal(g.CreateGroup.Group)
		if err != nil {
			return err
		}

		return dbx.JSONPrint(s)
	}

	return dbx.SingleRowTablePrint(g.CreateGroup.Group)
}
