package logs_metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateLogsMetricCmd = &cobra.Command{
	Use:     "create-logs-metric [payload]",
	Aliases: []string{"create"},
	Short:   "Create a log-based metric",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LogsMetricResponse
		var err error

		var body datadogV2.LogsMetricCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewLogsMetricsApi(client.NewAPIClient())
		res, _, err = api.CreateLogsMetric(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-logs-metric")

		cmd.Println(cmdutil.FormatJSON(res, "logs_metrics"))
	},
}

func init() {
	Cmd.AddCommand(CreateLogsMetricCmd)
}
