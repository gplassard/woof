package dora_metrics

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteDORADeploymentCmd = &cobra.Command{
	Use: "delete-d-o-r-ad-eployment [deployment_id]",

	Short: "Delete a deployment event",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		_, err := api.DeleteDORADeployment(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-d-o-r-ad-eployment")

	},
}

func init() {
	Cmd.AddCommand(DeleteDORADeploymentCmd)
}
