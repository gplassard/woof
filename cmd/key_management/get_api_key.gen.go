package key_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetAPIKeyCmd = &cobra.Command{
	Use: "get-api-key [api_key_id]",

	Short: "Get API key",
	Long: `Get API key
Documentation: https://docs.datadoghq.com/api/latest/key-management/#get-api-key`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.APIKeyResponse
		var err error

		optionalParams := datadogV2.NewGetAPIKeyOptionalParameters()

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetAPIKey(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to get-api-key")

		cmd.Println(cmdutil.FormatJSON(res, "api_keys"))
	},
}

func init() {

	GetAPIKeyCmd.Flags().String("include", "", "Comma separated list of resource paths for related resources to include in the response. Supported resource paths are 'created_by' and 'modified_by'.")

	Cmd.AddCommand(GetAPIKeyCmd)
}
