package dora_metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteDORADeploymentCmd = &cobra.Command{
	Use: "delete-d-o-r-ad-eployment [deployment_id]",

	Short: "Delete a deployment event",
	Long: `Delete a deployment event
Documentation: https://docs.datadoghq.com/api/latest/dora-metrics/#delete-d-o-r-ad-eployment`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteDORADeployment(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-d-o-r-ad-eployment")

	},
}

func init() {

	Cmd.AddCommand(DeleteDORADeploymentCmd)
}
