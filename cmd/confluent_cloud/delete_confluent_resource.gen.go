package confluent_cloud

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteConfluentResourceCmd = &cobra.Command{
	Use:   "delete_confluent_resource [account_id] [resource_id]",
	Short: "Delete resource from Confluent account",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		_, err := api.DeleteConfluentResource(client.NewContext(apiKey, appKey, site), args[0], args[1])
		if err != nil {
			log.Fatalf("failed to delete_confluent_resource: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteConfluentResourceCmd)
}
