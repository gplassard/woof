package incident_teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateIncidentTeamCmd = &cobra.Command{
	Use:   "updateincidentteam",
	Short: "Update an existing incident team",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/teams/{team_id}")
		fmt.Println("OperationID: UpdateIncidentTeam")
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentTeamCmd)
}
