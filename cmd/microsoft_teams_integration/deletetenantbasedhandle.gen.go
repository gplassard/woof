package microsoft_teams_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteTenantBasedHandleCmd = &cobra.Command{
	Use:   "deletetenantbasedhandle",
	Short: "Delete tenant-based handle",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/integration/ms-teams/configuration/tenant-based-handles/{handle_id}")
		fmt.Println("OperationID: DeleteTenantBasedHandle")
	},
}

func init() {
	Cmd.AddCommand(DeleteTenantBasedHandleCmd)
}
