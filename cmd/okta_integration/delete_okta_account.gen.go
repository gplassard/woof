package okta_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteOktaAccountCmd = &cobra.Command{
	Use: "delete-okta-account [account_id]",

	Short: "Delete Okta account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewOktaIntegrationApi(client.NewAPIClient())
		_, err := api.DeleteOktaAccount(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-okta-account")

	},
}

func init() {
	Cmd.AddCommand(DeleteOktaAccountCmd)
}
