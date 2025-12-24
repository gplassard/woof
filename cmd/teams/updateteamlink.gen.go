package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateTeamLinkCmd = &cobra.Command{
	Use:   "updateteamlink",
	Short: "Update a team link",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/team/{team_id}/links/{link_id}")
		fmt.Println("OperationID: UpdateTeamLink")
	},
}

func init() {
	Cmd.AddCommand(UpdateTeamLinkCmd)
}
