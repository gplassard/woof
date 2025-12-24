package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ConvertExistingSecurityMonitoringRuleCmd = &cobra.Command{
	Use:   "convertexistingsecuritymonitoringrule",
	Short: "Convert an existing rule from JSON to Terraform",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/security_monitoring/rules/{rule_id}/convert")
		fmt.Println("OperationID: ConvertExistingSecurityMonitoringRule")
	},
}

func init() {
	Cmd.AddCommand(ConvertExistingSecurityMonitoringRuleCmd)
}
