package fleet_automation

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateFleetDeploymentUpgradeCmd = &cobra.Command{
	Use:   "create-fleet-deployment-upgrade",
	
	Short: "Upgrade hosts",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		res, _, err := api.CreateFleetDeploymentUpgrade(client.NewContext(apiKey, appKey, site), datadogV2.FleetDeploymentPackageUpgradeCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create-fleet-deployment-upgrade: %v", err)
		}

		cmdutil.PrintJSON(res, "deployment")
	},
}

func init() {
	Cmd.AddCommand(CreateFleetDeploymentUpgradeCmd)
}
