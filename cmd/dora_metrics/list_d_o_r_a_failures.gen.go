package dora_metrics

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListDORAFailuresCmd = &cobra.Command{
	Use:   "list-d-o-r-a-failures",
	
	Short: "Get a list of failure events",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		res, _, err := api.ListDORAFailures(client.NewContext(apiKey, appKey, site), datadogV2.DORAListFailuresRequest{})
		if err != nil {
			log.Fatalf("failed to list-d-o-r-a-failures: %v", err)
		}

		cmdutil.PrintJSON(res, "dora_metrics")
	},
}

func init() {
	Cmd.AddCommand(ListDORAFailuresCmd)
}
