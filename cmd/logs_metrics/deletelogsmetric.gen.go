package logs_metrics

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteLogsMetricCmd = &cobra.Command{
	Use:   "deletelogsmetric [metric_id]",
	Short: "Delete a log-based metric",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewLogsMetricsApi(client.NewAPIClient())
		_, err := api.DeleteLogsMetric(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to deletelogsmetric: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteLogsMetricCmd)
}
