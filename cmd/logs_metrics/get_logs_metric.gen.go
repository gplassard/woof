package logs_metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetLogsMetricCmd = &cobra.Command{
	Use:     "get-logs-metric [metric_id]",
	Aliases: []string{"get"},
	Short:   "Get a log-based metric",
	Long: `Get a log-based metric
Documentation: https://docs.datadoghq.com/api/latest/logs-metrics/#get-logs-metric`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LogsMetricResponse
		var err error

		api := datadogV2.NewLogsMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetLogsMetric(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-logs-metric")

		cmd.Println(cmdutil.FormatJSON(res, "logs_metrics"))
	},
}

func init() {

	Cmd.AddCommand(GetLogsMetricCmd)
}
