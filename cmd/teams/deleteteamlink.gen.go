package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteTeamLinkCmd = &cobra.Command{
	Use:   "deleteteamlink",
	Short: "Remove a team link",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/team/{team_id}/links/{link_id}")
		fmt.Println("OperationID: DeleteTeamLink")
	},
}

func init() {
	Cmd.AddCommand(DeleteTeamLinkCmd)
}
