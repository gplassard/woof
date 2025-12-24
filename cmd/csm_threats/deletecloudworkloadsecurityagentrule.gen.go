package csm_threats

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteCloudWorkloadSecurityAgentRuleCmd = &cobra.Command{
	Use:   "deletecloudworkloadsecurityagentrule",
	Short: "Delete a Workload Protection agent rule (US1-FED)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/security_monitoring/cloud_workload_security/agent_rules/{agent_rule_id}")
		fmt.Println("OperationID: DeleteCloudWorkloadSecurityAgentRule")
	},
}

func init() {
	Cmd.AddCommand(DeleteCloudWorkloadSecurityAgentRuleCmd)
}
