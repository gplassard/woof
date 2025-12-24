package incident_teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListIncidentTeamsCmd = &cobra.Command{
	Use:   "listincidentteams",
	Short: "Get a list of all incident teams",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/teams")
		fmt.Println("OperationID: ListIncidentTeams")
	},
}

func init() {
	Cmd.AddCommand(ListIncidentTeamsCmd)
}
