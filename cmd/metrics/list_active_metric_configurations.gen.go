package metrics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListActiveMetricConfigurationsCmd = &cobra.Command{
	Use:     "list-active-metric-configurations [metric_name]",
	Aliases: []string{"list-active-configurations"},
	Short:   "List active tags and aggregations",
	Long: `List active tags and aggregations
Documentation: https://docs.datadoghq.com/api/latest/metrics/#list-active-metric-configurations`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MetricSuggestedTagsAndAggregationsResponse
		var err error

		optionalParams := datadogV2.NewListActiveMetricConfigurationsOptionalParameters()

		if cmd.Flags().Changed("window-seconds") {
			val, _ := cmd.Flags().GetInt64("window-seconds")
			optionalParams.WithWindowSeconds(val)
		}

		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListActiveMetricConfigurations(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to list-active-metric-configurations")

		fmt.Println(cmdutil.FormatJSON(res, "actively_queried_configurations"))
	},
}

func init() {

	ListActiveMetricConfigurationsCmd.Flags().Int64("window-seconds", 0, "The number of seconds of look back (from now). Default value is 604,800 (1 week), minimum value is 7200 (2 hours), maximum value is 2,630,000 (1 month).")

	Cmd.AddCommand(ListActiveMetricConfigurationsCmd)
}
