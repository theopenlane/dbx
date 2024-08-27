package version

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"

	dbx "github.com/theopenlane/dbx/cmd/cli/cmd"
	"github.com/theopenlane/dbx/internal/constants"
)

// VersionCmd is the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print dbx CLI version",
	Long:  `The version command prints the version of the dbx CLI`,
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Println(constants.VerboseCLIVersion)
		cmd.Printf("User Agent: %s\n", GetUserAgent())
	},
}

func init() {
	dbx.RootCmd.AddCommand(versionCmd)
}

func GetUserAgent() string {
	product := "dbx-cli"
	productVersion := constants.CLIVersion

	userAgent := fmt.Sprintf("%s/%s (%s) %s (%s)",
		product, productVersion, runtime.GOOS, runtime.GOARCH, runtime.Version())

	return userAgent
}
