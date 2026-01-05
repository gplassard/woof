package dora_metrics

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetDORADeploymentCmd = &cobra.Command{
	Use: "get-d-o-r-ad-eployment [deployment_id]",

	Short: "Get a deployment event",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		res, _, err := api.GetDORADeployment(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-d-o-r-ad-eployment")

		cmd.Println(cmdutil.FormatJSON(res, "dora_metrics"))
	},
}

func init() {
	Cmd.AddCommand(GetDORADeploymentCmd)
}
