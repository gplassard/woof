package users

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListUserPermissionsCmd = &cobra.Command{
	Use:   "listuserpermissions",
	Short: "Get a user permissions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/users/{user_id}/permissions")
		fmt.Println("OperationID: ListUserPermissions")
	},
}

func init() {
	Cmd.AddCommand(ListUserPermissionsCmd)
}
