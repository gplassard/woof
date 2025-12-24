package on_call

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetOnCallEscalationPolicyCmd = &cobra.Command{
	Use:   "getoncallescalationpolicy",
	Short: "Get On-Call escalation policy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/on-call/escalation-policies/{policy_id}")
		fmt.Println("OperationID: GetOnCallEscalationPolicy")
	},
}

func init() {
	Cmd.AddCommand(GetOnCallEscalationPolicyCmd)
}
