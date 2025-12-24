package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteTeamMembershipCmd = &cobra.Command{
	Use:   "deleteteammembership",
	Short: "Remove a user from a team",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/team/{team_id}/memberships/{user_id}")
		fmt.Println("OperationID: DeleteTeamMembership")
	},
}

func init() {
	Cmd.AddCommand(DeleteTeamMembershipCmd)
}
