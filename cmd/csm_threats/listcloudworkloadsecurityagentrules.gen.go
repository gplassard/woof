package csm_threats

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListCloudWorkloadSecurityAgentRulesCmd = &cobra.Command{
	Use:   "listcloudworkloadsecurityagentrules",
	Short: "Get all Workload Protection agent rules (US1-FED)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/security_monitoring/cloud_workload_security/agent_rules")
		fmt.Println("OperationID: ListCloudWorkloadSecurityAgentRules")
	},
}

func init() {
	Cmd.AddCommand(ListCloudWorkloadSecurityAgentRulesCmd)
}
