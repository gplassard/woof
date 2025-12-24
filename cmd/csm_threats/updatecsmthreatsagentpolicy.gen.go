package csm_threats

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateCSMThreatsAgentPolicyCmd = &cobra.Command{
	Use:   "updatecsmthreatsagentpolicy",
	Short: "Update a Workload Protection policy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/remote_config/products/cws/policy/{policy_id}")
		fmt.Println("OperationID: UpdateCSMThreatsAgentPolicy")
	},
}

func init() {
	Cmd.AddCommand(UpdateCSMThreatsAgentPolicyCmd)
}
