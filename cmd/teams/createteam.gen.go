package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateTeamCmd = &cobra.Command{
	Use:   "createteam",
	Short: "Create a team",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/team")
		fmt.Println("OperationID: CreateTeam")
	},
}

func init() {
	Cmd.AddCommand(CreateTeamCmd)
}
