package users

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListUsersCmd = &cobra.Command{
	Use:   "listusers",
	Short: "List all users",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/users")
		fmt.Println("OperationID: ListUsers")
	},
}

func init() {
	Cmd.AddCommand(ListUsersCmd)
}
