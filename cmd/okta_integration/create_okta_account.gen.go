package okta_integration

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateOktaAccountCmd = &cobra.Command{
	Use: "create-okta-account",

	Short: "Add Okta account",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOktaIntegrationApi(client.NewAPIClient())
		res, _, err := api.CreateOktaAccount(client.NewContext(apiKey, appKey, site), datadogV2.OktaAccountRequest{})
		cmdutil.HandleError(err, "failed to create-okta-account")

		cmd.Println(cmdutil.FormatJSON(res, "okta-accounts"))
	},
}

func init() {
	Cmd.AddCommand(CreateOktaAccountCmd)
}
