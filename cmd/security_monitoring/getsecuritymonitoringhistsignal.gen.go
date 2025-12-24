package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetSecurityMonitoringHistsignalCmd = &cobra.Command{
	Use:   "getsecuritymonitoringhistsignal",
	Short: "Get a hist signal's details",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/siem-threat-hunting/histsignals/{histsignal_id}")
		fmt.Println("OperationID: GetSecurityMonitoringHistsignal")
	},
}

func init() {
	Cmd.AddCommand(GetSecurityMonitoringHistsignalCmd)
}
