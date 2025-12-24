package fleet_automation

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateFleetScheduleCmd = &cobra.Command{
	Use:   "updatefleetschedule",
	Short: "Update a schedule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/unstable/fleet/schedules/{id}")
		fmt.Println("OperationID: UpdateFleetSchedule")
	},
}

func init() {
	Cmd.AddCommand(UpdateFleetScheduleCmd)
}
