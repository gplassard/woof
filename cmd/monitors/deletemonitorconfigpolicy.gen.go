package monitors

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteMonitorConfigPolicyCmd = &cobra.Command{
	Use:   "deletemonitorconfigpolicy",
	Short: "Delete a monitor configuration policy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/monitor/policy/{policy_id}")
		fmt.Println("OperationID: DeleteMonitorConfigPolicy")
	},
}

func init() {
	Cmd.AddCommand(DeleteMonitorConfigPolicyCmd)
}
