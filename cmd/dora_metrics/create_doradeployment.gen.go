package dora_metrics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateDORADeploymentCmd = &cobra.Command{
	Use: "create-doradeployment",

	Short: "Send a deployment event",
	Long: `Send a deployment event
Documentation: https://docs.datadoghq.com/api/latest/dora-metrics/#create-doradeployment`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DORADeploymentResponse
		var err error

		var body datadogV2.DORADeploymentRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateDORADeployment(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-doradeployment")

		fmt.Println(cmdutil.FormatJSON(res, "d_o_r_a_deployment"))
	},
}

func init() {

	CreateDORADeploymentCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateDORADeploymentCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateDORADeploymentCmd)
}
