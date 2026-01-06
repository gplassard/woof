package deployment_gates

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetDeploymentGateCmd = &cobra.Command{
	Use:     "get-deployment-gate [id]",
	Aliases: []string{"get"},
	Short:   "Get deployment gate",
	Long: `Get deployment gate
Documentation: https://docs.datadoghq.com/api/latest/deployment-gates/#get-deployment-gate`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DeploymentGateResponse
		var err error

		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetDeploymentGate(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-deployment-gate")

		cmd.Println(cmdutil.FormatJSON(res, "deployment_gate"))
	},
}

func init() {

	Cmd.AddCommand(GetDeploymentGateCmd)
}
