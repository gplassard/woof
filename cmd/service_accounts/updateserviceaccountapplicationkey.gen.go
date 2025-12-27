package service_accounts

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateServiceAccountApplicationKeyCmd = &cobra.Command{
	Use:   "updateserviceaccountapplicationkey [service_account_id] [app_key_id]",
	Short: "Edit an application key for this service account",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceAccountsApi(client.NewAPIClient())
		res, _, err := api.UpdateServiceAccountApplicationKey(client.NewContext(apiKey, appKey, site), args[0], args[1], datadogV2.ApplicationKeyUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to updateserviceaccountapplicationkey: %v", err)
		}

		cmdutil.PrintJSON(res, "service_accounts")
	},
}

func init() {
	Cmd.AddCommand(UpdateServiceAccountApplicationKeyCmd)
}
