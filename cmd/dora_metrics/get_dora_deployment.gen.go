package dora_metrics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetDORADeploymentCmd = &cobra.Command{
	Use: "get-dora-deployment [deployment_id]",

	Short: "Get a deployment event",
	Long: `Get a deployment event
Documentation: https://docs.datadoghq.com/api/latest/dora-metrics/#get-dora-deployment`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DORADeploymentFetchResponse
		var err error

		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetDORADeployment(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-dora-deployment")

		fmt.Println(cmdutil.FormatJSON(res, "dora_deployment"))
	},
}

func init() {

	Cmd.AddCommand(GetDORADeploymentCmd)
}
