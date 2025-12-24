package fleet_automation

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateFleetScheduleCmd = &cobra.Command{
	Use:   "createfleetschedule",
	Short: "Create a schedule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/unstable/fleet/schedules")
		fmt.Println("OperationID: CreateFleetSchedule")
	},
}

func init() {
	Cmd.AddCommand(CreateFleetScheduleCmd)
}
