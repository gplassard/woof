package csm_threats

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteCSMThreatsAgentRuleCmd = &cobra.Command{
	Use:   "deletecsmthreatsagentrule",
	Short: "Delete a Workload Protection agent rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/remote_config/products/cws/agent_rules/{agent_rule_id}")
		fmt.Println("OperationID: DeleteCSMThreatsAgentRule")
	},
}

func init() {
	Cmd.AddCommand(DeleteCSMThreatsAgentRuleCmd)
}
