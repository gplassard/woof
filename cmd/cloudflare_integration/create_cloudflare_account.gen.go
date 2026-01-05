package cloudflare_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateCloudflareAccountCmd = &cobra.Command{
	Use: "create-cloudflare-account [payload]",

	Short: "Add Cloudflare account",
	Long: `Add Cloudflare account
Documentation: https://docs.datadoghq.com/api/latest/cloudflare-integration/#create-cloudflare-account`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CloudflareAccountResponse
		var err error

		var body datadogV2.CloudflareAccountCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCloudflareIntegrationApi(client.NewAPIClient())
		res, _, err = api.CreateCloudflareAccount(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-cloudflare-account")

		cmd.Println(cmdutil.FormatJSON(res, "cloudflare-accounts"))
	},
}

func init() {
	Cmd.AddCommand(CreateCloudflareAccountCmd)
}
