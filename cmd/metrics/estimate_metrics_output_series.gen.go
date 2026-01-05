package metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var EstimateMetricsOutputSeriesCmd = &cobra.Command{
	Use:     "estimate-metrics-output-series [metric_name]",
	Aliases: []string{"estimate-output-series"},
	Short:   "Tag Configuration Cardinality Estimator",
	Long: `Tag Configuration Cardinality Estimator
Documentation: https://docs.datadoghq.com/api/latest/metrics/#estimate-metrics-output-series`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MetricEstimateResponse
		var err error

		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err = api.EstimateMetricsOutputSeries(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to estimate-metrics-output-series")

		cmd.Println(cmdutil.FormatJSON(res, "metric_cardinality_estimate"))
	},
}

func init() {

	Cmd.AddCommand(EstimateMetricsOutputSeriesCmd)
}
