package csm_threats

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateCloudWorkloadSecurityAgentRuleCmd = &cobra.Command{
	Use:   "updatecloudworkloadsecurityagentrule",
	Short: "Update a Workload Protection agent rule (US1-FED)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/security_monitoring/cloud_workload_security/agent_rules/{agent_rule_id}")
		fmt.Println("OperationID: UpdateCloudWorkloadSecurityAgentRule")
	},
}

func init() {
	Cmd.AddCommand(UpdateCloudWorkloadSecurityAgentRuleCmd)
}
