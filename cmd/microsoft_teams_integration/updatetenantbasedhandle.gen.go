package microsoft_teams_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateTenantBasedHandleCmd = &cobra.Command{
	Use:   "updatetenantbasedhandle",
	Short: "Update tenant-based handle",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/integration/ms-teams/configuration/tenant-based-handles/{handle_id}")
		fmt.Println("OperationID: UpdateTenantBasedHandle")
	},
}

func init() {
	Cmd.AddCommand(UpdateTenantBasedHandleCmd)
}
