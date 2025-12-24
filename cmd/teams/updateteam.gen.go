package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateTeamCmd = &cobra.Command{
	Use:   "updateteam",
	Short: "Update a team",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/team/{team_id}")
		fmt.Println("OperationID: UpdateTeam")
	},
}

func init() {
	Cmd.AddCommand(UpdateTeamCmd)
}
