package actions_datastores

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "actions-datastores",
	Short: "actions_datastores endpoints",
	Aliases: []string{
		"datastores",
	},
}
