package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListTeamConnectionsCmd = &cobra.Command{
	Use:   "listteamconnections",
	Short: "List team connections",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/team/connections")
		fmt.Println("OperationID: ListTeamConnections")
	},
}

func init() {
	Cmd.AddCommand(ListTeamConnectionsCmd)
}
