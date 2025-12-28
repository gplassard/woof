package spans_metrics

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteSpansMetricCmd = &cobra.Command{
	Use:   "delete-spans-metric [metric_id]",
	Aliases: []string{ "delete", },
	Short: "Delete a span-based metric",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSpansMetricsApi(client.NewAPIClient())
		_, err := api.DeleteSpansMetric(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete-spans-metric: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteSpansMetricCmd)
}
