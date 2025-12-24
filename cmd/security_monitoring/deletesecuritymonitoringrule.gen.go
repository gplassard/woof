package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteSecurityMonitoringRuleCmd = &cobra.Command{
	Use:   "deletesecuritymonitoringrule",
	Short: "Delete an existing rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/security_monitoring/rules/{rule_id}")
		fmt.Println("OperationID: DeleteSecurityMonitoringRule")
	},
}

func init() {
	Cmd.AddCommand(DeleteSecurityMonitoringRuleCmd)
}
