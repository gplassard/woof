package fleet_automation

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CancelFleetDeploymentCmd = &cobra.Command{
	Use: "cancel-fleet-deployment [deployment_id]",

	Short: "Cancel a deployment",
	Long: `Cancel a deployment
Documentation: https://docs.datadoghq.com/api/latest/fleet-automation/#cancel-fleet-deployment`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		_, err = api.CancelFleetDeployment(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to cancel-fleet-deployment")

	},
}

func init() {
	Cmd.AddCommand(CancelFleetDeploymentCmd)
}
