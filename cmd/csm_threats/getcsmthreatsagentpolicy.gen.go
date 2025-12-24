package csm_threats

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetCSMThreatsAgentPolicyCmd = &cobra.Command{
	Use:   "getcsmthreatsagentpolicy",
	Short: "Get a Workload Protection policy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/remote_config/products/cws/policy/{policy_id}")
		fmt.Println("OperationID: GetCSMThreatsAgentPolicy")
	},
}

func init() {
	Cmd.AddCommand(GetCSMThreatsAgentPolicyCmd)
}
