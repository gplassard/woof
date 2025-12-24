package metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var EstimateMetricsOutputSeriesCmd = &cobra.Command{
	Use:   "estimatemetricsoutputseries",
	Short: "Tag Configuration Cardinality Estimator",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/metrics/{metric_name}/estimate")
		fmt.Println("OperationID: EstimateMetricsOutputSeries")
	},
}

func init() {
	Cmd.AddCommand(EstimateMetricsOutputSeriesCmd)
}
