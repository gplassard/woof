package cloudflare_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateCloudflareAccountCmd = &cobra.Command{
	Use: "update-cloudflare-account [account_id] [payload]",

	Short: "Update Cloudflare account",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.CloudflareAccountUpdateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCloudflareIntegrationApi(client.NewAPIClient())
		res, _, err := api.UpdateCloudflareAccount(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-cloudflare-account")

		cmd.Println(cmdutil.FormatJSON(res, "cloudflare-accounts"))
	},
}

func init() {
	Cmd.AddCommand(UpdateCloudflareAccountCmd)
}
