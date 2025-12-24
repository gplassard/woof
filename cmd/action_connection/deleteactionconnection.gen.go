package action_connection

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteActionConnectionCmd = &cobra.Command{
	Use:   "deleteactionconnection",
	Short: "Delete an existing Action Connection",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/actions/connections/{connection_id}")
		fmt.Println("OperationID: DeleteActionConnection")
	},
}

func init() {
	Cmd.AddCommand(DeleteActionConnectionCmd)
}
