package downtimes

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListMonitorDowntimesCmd = &cobra.Command{
	Use:   "listmonitordowntimes",
	Short: "Get active downtimes for a monitor",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/monitor/{monitor_id}/downtime_matches")
		fmt.Println("OperationID: ListMonitorDowntimes")
	},
}

func init() {
	Cmd.AddCommand(ListMonitorDowntimesCmd)
}
