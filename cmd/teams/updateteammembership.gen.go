package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateTeamMembershipCmd = &cobra.Command{
	Use:   "updateteammembership",
	Short: "Update a user's membership attributes on a team",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/team/{team_id}/memberships/{user_id}")
		fmt.Println("OperationID: UpdateTeamMembership")
	},
}

func init() {
	Cmd.AddCommand(UpdateTeamMembershipCmd)
}
