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
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		res, _, err := api.UpdateFastlyAccount(client.NewContext(apiKey, appKey, site), args[0], datadogV2.FastlyAccountUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-fastly-account")

		cmd.Println(cmdutil.FormatJSON(res, "fastly-accounts"))
	},
}

func init() {
	Cmd.AddCommand(UpdateFastlyAccountCmd)
}
