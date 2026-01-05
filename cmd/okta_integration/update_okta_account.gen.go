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
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOktaIntegrationApi(client.NewAPIClient())
		res, _, err := api.UpdateOktaAccount(client.NewContext(apiKey, appKey, site), args[0], datadogV2.OktaAccountUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-okta-account")

		cmd.Println(cmdutil.FormatJSON(res, "okta-accounts"))
	},
}

func init() {
	Cmd.AddCommand(UpdateOktaAccountCmd)
}
