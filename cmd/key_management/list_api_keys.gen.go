package key_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAPIKeysCmd = &cobra.Command{
	Use: "list-api-keys",

	Short: "Get all API keys",
	Long: `Get all API keys
Documentation: https://docs.datadoghq.com/api/latest/key-management/#list-api-keys`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.APIKeysResponse
		var err error

		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListAPIKeys(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-api-keys")

		cmd.Println(cmdutil.FormatJSON(res, "api_keys"))
	},
}

func init() {

	Cmd.AddCommand(ListAPIKeysCmd)
}
