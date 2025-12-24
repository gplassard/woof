package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetSecurityMonitoringSignalCmd = &cobra.Command{
	Use:   "getsecuritymonitoringsignal",
	Short: "Get a signal's details",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/security_monitoring/signals/{signal_id}")
		fmt.Println("OperationID: GetSecurityMonitoringSignal")
	},
}

func init() {
	Cmd.AddCommand(GetSecurityMonitoringSignalCmd)
}
