package key_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteAPIKeyCmd = &cobra.Command{
	Use: "delete-api-key [api_key_id]",

	Short: "Delete an API key",
	Long: `Delete an API key
Documentation: https://docs.datadoghq.com/api/latest/key-management/#delete-api-key`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteAPIKey(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-api-key")

	},
}

func init() {

	Cmd.AddCommand(DeleteAPIKeyCmd)
}
