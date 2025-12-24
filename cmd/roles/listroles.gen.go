package roles

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListRolesCmd = &cobra.Command{
	Use:   "listroles",
	Short: "List roles",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/roles")
		fmt.Println("OperationID: ListRoles")
	},
}

func init() {
	Cmd.AddCommand(ListRolesCmd)
}
