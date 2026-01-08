package dora_metrics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListDORADeploymentsCmd = &cobra.Command{
	Use: "list-d-o-r-ad-eployments",

	Short: "Get a list of deployment events",
	Long: `Get a list of deployment events
Documentation: https://docs.datadoghq.com/api/latest/d-o-r-a-metrics/#list-d-o-r-ad-eployments`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DORAListResponse
		var err error

		var body datadogV2.DORAListDeploymentsRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListDORADeployments(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to list-d-o-r-ad-eployments")

		fmt.Println(cmdutil.FormatJSON(res, "d_o_r_a_deployment"))
	},
}

func init() {

	ListDORADeploymentsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	ListDORADeploymentsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(ListDORADeploymentsCmd)
}
