package monitors

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetMonitorConfigPolicyCmd = &cobra.Command{
	Use:   "getmonitorconfigpolicy",
	Short: "Get a monitor configuration policy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/monitor/policy/{policy_id}")
		fmt.Println("OperationID: GetMonitorConfigPolicy")
	},
}

func init() {
	Cmd.AddCommand(GetMonitorConfigPolicyCmd)
}
