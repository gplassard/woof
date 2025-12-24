package logs_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListLogsMetricsCmd = &cobra.Command{
	Use:   "listlogsmetrics",
	Short: "Get all log-based metrics",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/logs/config/metrics")
		fmt.Println("OperationID: ListLogsMetrics")
	},
}

func init() {
	Cmd.AddCommand(ListLogsMetricsCmd)
}
