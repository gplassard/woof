package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetSuppressionsAffectingRuleCmd = &cobra.Command{
	Use:   "getsuppressionsaffectingrule",
	Short: "Get suppressions affecting a specific rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/security_monitoring/configuration/suppressions/rules/{rule_id}")
		fmt.Println("OperationID: GetSuppressionsAffectingRule")
	},
}

func init() {
	Cmd.AddCommand(GetSuppressionsAffectingRuleCmd)
}
