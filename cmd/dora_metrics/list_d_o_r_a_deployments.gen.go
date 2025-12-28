package dora_metrics

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListDORADeploymentsCmd = &cobra.Command{
	Use:   "list-d-o-r-a-deployments",
	
	Short: "Get a list of deployment events",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		res, _, err := api.ListDORADeployments(client.NewContext(apiKey, appKey, site), datadogV2.DORAListDeploymentsRequest{})
		if err != nil {
			log.Fatalf("failed to list-d-o-r-a-deployments: %v", err)
		}

		cmdutil.PrintJSON(res, "dora_metrics")
	},
}

func init() {
	Cmd.AddCommand(ListDORADeploymentsCmd)
}
