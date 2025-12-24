package cloudflare_integration

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateCloudflareAccountCmd = &cobra.Command{
	Use:   "createcloudflareaccount",
	Short: "Add Cloudflare account",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCloudflareIntegrationApi(client.NewAPIClient())
		res, _, err := api.CreateCloudflareAccount(client.NewContext(apiKey, appKey, site), datadogV2.CloudflareAccountCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createcloudflareaccount: %v", err)
		}

		cmdutil.PrintJSON(res, "cloudflare_integration")
	},
}

func init() {
	Cmd.AddCommand(CreateCloudflareAccountCmd)
}
