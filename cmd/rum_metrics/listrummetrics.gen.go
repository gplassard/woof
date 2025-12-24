package rum_metrics

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListRumMetricsCmd = &cobra.Command{
	Use:   "listrummetrics",
	Short: "Get all rum-based metrics",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewRumMetricsApi(client.NewAPIClient())
		res, _, err := api.ListRumMetrics(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listrummetrics: %v", err)
		}

		cmdutil.PrintJSON(res, "rum_metrics")
	},
}

func init() {
	Cmd.AddCommand(ListRumMetricsCmd)
}
