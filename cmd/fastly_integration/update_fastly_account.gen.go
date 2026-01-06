package fastly_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateFastlyAccountCmd = &cobra.Command{
	Use: "update-fastly-account [account_id]",

	Short: "Update Fastly account",
	Long: `Update Fastly account
Documentation: https://docs.datadoghq.com/api/latest/fastly-integration/#update-fastly-account`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FastlyAccountResponse
		var err error

		var body datadogV2.FastlyAccountUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateFastlyAccount(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-fastly-account")

		cmd.Println(cmdutil.FormatJSON(res, "fastly-accounts"))
	},
}

func init() {

	UpdateFastlyAccountCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateFastlyAccountCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateFastlyAccountCmd)
}
