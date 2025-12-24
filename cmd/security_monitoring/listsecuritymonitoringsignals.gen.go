package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListSecurityMonitoringSignalsCmd = &cobra.Command{
	Use:   "listsecuritymonitoringsignals",
	Short: "Get a quick list of security signals",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/security_monitoring/signals")
		fmt.Println("OperationID: ListSecurityMonitoringSignals")
	},
}

func init() {
	Cmd.AddCommand(ListSecurityMonitoringSignalsCmd)
}
