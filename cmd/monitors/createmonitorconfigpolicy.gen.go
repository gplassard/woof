package monitors

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateMonitorConfigPolicyCmd = &cobra.Command{
	Use:   "createmonitorconfigpolicy",
	Short: "Create a monitor configuration policy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/monitor/policy")
		fmt.Println("OperationID: CreateMonitorConfigPolicy")
	},
}

func init() {
	Cmd.AddCommand(CreateMonitorConfigPolicyCmd)
}
