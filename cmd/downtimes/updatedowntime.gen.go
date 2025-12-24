package downtimes

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateDowntimeCmd = &cobra.Command{
	Use:   "updatedowntime",
	Short: "Update a downtime",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/downtime/{downtime_id}")
		fmt.Println("OperationID: UpdateDowntime")
	},
}

func init() {
	Cmd.AddCommand(UpdateDowntimeCmd)
}
