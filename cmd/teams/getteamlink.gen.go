package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetTeamLinkCmd = &cobra.Command{
	Use:   "getteamlink",
	Short: "Get a team link",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/team/{team_id}/links/{link_id}")
		fmt.Println("OperationID: GetTeamLink")
	},
}

func init() {
	Cmd.AddCommand(GetTeamLinkCmd)
}
