package logs_metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateLogsMetricCmd = &cobra.Command{
	Use:     "update-logs-metric [metric_id] [payload]",
	Aliases: []string{"update"},
	Short:   "Update a log-based metric",
	Long: `Update a log-based metric
Documentation: https://docs.datadoghq.com/api/latest/logs-metrics/#update-logs-metric`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LogsMetricResponse
		var err error

		var body datadogV2.LogsMetricUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewLogsMetricsApi(client.NewAPIClient())
		res, _, err = api.UpdateLogsMetric(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-logs-metric")

		cmd.Println(cmdutil.FormatJSON(res, "logs_metrics"))
	},
}

func init() {
	Cmd.AddCommand(UpdateLogsMetricCmd)
}
