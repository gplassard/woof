package roles

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListPermissionsCmd = &cobra.Command{
	Use:   "listpermissions",
	Short: "List permissions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/permissions")
		fmt.Println("OperationID: ListPermissions")
	},
}

func init() {
	Cmd.AddCommand(ListPermissionsCmd)
}
