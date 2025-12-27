package logs_metrics

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateLogsMetricCmd = &cobra.Command{
	Use:   "update_logs_metric [metric_id]",
	Short: "Update a log-based metric",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsMetricsApi(client.NewAPIClient())
		res, _, err := api.UpdateLogsMetric(client.NewContext(apiKey, appKey, site), args[0], datadogV2.LogsMetricUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to update_logs_metric: %v", err)
		}

		cmdutil.PrintJSON(res, "logs_metrics")
	},
}

func init() {
	Cmd.AddCommand(UpdateLogsMetricCmd)
}
