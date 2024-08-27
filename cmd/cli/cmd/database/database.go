package dbx

import (
	"github.com/spf13/cobra"

	dbx "github.com/theopenlane/dbx/cmd/cli/cmd"
)

// databaseCmd represents the base database command when called without any subcommands
var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "The subcommands for working with dbx databases",
}

func init() {
	dbx.RootCmd.AddCommand(databaseCmd)
}
