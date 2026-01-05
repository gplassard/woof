package cloudflare_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateCloudflareAccountCmd = &cobra.Command{
	Use: "update-cloudflare-account [account_id]",

	Short: "Update Cloudflare account",
	Long: `Update Cloudflare account
Documentation: https://docs.datadoghq.com/api/latest/cloudflare-integration/#update-cloudflare-account`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CloudflareAccountResponse
		var err error

		var body datadogV2.CloudflareAccountUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCloudflareIntegrationApi(client.NewAPIClient())
		res, _, err = api.UpdateCloudflareAccount(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-cloudflare-account")

		cmd.Println(cmdutil.FormatJSON(res, "cloudflare-accounts"))
	},
}

func init() {

	UpdateCloudflareAccountCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateCloudflareAccountCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateCloudflareAccountCmd)
}
