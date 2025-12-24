package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListSecurityMonitoringRulesCmd = &cobra.Command{
	Use:   "listsecuritymonitoringrules",
	Short: "List rules",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/security_monitoring/rules")
		fmt.Println("OperationID: ListSecurityMonitoringRules")
	},
}

func init() {
	Cmd.AddCommand(ListSecurityMonitoringRulesCmd)
}
