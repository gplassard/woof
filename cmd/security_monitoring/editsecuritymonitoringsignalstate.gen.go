package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var EditSecurityMonitoringSignalStateCmd = &cobra.Command{
	Use:   "editsecuritymonitoringsignalstate",
	Short: "Change the triage state of a security signal",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/security_monitoring/signals/{signal_id}/state")
		fmt.Println("OperationID: EditSecurityMonitoringSignalState")
	},
}

func init() {
	Cmd.AddCommand(EditSecurityMonitoringSignalStateCmd)
}
