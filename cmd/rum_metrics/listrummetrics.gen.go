package rum_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListRumMetricsCmd = &cobra.Command{
	Use:   "listrummetrics",
	Short: "Get all rum-based metrics",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/rum/config/metrics")
		fmt.Println("OperationID: ListRumMetrics")
	},
}

func init() {
	Cmd.AddCommand(ListRumMetricsCmd)
}
