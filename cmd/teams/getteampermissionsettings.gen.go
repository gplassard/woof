package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetTeamPermissionSettingsCmd = &cobra.Command{
	Use:   "getteampermissionsettings",
	Short: "Get permission settings for a team",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/team/{team_id}/permission-settings")
		fmt.Println("OperationID: GetTeamPermissionSettings")
	},
}

func init() {
	Cmd.AddCommand(GetTeamPermissionSettingsCmd)
}
