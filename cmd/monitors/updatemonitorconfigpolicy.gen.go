package monitors

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateMonitorConfigPolicyCmd = &cobra.Command{
	Use:   "updatemonitorconfigpolicy",
	Short: "Edit a monitor configuration policy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/monitor/policy/{policy_id}")
		fmt.Println("OperationID: UpdateMonitorConfigPolicy")
	},
}

func init() {
	Cmd.AddCommand(UpdateMonitorConfigPolicyCmd)
}
