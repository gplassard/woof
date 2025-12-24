package on_call

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateOnCallScheduleCmd = &cobra.Command{
	Use:   "updateoncallschedule",
	Short: "Update On-Call schedule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/on-call/schedules/{schedule_id}")
		fmt.Println("OperationID: UpdateOnCallSchedule")
	},
}

func init() {
	Cmd.AddCommand(UpdateOnCallScheduleCmd)
}
