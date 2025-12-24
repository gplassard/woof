package rum_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateRumMetricCmd = &cobra.Command{
	Use:   "updaterummetric",
	Short: "Update a rum-based metric",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/rum/config/metrics/{metric_id}")
		fmt.Println("OperationID: UpdateRumMetric")
	},
}

func init() {
	Cmd.AddCommand(UpdateRumMetricCmd)
}
