package roles

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListRoleUsersCmd = &cobra.Command{
	Use:   "listroleusers",
	Short: "Get all users of a role",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/roles/{role_id}/users")
		fmt.Println("OperationID: ListRoleUsers")
	},
}

func init() {
	Cmd.AddCommand(ListRoleUsersCmd)
}
