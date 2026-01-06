package okta_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateOktaAccountCmd = &cobra.Command{
	Use: "update-okta-account [account_id]",

	Short: "Update Okta account",
	Long: `Update Okta account
Documentation: https://docs.datadoghq.com/api/latest/okta-integration/#update-okta-account`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.OktaAccountResponse
		var err error

		var body datadogV2.OktaAccountUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewOktaIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateOktaAccount(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-okta-account")

		cmd.Println(cmdutil.FormatJSON(res, "okta-accounts"))
	},
}

func init() {

	UpdateOktaAccountCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateOktaAccountCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateOktaAccountCmd)
}
