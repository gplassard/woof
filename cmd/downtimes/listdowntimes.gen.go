package downtimes

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListDowntimesCmd = &cobra.Command{
	Use:   "listdowntimes",
	Short: "Get all downtimes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/downtime")
		fmt.Println("OperationID: ListDowntimes")
	},
}

func init() {
	Cmd.AddCommand(ListDowntimesCmd)
}
