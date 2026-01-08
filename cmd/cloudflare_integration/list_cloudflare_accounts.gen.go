package cloudflare_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCloudflareAccountsCmd = &cobra.Command{
	Use: "list-cloudflare-accounts",

	Short: "List Cloudflare accounts",
	Long: `List Cloudflare accounts
Documentation: https://docs.datadoghq.com/api/latest/cloudflare-integration/#list-cloudflare-accounts`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CloudflareAccountsResponse
		var err error

		api := datadogV2.NewCloudflareIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListCloudflareAccounts(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-cloudflare-accounts")

		fmt.Println(cmdutil.FormatJSON(res, "cloudflare_account"))
	},
}

func init() {

	Cmd.AddCommand(ListCloudflareAccountsCmd)
}
