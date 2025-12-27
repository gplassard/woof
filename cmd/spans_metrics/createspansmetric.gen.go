package spans_metrics

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateSpansMetricCmd = &cobra.Command{
	Use:   "createspansmetric",
	Short: "Create a span-based metric",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSpansMetricsApi(client.NewAPIClient())
		res, _, err := api.CreateSpansMetric(client.NewContext(apiKey, appKey, site), datadogV2.SpansMetricCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createspansmetric: %v", err)
		}

		cmdutil.PrintJSON(res, "spans_metrics")
	},
}

func init() {
	Cmd.AddCommand(CreateSpansMetricCmd)
}
