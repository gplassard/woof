package service_accounts

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateServiceAccountApplicationKeyCmd = &cobra.Command{
	Use:     "create-service-account-application-key [service_account_id]",
	Aliases: []string{"create-application-key"},
	Short:   "Create an application key for this service account",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceAccountsApi(client.NewAPIClient())
		res, _, err := api.CreateServiceAccountApplicationKey(client.NewContext(apiKey, appKey, site), args[0], datadogV2.ApplicationKeyCreateRequest{})
		cmdutil.HandleError(err, "failed to create-service-account-application-key")

		cmd.Println(cmdutil.FormatJSON(res, "application_keys"))
	},
}

func init() {
	Cmd.AddCommand(CreateServiceAccountApplicationKeyCmd)
}
