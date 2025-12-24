package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var AddMemberTeamCmd = &cobra.Command{
	Use:   "addmemberteam",
	Short: "Add a member team",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/team/{super_team_id}/member_teams")
		fmt.Println("OperationID: AddMemberTeam")
	},
}

func init() {
	Cmd.AddCommand(AddMemberTeamCmd)
}
