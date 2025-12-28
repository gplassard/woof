package fleet_automation

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateFleetDeploymentConfigureCmd = &cobra.Command{
	Use:   "create_fleet_deployment_configure",
	Short: "Create a configuration deployment",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		res, _, err := api.CreateFleetDeploymentConfigure(client.NewContext(apiKey, appKey, site), datadogV2.FleetDeploymentConfigureCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create_fleet_deployment_configure: %v", err)
		}

		cmdutil.PrintJSON(res, "deployment")
	},
}

func init() {
	Cmd.AddCommand(CreateFleetDeploymentConfigureCmd)
}
