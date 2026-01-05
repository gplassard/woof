package logs_metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListLogsMetricsCmd = &cobra.Command{
	Use:     "list-logs-metrics",
	Aliases: []string{"list"},
	Short:   "Get all log-based metrics",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LogsMetricsResponse
		var err error

		api := datadogV2.NewLogsMetricsApi(client.NewAPIClient())
		res, _, err = api.ListLogsMetrics(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-logs-metrics")

		cmd.Println(cmdutil.FormatJSON(res, "logs_metrics"))
	},
}

func init() {
	Cmd.AddCommand(ListLogsMetricsCmd)
}
