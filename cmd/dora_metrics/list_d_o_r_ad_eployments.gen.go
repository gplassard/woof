package dora_metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListDORADeploymentsCmd = &cobra.Command{
	Use: "list-d-o-r-ad-eployments",

	Short: "Get a list of deployment events",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		res, _, err := api.ListDORADeployments(client.NewContext(apiKey, appKey, site), datadogV2.DORAListDeploymentsRequest{})
		cmdutil.HandleError(err, "failed to list-d-o-r-ad-eployments")

		cmd.Println(cmdutil.FormatJSON(res, "dora_metrics"))
	},
}

func init() {
	Cmd.AddCommand(ListDORADeploymentsCmd)
}
