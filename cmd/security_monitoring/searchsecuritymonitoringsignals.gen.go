package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var SearchSecurityMonitoringSignalsCmd = &cobra.Command{
	Use:   "searchsecuritymonitoringsignals",
	Short: "Get a list of security signals",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/security_monitoring/signals/search")
		fmt.Println("OperationID: SearchSecurityMonitoringSignals")
	},
}

func init() {
	Cmd.AddCommand(SearchSecurityMonitoringSignalsCmd)
}
