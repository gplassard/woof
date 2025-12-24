package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RemoveMemberTeamCmd = &cobra.Command{
	Use:   "removememberteam",
	Short: "Remove a member team",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/team/{super_team_id}/member_teams/{member_team_id}")
		fmt.Println("OperationID: RemoveMemberTeam")
	},
}

func init() {
	Cmd.AddCommand(RemoveMemberTeamCmd)
}
