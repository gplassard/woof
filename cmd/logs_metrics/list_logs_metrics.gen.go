package logs_metrics

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListLogsMetricsCmd = &cobra.Command{
	Use:   "list-logs-metrics",
	Short: "Get all log-based metrics",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsMetricsApi(client.NewAPIClient())
		res, _, err := api.ListLogsMetrics(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-logs-metrics: %v", err)
		}

		cmdutil.PrintJSON(res, "logs_metrics")
	},
}

func init() {
	Cmd.AddCommand(ListLogsMetricsCmd)
}
