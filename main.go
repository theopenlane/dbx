// package main is the entry point
package main

import (
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"

	"github.com/theopenlane/dbx/cmd"
	_ "github.com/theopenlane/dbx/internal/ent/generated/runtime"
)

func main() {
	cmd.Execute()
}
