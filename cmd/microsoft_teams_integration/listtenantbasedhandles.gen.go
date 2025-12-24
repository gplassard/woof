package microsoft_teams_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListTenantBasedHandlesCmd = &cobra.Command{
	Use:   "listtenantbasedhandles",
	Short: "Get all tenant-based handles",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integration/ms-teams/configuration/tenant-based-handles")
		fmt.Println("OperationID: ListTenantBasedHandles")
	},
}

func init() {
	Cmd.AddCommand(ListTenantBasedHandlesCmd)
}
