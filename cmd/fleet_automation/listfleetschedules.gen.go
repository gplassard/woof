package fleet_automation

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListFleetSchedulesCmd = &cobra.Command{
	Use:   "listfleetschedules",
	Short: "List all schedules",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/unstable/fleet/schedules")
		fmt.Println("OperationID: ListFleetSchedules")
	},
}

func init() {
	Cmd.AddCommand(ListFleetSchedulesCmd)
}
