package microsoft_teams_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateTenantBasedHandleCmd = &cobra.Command{
	Use:   "createtenantbasedhandle",
	Short: "Create tenant-based handle",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/integration/ms-teams/configuration/tenant-based-handles")
		fmt.Println("OperationID: CreateTenantBasedHandle")
	},
}

func init() {
	Cmd.AddCommand(CreateTenantBasedHandleCmd)
}
