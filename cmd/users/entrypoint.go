package users

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "users",
	Short: "Manage Datadog users",
}

func init() {
	Cmd.AddCommand(
		listUsers,
		listUserOrganizations,
	)
}
