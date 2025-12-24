package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetSecurityMonitoringRuleCmd = &cobra.Command{
	Use:   "getsecuritymonitoringrule",
	Short: "Get a rule's details",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/security_monitoring/rules/{rule_id}")
		fmt.Println("OperationID: GetSecurityMonitoringRule")
	},
}

func init() {
	Cmd.AddCommand(GetSecurityMonitoringRuleCmd)
}
