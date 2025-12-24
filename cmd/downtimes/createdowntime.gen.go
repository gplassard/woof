package downtimes

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateDowntimeCmd = &cobra.Command{
	Use:   "createdowntime",
	Short: "Schedule a downtime",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/downtime")
		fmt.Println("OperationID: CreateDowntime")
	},
}

func init() {
	Cmd.AddCommand(CreateDowntimeCmd)
}
