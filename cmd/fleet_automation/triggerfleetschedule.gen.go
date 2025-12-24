package fleet_automation

import (
	"fmt"
	"github.com/spf13/cobra"
)

var TriggerFleetScheduleCmd = &cobra.Command{
	Use:   "triggerfleetschedule",
	Short: "Trigger a schedule deployment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/unstable/fleet/schedules/{id}/trigger")
		fmt.Println("OperationID: TriggerFleetSchedule")
	},
}

func init() {
	Cmd.AddCommand(TriggerFleetScheduleCmd)
}
