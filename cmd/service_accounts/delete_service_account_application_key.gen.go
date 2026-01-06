package service_accounts

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteServiceAccountApplicationKeyCmd = &cobra.Command{
	Use:     "delete-service-account-application-key [service_account_id] [app_key_id]",
	Aliases: []string{"delete-application-key"},
	Short:   "Delete an application key for this service account",
	Long: `Delete an application key for this service account
Documentation: https://docs.datadoghq.com/api/latest/service-accounts/#delete-service-account-application-key`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewServiceAccountsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteServiceAccountApplicationKey(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to delete-service-account-application-key")

	},
}

func init() {

	Cmd.AddCommand(DeleteServiceAccountApplicationKeyCmd)
}
