package on_call

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetOnCallTeamRoutingRulesCmd = &cobra.Command{
	Use:   "getoncallteamroutingrules",
	Short: "Get On-Call team routing rules",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/on-call/teams/{team_id}/routing-rules")
		fmt.Println("OperationID: GetOnCallTeamRoutingRules")
	},
}

func init() {
	Cmd.AddCommand(GetOnCallTeamRoutingRulesCmd)
}
