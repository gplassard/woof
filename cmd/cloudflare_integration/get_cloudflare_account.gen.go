package cloudflare_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetCloudflareAccountCmd = &cobra.Command{
	Use:   "get_cloudflare_account [account_id]",
	Short: "Get Cloudflare account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudflareIntegrationApi(client.NewAPIClient())
		res, _, err := api.GetCloudflareAccount(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get_cloudflare_account: %v", err)
		}

		cmdutil.PrintJSON(res, "cloudflare_integration")
	},
}

func init() {
	Cmd.AddCommand(GetCloudflareAccountCmd)
}
