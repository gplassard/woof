package logs

import (
	"fmt"
	"github.com/spf13/cobra"
)

var AggregateLogsCmd = &cobra.Command{
	Use:   "aggregatelogs",
	Short: "Aggregate events",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/logs/analytics/aggregate")
		fmt.Println("OperationID: AggregateLogs")
	},
}

func init() {
	Cmd.AddCommand(AggregateLogsCmd)
}
