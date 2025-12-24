package csm_threats

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListCSMThreatsAgentRulesCmd = &cobra.Command{
	Use:   "listcsmthreatsagentrules",
	Short: "Get all Workload Protection agent rules",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/remote_config/products/cws/agent_rules")
		fmt.Println("OperationID: ListCSMThreatsAgentRules")
	},
}

func init() {
	Cmd.AddCommand(ListCSMThreatsAgentRulesCmd)
}
