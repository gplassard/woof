package incident_teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteIncidentTeamCmd = &cobra.Command{
	Use:   "deleteincidentteam",
	Short: "Delete an existing incident team",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/teams/{team_id}")
		fmt.Println("OperationID: DeleteIncidentTeam")
	},
}

func init() {
	Cmd.AddCommand(DeleteIncidentTeamCmd)
}
