package deployment_gates

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateDeploymentGateCmd = &cobra.Command{
	Use:     "update-deployment-gate [id]",
	Aliases: []string{"update"},
	Short:   "Update deployment gate",
	Long: `Update deployment gate
Documentation: https://docs.datadoghq.com/api/latest/deployment-gates/#update-deployment-gate`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DeploymentGateResponse
		var err error

		var body datadogV2.UpdateDeploymentGateParams
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateDeploymentGate(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-deployment-gate")

		fmt.Println(cmdutil.FormatJSON(res, "deployment_gate"))
	},
}

func init() {

	UpdateDeploymentGateCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateDeploymentGateCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateDeploymentGateCmd)
}
