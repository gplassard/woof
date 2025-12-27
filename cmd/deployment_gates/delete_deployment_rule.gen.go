package deployment_gates

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteDeploymentRuleCmd = &cobra.Command{
	Use:   "delete_deployment_rule [gate_id] [id]",
	Short: "Delete deployment rule",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		_, err := api.DeleteDeploymentRule(client.NewContext(apiKey, appKey, site), args[0], args[1])
		if err != nil {
			log.Fatalf("failed to delete_deployment_rule: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteDeploymentRuleCmd)
}
