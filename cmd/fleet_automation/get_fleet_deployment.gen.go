package fleet_automation

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetFleetDeploymentCmd = &cobra.Command{
	Use: "get-fleet-deployment [deployment_id]",

	Short: "Get a configuration deployment by ID",
	Long: `Get a configuration deployment by ID
Documentation: https://docs.datadoghq.com/api/latest/fleet-automation/#get-fleet-deployment`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FleetDeploymentResponse
		var err error

		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		res, _, err = api.GetFleetDeployment(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-fleet-deployment")

		cmd.Println(cmdutil.FormatJSON(res, "deployment"))
	},
}

func init() {
	Cmd.AddCommand(GetFleetDeploymentCmd)
}
