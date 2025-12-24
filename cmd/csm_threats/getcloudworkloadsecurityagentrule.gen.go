package csm_threats

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetCloudWorkloadSecurityAgentRuleCmd = &cobra.Command{
	Use:   "getcloudworkloadsecurityagentrule",
	Short: "Get a Workload Protection agent rule (US1-FED)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/security_monitoring/cloud_workload_security/agent_rules/{agent_rule_id}")
		fmt.Println("OperationID: GetCloudWorkloadSecurityAgentRule")
	},
}

func init() {
	Cmd.AddCommand(GetCloudWorkloadSecurityAgentRuleCmd)
}
