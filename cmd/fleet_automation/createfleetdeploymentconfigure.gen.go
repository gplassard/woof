package fleet_automation

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateFleetDeploymentConfigureCmd = &cobra.Command{
	Use:   "createfleetdeploymentconfigure",
	Short: "Create a configuration deployment",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		res, _, err := api.CreateFleetDeploymentConfigure(client.NewContext(apiKey, appKey, site), datadogV2.FleetDeploymentConfigureCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createfleetdeploymentconfigure: %v", err)
		}

		cmdutil.PrintJSON(res, "fleet_automation")
	},
}

func init() {
	Cmd.AddCommand(CreateFleetDeploymentConfigureCmd)
}
