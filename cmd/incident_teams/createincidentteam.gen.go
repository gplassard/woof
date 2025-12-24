package incident_teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateIncidentTeamCmd = &cobra.Command{
	Use:   "createincidentteam",
	Short: "Create a new incident team",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/teams")
		fmt.Println("OperationID: CreateIncidentTeam")
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentTeamCmd)
}
