package action_connection

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateActionConnectionCmd = &cobra.Command{
	Use:   "updateactionconnection",
	Short: "Update an existing Action Connection",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/actions/connections/{connection_id}")
		fmt.Println("OperationID: UpdateActionConnection")
	},
}

func init() {
	Cmd.AddCommand(UpdateActionConnectionCmd)
}
