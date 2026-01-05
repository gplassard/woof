package key_management

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetAPIKeyCmd = &cobra.Command{
	Use: "get-api-key [api_key_id]",

	Short: "Get API key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		res, _, err := api.GetAPIKey(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-api-key")

		cmd.Println(cmdutil.FormatJSON(res, "api_keys"))
	},
}

func init() {
	Cmd.AddCommand(GetAPIKeyCmd)
}
