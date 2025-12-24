package on_call

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetScheduleOnCallUserCmd = &cobra.Command{
	Use:   "getscheduleoncalluser",
	Short: "Get scheduled on-call user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/on-call/schedules/{schedule_id}/on-call")
		fmt.Println("OperationID: GetScheduleOnCallUser")
	},
}

func init() {
	Cmd.AddCommand(GetScheduleOnCallUserCmd)
}
