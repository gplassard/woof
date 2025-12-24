package fleet_automation

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteFleetScheduleCmd = &cobra.Command{
	Use:   "deletefleetschedule",
	Short: "Delete a schedule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/unstable/fleet/schedules/{id}")
		fmt.Println("OperationID: DeleteFleetSchedule")
	},
}

func init() {
	Cmd.AddCommand(DeleteFleetScheduleCmd)
}
