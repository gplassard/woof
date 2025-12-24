package usage_metering

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetUsageApplicationSecurityMonitoringCmd = &cobra.Command{
	Use:   "getusageapplicationsecuritymonitoring",
	Short: "Get hourly usage for application security",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/usage/application_security")
		fmt.Println("OperationID: GetUsageApplicationSecurityMonitoring")
	},
}

func init() {
	Cmd.AddCommand(GetUsageApplicationSecurityMonitoringCmd)
}
