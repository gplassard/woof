package csm_threats

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateCSMThreatsAgentRuleCmd = &cobra.Command{
	Use:   "updatecsmthreatsagentrule",
	Short: "Update a Workload Protection agent rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/remote_config/products/cws/agent_rules/{agent_rule_id}")
		fmt.Println("OperationID: UpdateCSMThreatsAgentRule")
	},
}

func init() {
	Cmd.AddCommand(UpdateCSMThreatsAgentRuleCmd)
}
