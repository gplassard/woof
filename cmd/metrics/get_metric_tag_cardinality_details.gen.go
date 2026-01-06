package metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetMetricTagCardinalityDetailsCmd = &cobra.Command{
	Use:     "get-metric-tag-cardinality-details [metric_name]",
	Aliases: []string{"get-tag-cardinality-details"},
	Short:   "Get tag key cardinality details",
	Long: `Get tag key cardinality details
Documentation: https://docs.datadoghq.com/api/latest/metrics/#get-metric-tag-cardinality-details`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MetricTagCardinalitiesResponse
		var err error

		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetMetricTagCardinalityDetails(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-metric-tag-cardinality-details")

		cmd.Println(cmdutil.FormatJSON(res, "metrics"))
	},
}

func init() {

	Cmd.AddCommand(GetMetricTagCardinalityDetailsCmd)
}
