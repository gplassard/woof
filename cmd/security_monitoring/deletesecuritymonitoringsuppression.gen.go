package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteSecurityMonitoringSuppressionCmd = &cobra.Command{
	Use:   "deletesecuritymonitoringsuppression",
	Short: "Delete a suppression rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/security_monitoring/configuration/suppressions/{suppression_id}")
		fmt.Println("OperationID: DeleteSecurityMonitoringSuppression")
	},
}

func init() {
	Cmd.AddCommand(DeleteSecurityMonitoringSuppressionCmd)
}
