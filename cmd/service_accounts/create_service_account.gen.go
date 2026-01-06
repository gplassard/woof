package service_accounts

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateServiceAccountCmd = &cobra.Command{
	Use:     "create-service-account",
	Aliases: []string{"create"},
	Short:   "Create a service account",
	Long: `Create a service account
Documentation: https://docs.datadoghq.com/api/latest/service-accounts/#create-service-account`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UserResponse
		var err error

		var body datadogV2.ServiceAccountCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewServiceAccountsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateServiceAccount(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-service-account")

		cmd.Println(cmdutil.FormatJSON(res, "users"))
	},
}

func init() {

	CreateServiceAccountCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateServiceAccountCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateServiceAccountCmd)
}
