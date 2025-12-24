package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var TestSecurityMonitoringRuleCmd = &cobra.Command{
	Use:   "testsecuritymonitoringrule",
	Short: "Test a rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/security_monitoring/rules/test")
		fmt.Println("OperationID: TestSecurityMonitoringRule")
	},
}

func init() {
	Cmd.AddCommand(TestSecurityMonitoringRuleCmd)
}
