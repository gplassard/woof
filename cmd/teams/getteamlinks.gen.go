package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetTeamLinksCmd = &cobra.Command{
	Use:   "getteamlinks",
	Short: "Get links for a team",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/team/{team_id}/links")
		fmt.Println("OperationID: GetTeamLinks")
	},
}

func init() {
	Cmd.AddCommand(GetTeamLinksCmd)
}
