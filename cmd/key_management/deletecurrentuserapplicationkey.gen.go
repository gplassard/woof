package key_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteCurrentUserApplicationKeyCmd = &cobra.Command{
	Use:   "deletecurrentuserapplicationkey [app_key_id]",
	Short: "Delete an application key owned by current user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		_, err := api.DeleteCurrentUserApplicationKey(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to deletecurrentuserapplicationkey: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteCurrentUserApplicationKeyCmd)
}
