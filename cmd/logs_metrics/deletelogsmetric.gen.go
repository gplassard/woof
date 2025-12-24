package logs_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteLogsMetricCmd = &cobra.Command{
	Use:   "deletelogsmetric",
	Short: "Delete a log-based metric",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/logs/config/metrics/{metric_id}")
		fmt.Println("OperationID: DeleteLogsMetric")
	},
}

func init() {
	Cmd.AddCommand(DeleteLogsMetricCmd)
}
