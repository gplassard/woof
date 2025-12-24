package roles

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListRoleTemplatesCmd = &cobra.Command{
	Use:   "listroletemplates",
	Short: "List role templates",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/roles/templates")
		fmt.Println("OperationID: ListRoleTemplates")
	},
}

func init() {
	Cmd.AddCommand(ListRoleTemplatesCmd)
}
