package dora_metrics

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateDORADeploymentCmd = &cobra.Command{
	Use:   "create-d-o-r-ad-eployment",
	
	Short: "Send a deployment event",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		res, _, err := api.CreateDORADeployment(client.NewContext(apiKey, appKey, site), datadogV2.DORADeploymentRequest{})
		cmdutil.HandleError(err, "failed to create-d-o-r-ad-eployment")

		cmdutil.PrintJSON(res, "dora_deployment")
	},
}

func init() {
	Cmd.AddCommand(CreateDORADeploymentCmd)
}
