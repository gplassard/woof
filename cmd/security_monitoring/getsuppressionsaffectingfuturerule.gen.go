package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetSuppressionsAffectingFutureRuleCmd = &cobra.Command{
	Use:   "getsuppressionsaffectingfuturerule",
	Short: "Get suppressions affecting future rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/security_monitoring/configuration/suppressions/rules")
		fmt.Println("OperationID: GetSuppressionsAffectingFutureRule")
	},
}

func init() {
	Cmd.AddCommand(GetSuppressionsAffectingFutureRuleCmd)
}
