package csm_threats

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteCSMThreatsAgentPolicyCmd = &cobra.Command{
	Use:   "deletecsmthreatsagentpolicy",
	Short: "Delete a Workload Protection policy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/remote_config/products/cws/policy/{policy_id}")
		fmt.Println("OperationID: DeleteCSMThreatsAgentPolicy")
	},
}

func init() {
	Cmd.AddCommand(DeleteCSMThreatsAgentPolicyCmd)
}
