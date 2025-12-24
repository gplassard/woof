package rum_metrics

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateRumMetricCmd = &cobra.Command{
	Use:   "createrummetric",
	Short: "Create a rum-based metric",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewRumMetricsApi(client.NewAPIClient())
		res, _, err := api.CreateRumMetric(client.NewContext(apiKey, appKey, site), datadogV2.RumMetricCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createrummetric: %v", err)
		}

		cmdutil.PrintJSON(res, "rum_metrics")
	},
}

func init() {
	Cmd.AddCommand(CreateRumMetricCmd)
}
