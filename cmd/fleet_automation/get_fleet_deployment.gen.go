package fleet_automation

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetFleetDeploymentCmd = &cobra.Command{
	Use:   "get-fleet-deployment [deployment_id]",
	
	Short: "Get a configuration deployment by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		res, _, err := api.GetFleetDeployment(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-fleet-deployment")

		cmdutil.PrintJSON(res, "deployment")
	},
}

func init() {
	Cmd.AddCommand(GetFleetDeploymentCmd)
}
