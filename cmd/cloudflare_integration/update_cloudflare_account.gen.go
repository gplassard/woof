package cloudflare_integration

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateCloudflareAccountCmd = &cobra.Command{
	Use:   "update-cloudflare-account [account_id]",
	
	Short: "Update Cloudflare account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudflareIntegrationApi(client.NewAPIClient())
		res, _, err := api.UpdateCloudflareAccount(client.NewContext(apiKey, appKey, site), args[0], datadogV2.CloudflareAccountUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-cloudflare-account")

		cmdutil.PrintJSON(res, "cloudflare-accounts")
	},
}

func init() {
	Cmd.AddCommand(UpdateCloudflareAccountCmd)
}
