package on_call

import (
	"fmt"
	"github.com/spf13/cobra"
)

var SetOnCallTeamRoutingRulesCmd = &cobra.Command{
	Use:   "setoncallteamroutingrules",
	Short: "Set On-Call team routing rules",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/on-call/teams/{team_id}/routing-rules")
		fmt.Println("OperationID: SetOnCallTeamRoutingRules")
	},
}

func init() {
	Cmd.AddCommand(SetOnCallTeamRoutingRulesCmd)
}
