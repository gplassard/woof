package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteTeamCmd = &cobra.Command{
	Use:   "deleteteam",
	Short: "Remove a team",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/team/{team_id}")
		fmt.Println("OperationID: DeleteTeam")
	},
}

func init() {
	Cmd.AddCommand(DeleteTeamCmd)
}
