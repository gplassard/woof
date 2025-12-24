package roles

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RemovePermissionFromRoleCmd = &cobra.Command{
	Use:   "removepermissionfromrole",
	Short: "Revoke permission",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/roles/{role_id}/permissions")
		fmt.Println("OperationID: RemovePermissionFromRole")
	},
}

func init() {
	Cmd.AddCommand(RemovePermissionFromRoleCmd)
}
