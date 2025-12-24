package fleet_automation

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CancelFleetDeploymentCmd = &cobra.Command{
	Use:   "cancelfleetdeployment [deployment_id]",
	Short: "Cancel a deployment",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		_, err := api.CancelFleetDeployment(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to cancelfleetdeployment: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(CancelFleetDeploymentCmd)
}
