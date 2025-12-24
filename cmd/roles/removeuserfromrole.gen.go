package roles

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RemoveUserFromRoleCmd = &cobra.Command{
	Use:   "removeuserfromrole",
	Short: "Remove a user from a role",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/roles/{role_id}/users")
		fmt.Println("OperationID: RemoveUserFromRole")
	},
}

func init() {
	Cmd.AddCommand(RemoveUserFromRoleCmd)
}
