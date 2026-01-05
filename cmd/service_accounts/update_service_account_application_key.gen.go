package service_accounts

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateServiceAccountApplicationKeyCmd = &cobra.Command{
	Use:     "update-service-account-application-key [service_account_id] [app_key_id]",
	Aliases: []string{"update-application-key"},
	Short:   "Edit an application key for this service account",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceAccountsApi(client.NewAPIClient())
		res, _, err := api.UpdateServiceAccountApplicationKey(client.NewContext(apiKey, appKey, site), args[0], args[1], datadogV2.ApplicationKeyUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-service-account-application-key")

		cmd.Println(cmdutil.FormatJSON(res, "application_keys"))
	},
}

func init() {
	Cmd.AddCommand(UpdateServiceAccountApplicationKeyCmd)
}
