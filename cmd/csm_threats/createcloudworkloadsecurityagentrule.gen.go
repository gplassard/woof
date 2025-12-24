package csm_threats

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateCloudWorkloadSecurityAgentRuleCmd = &cobra.Command{
	Use:   "createcloudworkloadsecurityagentrule",
	Short: "Create a Workload Protection agent rule (US1-FED)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/security_monitoring/cloud_workload_security/agent_rules")
		fmt.Println("OperationID: CreateCloudWorkloadSecurityAgentRule")
	},
}

func init() {
	Cmd.AddCommand(CreateCloudWorkloadSecurityAgentRuleCmd)
}
