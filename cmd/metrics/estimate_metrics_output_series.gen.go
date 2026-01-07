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

		optionalParams := datadogV2.NewEstimateMetricsOutputSeriesOptionalParameters()

		if cmd.Flags().Changed("filter-groups") {
			val, _ := cmd.Flags().GetString("filter-groups")
			optionalParams.WithFilterGroups(val)
		}

		if cmd.Flags().Changed("filter-hours-ago") {
			val, _ := cmd.Flags().GetInt64("filter-hours-ago")
			optionalParams.WithFilterHoursAgo(val)
		}

		if cmd.Flags().Changed("filter-num-aggregations") {
			val, _ := cmd.Flags().GetInt64("filter-num-aggregations")
			optionalParams.WithFilterNumAggregations(val)
		}

		if cmd.Flags().Changed("filter-pct") {
			val, _ := cmd.Flags().GetString("filter-pct")
			optionalParams.WithFilterPct(val)
		}

		if cmd.Flags().Changed("filter-timespan-h") {
			val, _ := cmd.Flags().GetInt64("filter-timespan-h")
			optionalParams.WithFilterTimespanH(val)
		}

		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.EstimateMetricsOutputSeries(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to estimate-metrics-output-series")

		cmd.Println(cmdutil.FormatJSON(res, "metric_cardinality_estimate"))
	},
}

func init() {

	EstimateMetricsOutputSeriesCmd.Flags().String("filter-groups", "", "Filtered tag keys that the metric is configured to query with.")

	EstimateMetricsOutputSeriesCmd.Flags().Int64("filter-hours-ago", 0, "The number of hours of look back (from now) to estimate cardinality with. If unspecified, it defaults to 0 hours.")

	EstimateMetricsOutputSeriesCmd.Flags().Int64("filter-num-aggregations", 0, "Deprecated. Number of aggregations has no impact on volume.")

	EstimateMetricsOutputSeriesCmd.Flags().String("filter-pct", "", "A boolean, for distribution metrics only, to estimate cardinality if the metric includes additional percentile aggregators.")

	EstimateMetricsOutputSeriesCmd.Flags().Int64("filter-timespan-h", 0, "A window, in hours, from the look back to estimate cardinality with. The minimum and default is 1 hour.")

	Cmd.AddCommand(EstimateMetricsOutputSeriesCmd)
}
