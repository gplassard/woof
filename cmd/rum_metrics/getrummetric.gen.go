package rum_metrics

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetRumMetricCmd = &cobra.Command{
	Use:   "getrummetric [metric_id]",
	Short: "Get a rum-based metric",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewRumMetricsApi(client.NewAPIClient())
		res, _, err := api.GetRumMetric(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getrummetric: %v", err)
		}

		cmdutil.PrintJSON(res, "rum_metrics")
	},
}

func init() {
	Cmd.AddCommand(GetRumMetricCmd)
}
