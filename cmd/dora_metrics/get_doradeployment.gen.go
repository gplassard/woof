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
	Use: "get-d-o-r-ad-eployment [deployment_id]",

	Short: "Get a deployment event",
	Long: `Get a deployment event
Documentation: https://docs.datadoghq.com/api/latest/d-o-r-a-metrics/#get-d-o-r-ad-eployment`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DORAFetchResponse
		var err error

		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetDORADeployment(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-d-o-r-ad-eployment")

		fmt.Println(cmdutil.FormatJSON(res, "d_o_r_a_deployment"))
	},
}

func init() {

	Cmd.AddCommand(GetDORADeploymentCmd)
}
