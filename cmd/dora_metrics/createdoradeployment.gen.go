package dora_metrics

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateDORADeploymentCmd = &cobra.Command{
	Use:   "createdoradeployment",
	Short: "Send a deployment event",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		res, _, err := api.CreateDORADeployment(client.NewContext(apiKey, appKey, site), datadogV2.DORADeploymentRequest{})
		if err != nil {
			log.Fatalf("failed to createdoradeployment: %v", err)
		}

		cmdutil.PrintJSON(res, "dora_metrics")
	},
}

func init() {
	Cmd.AddCommand(CreateDORADeploymentCmd)
}
