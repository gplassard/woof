package spans_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteSpansMetricCmd = &cobra.Command{
	Use:   "deletespansmetric",
	Short: "Delete a span-based metric",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/apm/config/metrics/{metric_id}")
		fmt.Println("OperationID: DeleteSpansMetric")
	},
}

func init() {
	Cmd.AddCommand(DeleteSpansMetricCmd)
}
