package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateSecurityMonitoringSuppressionCmd = &cobra.Command{
	Use:   "updatesecuritymonitoringsuppression",
	Short: "Update a suppression rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/security_monitoring/configuration/suppressions/{suppression_id}")
		fmt.Println("OperationID: UpdateSecurityMonitoringSuppression")
	},
}

func init() {
	Cmd.AddCommand(UpdateSecurityMonitoringSuppressionCmd)
}
