package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateTeamMembershipCmd = &cobra.Command{
	Use:   "createteammembership",
	Short: "Add a user to a team",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/team/{team_id}/memberships")
		fmt.Println("OperationID: CreateTeamMembership")
	},
}

func init() {
	Cmd.AddCommand(CreateTeamMembershipCmd)
}
