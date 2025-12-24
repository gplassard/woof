package roles

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteRoleCmd = &cobra.Command{
	Use:   "deleterole",
	Short: "Delete role",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/roles/{role_id}")
		fmt.Println("OperationID: DeleteRole")
	},
}

func init() {
	Cmd.AddCommand(DeleteRoleCmd)
}
