package logs_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateLogsMetricCmd = &cobra.Command{
	Use:   "updatelogsmetric",
	Short: "Update a log-based metric",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/logs/config/metrics/{metric_id}")
		fmt.Println("OperationID: UpdateLogsMetric")
	},
}

func init() {
	Cmd.AddCommand(UpdateLogsMetricCmd)
}
