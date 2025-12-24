package spans_metrics

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteSpansMetricCmd = &cobra.Command{
	Use:   "deletespansmetric [metric_id]",
	Short: "Delete a span-based metric",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewSpansMetricsApi(client.NewAPIClient())
		_, err := api.DeleteSpansMetric(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to deletespansmetric: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteSpansMetricCmd)
}
