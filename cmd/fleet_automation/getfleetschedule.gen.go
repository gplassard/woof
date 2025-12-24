package fleet_automation

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetFleetScheduleCmd = &cobra.Command{
	Use:   "getfleetschedule",
	Short: "Get a schedule by ID",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/unstable/fleet/schedules/{id}")
		fmt.Println("OperationID: GetFleetSchedule")
	},
}

func init() {
	Cmd.AddCommand(GetFleetScheduleCmd)
}
