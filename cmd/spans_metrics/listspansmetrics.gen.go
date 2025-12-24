package spans_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListSpansMetricsCmd = &cobra.Command{
	Use:   "listspansmetrics",
	Short: "Get all span-based metrics",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/apm/config/metrics")
		fmt.Println("OperationID: ListSpansMetrics")
	},
}

func init() {
	Cmd.AddCommand(ListSpansMetricsCmd)
}
