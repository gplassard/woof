package microsoft_teams_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetTenantBasedHandleCmd = &cobra.Command{
	Use:   "gettenantbasedhandle",
	Short: "Get tenant-based handle information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integration/ms-teams/configuration/tenant-based-handles/{handle_id}")
		fmt.Println("OperationID: GetTenantBasedHandle")
	},
}

func init() {
	Cmd.AddCommand(GetTenantBasedHandleCmd)
}
