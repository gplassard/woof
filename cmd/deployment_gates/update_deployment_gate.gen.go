package deployment_gates

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateDeploymentGateCmd = &cobra.Command{
	Use:     "update-deployment-gate [id] [payload]",
	Aliases: []string{"update"},
	Short:   "Update deployment gate",
	Long: `Update deployment gate
Documentation: https://docs.datadoghq.com/api/latest/deployment-gates/#update-deployment-gate`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DeploymentGateResponse
		var err error

		var body datadogV2.UpdateDeploymentGateParams
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		res, _, err = api.UpdateDeploymentGate(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-deployment-gate")

		cmd.Println(cmdutil.FormatJSON(res, "deployment_gate"))
	},
}

func init() {
	Cmd.AddCommand(UpdateDeploymentGateCmd)
}
