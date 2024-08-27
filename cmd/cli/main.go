package main

import (
	dbx "github.com/theopenlane/dbx/cmd/cli/cmd"

	// since the cmds are no longer part of the same package
	// they must all be imported in main
	_ "github.com/theopenlane/dbx/cmd/cli/cmd/database"
	_ "github.com/theopenlane/dbx/cmd/cli/cmd/group"
	_ "github.com/theopenlane/dbx/cmd/cli/cmd/version"
)

func main() {
	dbx.Execute()
}
