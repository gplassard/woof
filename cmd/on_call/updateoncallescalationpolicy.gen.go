package on_call

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateOnCallEscalationPolicyCmd = &cobra.Command{
	Use:   "updateoncallescalationpolicy",
	Short: "Update On-Call escalation policy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/on-call/escalation-policies/{policy_id}")
		fmt.Println("OperationID: UpdateOnCallEscalationPolicy")
	},
}

func init() {
	Cmd.AddCommand(UpdateOnCallEscalationPolicyCmd)
}
