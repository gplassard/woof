package logs_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetLogsMetricCmd = &cobra.Command{
	Use:   "getlogsmetric",
	Short: "Get a log-based metric",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/logs/config/metrics/{metric_id}")
		fmt.Println("OperationID: GetLogsMetric")
	},
}

func init() {
	Cmd.AddCommand(GetLogsMetricCmd)
}
