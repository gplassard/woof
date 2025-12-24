package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListMemberTeamsCmd = &cobra.Command{
	Use:   "listmemberteams",
	Short: "Get all member teams",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/team/{super_team_id}/member_teams")
		fmt.Println("OperationID: ListMemberTeams")
	},
}

func init() {
	Cmd.AddCommand(ListMemberTeamsCmd)
}
