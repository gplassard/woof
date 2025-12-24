package cloudflare_integration

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteCloudflareAccountCmd = &cobra.Command{
	Use:   "deletecloudflareaccount [account_id]",
	Short: "Delete Cloudflare account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCloudflareIntegrationApi(client.NewAPIClient())
		_, err := api.DeleteCloudflareAccount(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to deletecloudflareaccount: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteCloudflareAccountCmd)
}
