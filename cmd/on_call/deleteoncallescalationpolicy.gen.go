package on_call

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteOnCallEscalationPolicyCmd = &cobra.Command{
	Use:   "deleteoncallescalationpolicy",
	Short: "Delete On-Call escalation policy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/on-call/escalation-policies/{policy_id}")
		fmt.Println("OperationID: DeleteOnCallEscalationPolicy")
	},
}

func init() {
	Cmd.AddCommand(DeleteOnCallEscalationPolicyCmd)
}
