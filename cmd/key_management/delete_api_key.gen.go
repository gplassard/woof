package key_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteAPIKeyCmd = &cobra.Command{
	Use:   "delete-api-key [api_key_id]",
	
	Short: "Delete an API key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		_, err := api.DeleteAPIKey(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete-api-key: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteAPIKeyCmd)
}
