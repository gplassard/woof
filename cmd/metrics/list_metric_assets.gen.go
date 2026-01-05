package metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListMetricAssetsCmd = &cobra.Command{
	Use:     "list-metric-assets [metric_name]",
	Aliases: []string{"list-assets"},
	Short:   "Related Assets to a Metric",
	Long: `Related Assets to a Metric
Documentation: https://docs.datadoghq.com/api/latest/metrics/#list-metric-assets`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MetricAssetsResponse
		var err error

		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err = api.ListMetricAssets(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-metric-assets")

		cmd.Println(cmdutil.FormatJSON(res, "metrics"))
	},
}

func init() {

	Cmd.AddCommand(ListMetricAssetsCmd)
}
