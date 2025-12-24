package service_accounts

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListServiceAccountApplicationKeysCmd = &cobra.Command{
	Use:   "listserviceaccountapplicationkeys [service_account_id]",
	Short: "List application keys for this service account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewServiceAccountsApi(client.NewAPIClient())
		res, _, err := api.ListServiceAccountApplicationKeys(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to listserviceaccountapplicationkeys: %v", err)
		}

		cmdutil.PrintJSON(res, "service_accounts")
	},
}

func init() {
	Cmd.AddCommand(ListServiceAccountApplicationKeysCmd)
}
