package rum_metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetRumMetricCmd = &cobra.Command{
	Use:     "get-rum-metric [metric_id]",
	Aliases: []string{"get"},
	Short:   "Get a rum-based metric",
	Long: `Get a rum-based metric
Documentation: https://docs.datadoghq.com/api/latest/rum-metrics/#get-rum-metric`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RumMetricResponse
		var err error

		api := datadogV2.NewRumMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetRumMetric(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-rum-metric")

		cmd.Println(cmdutil.FormatJSON(res, "rum_metrics"))
	},
}

func init() {

	Cmd.AddCommand(GetRumMetricCmd)
}
