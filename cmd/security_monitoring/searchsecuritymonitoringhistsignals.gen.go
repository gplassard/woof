package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var SearchSecurityMonitoringHistsignalsCmd = &cobra.Command{
	Use:   "searchsecuritymonitoringhistsignals",
	Short: "Search hist signals",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/siem-threat-hunting/histsignals/search")
		fmt.Println("OperationID: SearchSecurityMonitoringHistsignals")
	},
}

func init() {
	Cmd.AddCommand(SearchSecurityMonitoringHistsignalsCmd)
}
