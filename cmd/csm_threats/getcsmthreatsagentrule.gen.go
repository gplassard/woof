package csm_threats

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetCSMThreatsAgentRuleCmd = &cobra.Command{
	Use:   "getcsmthreatsagentrule",
	Short: "Get a Workload Protection agent rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/remote_config/products/cws/agent_rules/{agent_rule_id}")
		fmt.Println("OperationID: GetCSMThreatsAgentRule")
	},
}

func init() {
	Cmd.AddCommand(GetCSMThreatsAgentRuleCmd)
}
