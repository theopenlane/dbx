// package main is the entry point
package main

import (
	"github.com/theopenlane/dbx/cmd"
	_ "github.com/theopenlane/dbx/internal/ent/generated/runtime"
)

func main() {
	cmd.Execute()
}
