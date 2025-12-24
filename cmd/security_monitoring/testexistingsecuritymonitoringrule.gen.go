package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var TestExistingSecurityMonitoringRuleCmd = &cobra.Command{
	Use:   "testexistingsecuritymonitoringrule",
	Short: "Test an existing rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/security_monitoring/rules/{rule_id}/test")
		fmt.Println("OperationID: TestExistingSecurityMonitoringRule")
	},
}

func init() {
	Cmd.AddCommand(TestExistingSecurityMonitoringRuleCmd)
}
