package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListSecurityMonitoringSuppressionsCmd = &cobra.Command{
	Use:   "listsecuritymonitoringsuppressions",
	Short: "Get all suppression rules",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/security_monitoring/configuration/suppressions")
		fmt.Println("OperationID: ListSecurityMonitoringSuppressions")
	},
}

func init() {
	Cmd.AddCommand(ListSecurityMonitoringSuppressionsCmd)
}
