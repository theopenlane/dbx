package dbxgroup

import (
	"github.com/spf13/cobra"

	dbx "github.com/theopenlane/dbx/cmd/cli/cmd"
)

// groupCmd represents the base group command when called without any subcommands
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "The subcommands for working with dbx groups",
}

func init() {
	dbx.RootCmd.AddCommand(groupCmd)
}
