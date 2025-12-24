package on_call

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteOnCallScheduleCmd = &cobra.Command{
	Use:   "deleteoncallschedule",
	Short: "Delete On-Call schedule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/on-call/schedules/{schedule_id}")
		fmt.Println("OperationID: DeleteOnCallSchedule")
	},
}

func init() {
	Cmd.AddCommand(DeleteOnCallScheduleCmd)
}
