package service_accounts

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetServiceAccountApplicationKeyCmd = &cobra.Command{
	Use:     "get-service-account-application-key [service_account_id] [app_key_id]",
	Aliases: []string{"get-application-key"},
	Short:   "Get one application key for this service account",
	Long: `Get one application key for this service account
Documentation: https://docs.datadoghq.com/api/latest/service-accounts/#get-service-account-application-key`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.PartialApplicationKeyResponse
		var err error

		api := datadogV2.NewServiceAccountsApi(client.NewAPIClient())
		res, _, err = api.GetServiceAccountApplicationKey(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to get-service-account-application-key")

		cmd.Println(cmdutil.FormatJSON(res, "application_keys"))
	},
}

func init() {
	Cmd.AddCommand(GetServiceAccountApplicationKeyCmd)
}
