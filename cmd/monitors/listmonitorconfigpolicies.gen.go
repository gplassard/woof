package monitors

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListMonitorConfigPoliciesCmd = &cobra.Command{
	Use:   "listmonitorconfigpolicies",
	Short: "Get all monitor configuration policies",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/monitor/policy")
		fmt.Println("OperationID: ListMonitorConfigPolicies")
	},
}

func init() {
	Cmd.AddCommand(ListMonitorConfigPoliciesCmd)
}
