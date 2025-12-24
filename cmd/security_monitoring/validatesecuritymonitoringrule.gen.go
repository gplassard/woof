package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ValidateSecurityMonitoringRuleCmd = &cobra.Command{
	Use:   "validatesecuritymonitoringrule",
	Short: "Validate a detection rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/security_monitoring/rules/validation")
		fmt.Println("OperationID: ValidateSecurityMonitoringRule")
	},
}

func init() {
	Cmd.AddCommand(ValidateSecurityMonitoringRuleCmd)
}
