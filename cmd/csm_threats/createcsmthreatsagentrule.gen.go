package csm_threats

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateCSMThreatsAgentRuleCmd = &cobra.Command{
	Use:   "createcsmthreatsagentrule",
	Short: "Create a Workload Protection agent rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/remote_config/products/cws/agent_rules")
		fmt.Println("OperationID: CreateCSMThreatsAgentRule")
	},
}

func init() {
	Cmd.AddCommand(CreateCSMThreatsAgentRuleCmd)
}
