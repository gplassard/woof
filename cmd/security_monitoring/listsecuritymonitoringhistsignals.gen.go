package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListSecurityMonitoringHistsignalsCmd = &cobra.Command{
	Use:   "listsecuritymonitoringhistsignals",
	Short: "List hist signals",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/siem-threat-hunting/histsignals")
		fmt.Println("OperationID: ListSecurityMonitoringHistsignals")
	},
}

func init() {
	Cmd.AddCommand(ListSecurityMonitoringHistsignalsCmd)
}
