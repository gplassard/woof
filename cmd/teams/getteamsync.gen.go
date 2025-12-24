package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetTeamSyncCmd = &cobra.Command{
	Use:   "getteamsync",
	Short: "Get team sync configurations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/team/sync")
		fmt.Println("OperationID: GetTeamSync")
	},
}

func init() {
	Cmd.AddCommand(GetTeamSyncCmd)
}
