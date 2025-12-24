package on_call

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateOnCallEscalationPolicyCmd = &cobra.Command{
	Use:   "createoncallescalationpolicy",
	Short: "Create On-Call escalation policy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/on-call/escalation-policies")
		fmt.Println("OperationID: CreateOnCallEscalationPolicy")
	},
}

func init() {
	Cmd.AddCommand(CreateOnCallEscalationPolicyCmd)
}
