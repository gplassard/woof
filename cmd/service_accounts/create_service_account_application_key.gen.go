package service_accounts

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateServiceAccountApplicationKeyCmd = &cobra.Command{
	Use:   "create_service_account_application_key [service_account_id]",
	Short: "Create an application key for this service account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceAccountsApi(client.NewAPIClient())
		res, _, err := api.CreateServiceAccountApplicationKey(client.NewContext(apiKey, appKey, site), args[0], datadogV2.ApplicationKeyCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create_service_account_application_key: %v", err)
		}

		cmdutil.PrintJSON(res, "application_keys")
	},
}

func init() {
	Cmd.AddCommand(CreateServiceAccountApplicationKeyCmd)
}
