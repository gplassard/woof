package spans_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateSpansMetricCmd = &cobra.Command{
	Use:   "createspansmetric",
	Short: "Create a span-based metric",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/apm/config/metrics")
		fmt.Println("OperationID: CreateSpansMetric")
	},
}

func init() {
	Cmd.AddCommand(CreateSpansMetricCmd)
}
