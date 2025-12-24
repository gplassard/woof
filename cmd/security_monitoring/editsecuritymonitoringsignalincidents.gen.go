package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var EditSecurityMonitoringSignalIncidentsCmd = &cobra.Command{
	Use:   "editsecuritymonitoringsignalincidents",
	Short: "Change the related incidents of a security signal",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/security_monitoring/signals/{signal_id}/incidents")
		fmt.Println("OperationID: EditSecurityMonitoringSignalIncidents")
	},
}

func init() {
	Cmd.AddCommand(EditSecurityMonitoringSignalIncidentsCmd)
}
