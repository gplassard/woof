package spans_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetSpansMetricCmd = &cobra.Command{
	Use:   "getspansmetric",
	Short: "Get a span-based metric",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/apm/config/metrics/{metric_id}")
		fmt.Println("OperationID: GetSpansMetric")
	},
}

func init() {
	Cmd.AddCommand(GetSpansMetricCmd)
}
