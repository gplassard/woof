package downtimes

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetDowntimeCmd = &cobra.Command{
	Use:   "getdowntime",
	Short: "Get a downtime",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/downtime/{downtime_id}")
		fmt.Println("OperationID: GetDowntime")
	},
}

func init() {
	Cmd.AddCommand(GetDowntimeCmd)
}
