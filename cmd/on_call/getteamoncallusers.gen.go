package on_call

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetTeamOnCallUsersCmd = &cobra.Command{
	Use:   "getteamoncallusers",
	Short: "Get team on-call users",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/on-call/teams/{team_id}/on-call")
		fmt.Println("OperationID: GetTeamOnCallUsers")
	},
}

func init() {
	Cmd.AddCommand(GetTeamOnCallUsersCmd)
}
