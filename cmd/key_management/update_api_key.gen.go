package key_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateAPIKeyCmd = &cobra.Command{
	Use: "update-api-key [api_key_id]",

	Short: "Edit an API key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		res, _, err := api.UpdateAPIKey(client.NewContext(apiKey, appKey, site), args[0], datadogV2.APIKeyUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-api-key")

		cmd.Println(cmdutil.FormatJSON(res, "api_keys"))
	},
}

func init() {
	Cmd.AddCommand(UpdateAPIKeyCmd)
}
