package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetTeamCmd = &cobra.Command{
	Use:   "getteam",
	Short: "Get a team",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/team/{team_id}")
		fmt.Println("OperationID: GetTeam")
	},
}

func init() {
	Cmd.AddCommand(GetTeamCmd)
}
