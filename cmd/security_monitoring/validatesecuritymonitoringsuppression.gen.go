package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ValidateSecurityMonitoringSuppressionCmd = &cobra.Command{
	Use:   "validatesecuritymonitoringsuppression",
	Short: "Validate a suppression rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/security_monitoring/configuration/suppressions/validation")
		fmt.Println("OperationID: ValidateSecurityMonitoringSuppression")
	},
}

func init() {
	Cmd.AddCommand(ValidateSecurityMonitoringSuppressionCmd)
}
