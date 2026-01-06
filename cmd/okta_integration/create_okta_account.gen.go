package okta_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateOktaAccountCmd = &cobra.Command{
	Use: "create-okta-account",

	Short: "Add Okta account",
	Long: `Add Okta account
Documentation: https://docs.datadoghq.com/api/latest/okta-integration/#create-okta-account`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.OktaAccountResponse
		var err error

		var body datadogV2.OktaAccountRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewOktaIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateOktaAccount(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-okta-account")

		cmd.Println(cmdutil.FormatJSON(res, "okta-accounts"))
	},
}

func init() {

	CreateOktaAccountCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateOktaAccountCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateOktaAccountCmd)
}
