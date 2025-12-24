package roles

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "roles",
	Short: "Manage Datadog roles",
}

func init() {
	Cmd.AddCommand(listRoles)
}
