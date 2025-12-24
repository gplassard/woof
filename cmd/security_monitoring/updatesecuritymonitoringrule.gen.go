package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateSecurityMonitoringRuleCmd = &cobra.Command{
	Use:   "updatesecuritymonitoringrule",
	Short: "Update an existing rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/security_monitoring/rules/{rule_id}")
		fmt.Println("OperationID: UpdateSecurityMonitoringRule")
	},
}

func init() {
	Cmd.AddCommand(UpdateSecurityMonitoringRuleCmd)
}
