package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateTeamPermissionSettingCmd = &cobra.Command{
	Use:   "updateteampermissionsetting",
	Short: "Update permission setting for team",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/team/{team_id}/permission-settings/{action}")
		fmt.Println("OperationID: UpdateTeamPermissionSetting")
	},
}

func init() {
	Cmd.AddCommand(UpdateTeamPermissionSettingCmd)
}
