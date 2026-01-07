package service_accounts

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateServiceAccountApplicationKeyCmd = &cobra.Command{
	Use:     "create-service-account-application-key [service_account_id]",
	Aliases: []string{"create-application-key"},
	Short:   "Create an application key for this service account",
	Long: `Create an application key for this service account
Documentation: https://docs.datadoghq.com/api/latest/service-accounts/#create-service-account-application-key`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ApplicationKeyResponse
		var err error

		var body datadogV2.ApplicationKeyCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewServiceAccountsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateServiceAccountApplicationKey(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-service-account-application-key")

		fmt.Println(cmdutil.FormatJSON(res, "application_keys"))
	},
}

func init() {

	CreateServiceAccountApplicationKeyCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateServiceAccountApplicationKeyCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateServiceAccountApplicationKeyCmd)
}
