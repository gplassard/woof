package service_accounts

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetServiceAccountApplicationKeyCmd = &cobra.Command{
	Use:   "get-service-account-application-key [service_account_id] [app_key_id]",
	Aliases: []string{ "get-application-key", },
	Short: "Get one application key for this service account",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceAccountsApi(client.NewAPIClient())
		res, _, err := api.GetServiceAccountApplicationKey(client.NewContext(apiKey, appKey, site), args[0], args[1])
		if err != nil {
			log.Fatalf("failed to get-service-account-application-key: %v", err)
		}

		cmdutil.PrintJSON(res, "application_keys")
	},
}

func init() {
	Cmd.AddCommand(GetServiceAccountApplicationKeyCmd)
}
