package deployment_gates

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteDeploymentGateCmd = &cobra.Command{
	Use:   "delete-deployment-gate [id]",
	Aliases: []string{ "delete", },
	Short: "Delete deployment gate",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		_, err := api.DeleteDeploymentGate(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete-deployment-gate: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteDeploymentGateCmd)
}
