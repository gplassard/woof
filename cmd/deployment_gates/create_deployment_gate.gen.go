package deployment_gates

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateDeploymentGateCmd = &cobra.Command{
	Use:     "create-deployment-gate",
	Aliases: []string{"create"},
	Short:   "Create deployment gate",
	Long: `Create deployment gate
Documentation: https://docs.datadoghq.com/api/latest/deployment-gates/#create-deployment-gate`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DeploymentGateResponse
		var err error

		var body datadogV2.CreateDeploymentGateParams
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		res, _, err = api.CreateDeploymentGate(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-deployment-gate")

		cmd.Println(cmdutil.FormatJSON(res, "deployment_gate"))
	},
}

func init() {

	CreateDeploymentGateCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateDeploymentGateCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateDeploymentGateCmd)
}
