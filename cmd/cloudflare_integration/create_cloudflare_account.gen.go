package cloudflare_integration

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCloudflareAccountCmd = &cobra.Command{
	Use: "create-cloudflare-account",

	Short: "Add Cloudflare account",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudflareIntegrationApi(client.NewAPIClient())
		res, _, err := api.CreateCloudflareAccount(client.NewContext(apiKey, appKey, site), datadogV2.CloudflareAccountCreateRequest{})
		cmdutil.HandleError(err, "failed to create-cloudflare-account")

		cmd.Println(cmdutil.FormatJSON(res, "cloudflare-accounts"))
	},
}

func init() {
	Cmd.AddCommand(CreateCloudflareAccountCmd)
}
