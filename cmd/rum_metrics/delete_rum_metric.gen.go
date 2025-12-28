package rum_metrics

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteRumMetricCmd = &cobra.Command{
	Use:   "delete-rum-metric [metric_id]",
	Short: "Delete a rum-based metric",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRumMetricsApi(client.NewAPIClient())
		_, err := api.DeleteRumMetric(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete-rum-metric: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteRumMetricCmd)
}
