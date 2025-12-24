package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListTeamsCmd = &cobra.Command{
	Use:   "listteams",
	Short: "Get all teams",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/team")
		fmt.Println("OperationID: ListTeams")
	},
}

func init() {
	Cmd.AddCommand(ListTeamsCmd)
}
