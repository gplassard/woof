package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetTeamMembershipsCmd = &cobra.Command{
	Use:   "getteammemberships",
	Short: "Get team memberships",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/team/{team_id}/memberships")
		fmt.Println("OperationID: GetTeamMemberships")
	},
}

func init() {
	Cmd.AddCommand(GetTeamMembershipsCmd)
}
