package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateTeamLinkCmd = &cobra.Command{
	Use:   "createteamlink",
	Short: "Create a team link",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/team/{team_id}/links")
		fmt.Println("OperationID: CreateTeamLink")
	},
}

func init() {
	Cmd.AddCommand(CreateTeamLinkCmd)
}
