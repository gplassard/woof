package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetSecurityMonitoringSuppressionCmd = &cobra.Command{
	Use:   "getsecuritymonitoringsuppression",
	Short: "Get a suppression rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/security_monitoring/configuration/suppressions/{suppression_id}")
		fmt.Println("OperationID: GetSecurityMonitoringSuppression")
	},
}

func init() {
	Cmd.AddCommand(GetSecurityMonitoringSuppressionCmd)
}
