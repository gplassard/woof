package downtimes

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CancelDowntimeCmd = &cobra.Command{
	Use:   "canceldowntime",
	Short: "Cancel a downtime",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/downtime/{downtime_id}")
		fmt.Println("OperationID: CancelDowntime")
	},
}

func init() {
	Cmd.AddCommand(CancelDowntimeCmd)
}
