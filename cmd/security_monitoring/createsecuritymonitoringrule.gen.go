package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateSecurityMonitoringRuleCmd = &cobra.Command{
	Use:   "createsecuritymonitoringrule",
	Short: "Create a detection rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/security_monitoring/rules")
		fmt.Println("OperationID: CreateSecurityMonitoringRule")
	},
}

func init() {
	Cmd.AddCommand(CreateSecurityMonitoringRuleCmd)
}
