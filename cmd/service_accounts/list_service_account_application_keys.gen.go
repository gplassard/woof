package service_accounts

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListServiceAccountApplicationKeysCmd = &cobra.Command{
	Use:     "list-service-account-application-keys [service_account_id]",
	Aliases: []string{"list-application-keys"},
	Short:   "List application keys for this service account",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceAccountsApi(client.NewAPIClient())
		res, _, err := api.ListServiceAccountApplicationKeys(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-service-account-application-keys")

		cmd.Println(cmdutil.FormatJSON(res, "application_keys"))
	},
}

func init() {
	Cmd.AddCommand(ListServiceAccountApplicationKeysCmd)
}
