package on_call

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateOnCallScheduleCmd = &cobra.Command{
	Use:   "createoncallschedule",
	Short: "Create On-Call schedule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/on-call/schedules")
		fmt.Println("OperationID: CreateOnCallSchedule")
	},
}

func init() {
	Cmd.AddCommand(CreateOnCallScheduleCmd)
}
