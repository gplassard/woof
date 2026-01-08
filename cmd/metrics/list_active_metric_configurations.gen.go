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

		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListActiveMetricConfigurations(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-active-metric-configurations")

		fmt.Println(cmdutil.FormatJSON(res, "active_metric_configuration"))
	},
}

func init() {

	Cmd.AddCommand(ListActiveMetricConfigurationsCmd)
}
