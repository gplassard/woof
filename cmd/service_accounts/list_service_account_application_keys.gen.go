package service_accounts

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListServiceAccountApplicationKeysCmd = &cobra.Command{
	Use:     "list-service-account-application-keys [service_account_id]",
	Aliases: []string{"list-application-keys"},
	Short:   "List application keys for this service account",
	Long: `List application keys for this service account
Documentation: https://docs.datadoghq.com/api/latest/service-accounts/#list-service-account-application-keys`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListApplicationKeysResponse
		var err error

		api := datadogV2.NewServiceAccountsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListServiceAccountApplicationKeys(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-service-account-application-keys")

		fmt.Println(cmdutil.FormatJSON(res, "application_keys"))
	},
}

func init() {

	Cmd.AddCommand(ListServiceAccountApplicationKeysCmd)
}
