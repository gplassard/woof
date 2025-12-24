package rum_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetRumMetricCmd = &cobra.Command{
	Use:   "getrummetric",
	Short: "Get a rum-based metric",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/rum/config/metrics/{metric_id}")
		fmt.Println("OperationID: GetRumMetric")
	},
}

func init() {
	Cmd.AddCommand(GetRumMetricCmd)
}
