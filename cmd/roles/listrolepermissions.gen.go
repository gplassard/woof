package roles

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListRolePermissionsCmd = &cobra.Command{
	Use:   "listrolepermissions",
	Short: "List permissions for a role",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/roles/{role_id}/permissions")
		fmt.Println("OperationID: ListRolePermissions")
	},
}

func init() {
	Cmd.AddCommand(ListRolePermissionsCmd)
}
