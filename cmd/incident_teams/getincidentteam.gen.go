package incident_teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetIncidentTeamCmd = &cobra.Command{
	Use:   "getincidentteam",
	Short: "Get details of an incident team",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/teams/{team_id}")
		fmt.Println("OperationID: GetIncidentTeam")
	},
}

func init() {
	Cmd.AddCommand(GetIncidentTeamCmd)
}
