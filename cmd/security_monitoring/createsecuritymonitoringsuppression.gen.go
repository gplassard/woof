package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateSecurityMonitoringSuppressionCmd = &cobra.Command{
	Use:   "createsecuritymonitoringsuppression",
	Short: "Create a suppression rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/security_monitoring/configuration/suppressions")
		fmt.Println("OperationID: CreateSecurityMonitoringSuppression")
	},
}

func init() {
	Cmd.AddCommand(CreateSecurityMonitoringSuppressionCmd)
}
