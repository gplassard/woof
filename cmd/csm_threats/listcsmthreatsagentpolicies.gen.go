package csm_threats

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListCSMThreatsAgentPoliciesCmd = &cobra.Command{
	Use:   "listcsmthreatsagentpolicies",
	Short: "Get all Workload Protection policies",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/remote_config/products/cws/policy")
		fmt.Println("OperationID: ListCSMThreatsAgentPolicies")
	},
}

func init() {
	Cmd.AddCommand(ListCSMThreatsAgentPoliciesCmd)
}
