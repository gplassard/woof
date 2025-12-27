package service_accounts

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteServiceAccountApplicationKeyCmd = &cobra.Command{
	Use:   "deleteserviceaccountapplicationkey [service_account_id] [app_key_id]",
	Short: "Delete an application key for this service account",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceAccountsApi(client.NewAPIClient())
		_, err := api.DeleteServiceAccountApplicationKey(client.NewContext(apiKey, appKey, site), args[0], args[1])
		if err != nil {
			log.Fatalf("failed to deleteserviceaccountapplicationkey: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteServiceAccountApplicationKeyCmd)
}
