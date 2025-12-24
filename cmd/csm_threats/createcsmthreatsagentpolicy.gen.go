package csm_threats

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateCSMThreatsAgentPolicyCmd = &cobra.Command{
	Use:   "createcsmthreatsagentpolicy",
	Short: "Create a Workload Protection policy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/remote_config/products/cws/policy")
		fmt.Println("OperationID: CreateCSMThreatsAgentPolicy")
	},
}

func init() {
	Cmd.AddCommand(CreateCSMThreatsAgentPolicyCmd)
}
