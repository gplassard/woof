package logs_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateLogsMetricCmd = &cobra.Command{
	Use:   "createlogsmetric",
	Short: "Create a log-based metric",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/logs/config/metrics")
		fmt.Println("OperationID: CreateLogsMetric")
	},
}

func init() {
	Cmd.AddCommand(CreateLogsMetricCmd)
}
