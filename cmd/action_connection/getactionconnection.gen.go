package action_connection

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetActionConnectionCmd = &cobra.Command{
	Use:   "getactionconnection",
	Short: "Get an existing Action Connection",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/actions/connections/{connection_id}")
		fmt.Println("OperationID: GetActionConnection")
	},
}

func init() {
	Cmd.AddCommand(GetActionConnectionCmd)
}
