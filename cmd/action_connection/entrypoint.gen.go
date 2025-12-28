package action_connection

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "action-connection",
	Short: "action_connection endpoints",
	Aliases: []string{
		"connection",
	},
}
