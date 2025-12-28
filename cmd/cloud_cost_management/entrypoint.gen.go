package cloud_cost_management

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "cloud_cost_management",
	Short: "cloud_cost_management endpoints",
	Aliases: []string{
		"ccm",
	},
}
