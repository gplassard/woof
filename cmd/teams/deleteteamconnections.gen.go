package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteTeamConnectionsCmd = &cobra.Command{
	Use:   "deleteteamconnections",
	Short: "Delete team connections",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/team/connections")
		fmt.Println("OperationID: DeleteTeamConnections")
	},
}

func init() {
	Cmd.AddCommand(DeleteTeamConnectionsCmd)
}
