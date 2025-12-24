package action_connection

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateActionConnectionCmd = &cobra.Command{
	Use:   "createactionconnection",
	Short: "Create a new Action Connection",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/actions/connections")
		fmt.Println("OperationID: CreateActionConnection")
	},
}

func init() {
	Cmd.AddCommand(CreateActionConnectionCmd)
}
