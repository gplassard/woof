package cloudflare_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetCloudflareAccountCmd = &cobra.Command{
	Use: "get-cloudflare-account [account_id]",

	Short: "Get Cloudflare account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewCloudflareIntegrationApi(client.NewAPIClient())
		res, _, err := api.GetCloudflareAccount(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-cloudflare-account")

		cmd.Println(cmdutil.FormatJSON(res, "cloudflare-accounts"))
	},
}

func init() {
	Cmd.AddCommand(GetCloudflareAccountCmd)
}
