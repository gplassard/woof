package microsoft_teams_integration

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "microsoft-teams-integration",
	Short: "microsoft_teams_integration endpoints",
	Aliases: []string{
		"msteams-integration",
	},
}
