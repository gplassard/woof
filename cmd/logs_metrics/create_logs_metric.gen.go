package logs_metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateLogsMetricCmd = &cobra.Command{
	Use:     "create-logs-metric",
	Aliases: []string{"create"},
	Short:   "Create a log-based metric",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsMetricsApi(client.NewAPIClient())
		res, _, err := api.CreateLogsMetric(client.NewContext(apiKey, appKey, site), datadogV2.LogsMetricCreateRequest{})
		cmdutil.HandleError(err, "failed to create-logs-metric")

		cmd.Println(cmdutil.FormatJSON(res, "logs_metrics"))
	},
}

func init() {
	Cmd.AddCommand(CreateLogsMetricCmd)
}
