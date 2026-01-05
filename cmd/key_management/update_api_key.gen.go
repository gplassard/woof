package key_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateAPIKeyCmd = &cobra.Command{
	Use: "update-api-key [api_key_id] [payload]",

	Short: "Edit an API key",
	Long: `Edit an API key
Documentation: https://docs.datadoghq.com/api/latest/key-management/#update-api-key`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.APIKeyResponse
		var err error

		var body datadogV2.APIKeyUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		res, _, err = api.UpdateAPIKey(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-api-key")

		cmd.Println(cmdutil.FormatJSON(res, "api_keys"))
	},
}

func init() {
	Cmd.AddCommand(UpdateAPIKeyCmd)
}
