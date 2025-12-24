package roles

import (
	"fmt"
	"github.com/spf13/cobra"
)

var AddPermissionToRoleCmd = &cobra.Command{
	Use:   "addpermissiontorole",
	Short: "Grant permission to a role",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/roles/{role_id}/permissions")
		fmt.Println("OperationID: AddPermissionToRole")
	},
}

func init() {
	Cmd.AddCommand(AddPermissionToRoleCmd)
}
