package key_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateAPIKeyCmd = &cobra.Command{
	Use: "create-api-key [payload]",

	Short: "Create an API key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.APIKeyCreateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		res, _, err := api.CreateAPIKey(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-api-key")

		cmd.Println(cmdutil.FormatJSON(res, "api_keys"))
	},
}

func init() {
	Cmd.AddCommand(CreateAPIKeyCmd)
}
