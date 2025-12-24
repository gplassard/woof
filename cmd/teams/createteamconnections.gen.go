package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateTeamConnectionsCmd = &cobra.Command{
	Use:   "createteamconnections",
	Short: "Create team connections",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/team/connections")
		fmt.Println("OperationID: CreateTeamConnections")
	},
}

func init() {
	Cmd.AddCommand(CreateTeamConnectionsCmd)
}
