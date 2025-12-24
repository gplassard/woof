package spans_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateSpansMetricCmd = &cobra.Command{
	Use:   "updatespansmetric",
	Short: "Update a span-based metric",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/apm/config/metrics/{metric_id}")
		fmt.Println("OperationID: UpdateSpansMetric")
	},
}

func init() {
	Cmd.AddCommand(UpdateSpansMetricCmd)
}
