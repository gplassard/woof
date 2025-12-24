package on_call

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetOnCallScheduleCmd = &cobra.Command{
	Use:   "getoncallschedule",
	Short: "Get On-Call schedule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/on-call/schedules/{schedule_id}")
		fmt.Println("OperationID: GetOnCallSchedule")
	},
}

func init() {
	Cmd.AddCommand(GetOnCallScheduleCmd)
}
