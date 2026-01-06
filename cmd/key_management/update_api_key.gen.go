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
	Long: `Edit an API key
Documentation: https://docs.datadoghq.com/api/latest/key-management/#update-api-key`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.APIKeyResponse
		var err error

		var body datadogV2.APIKeyUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateAPIKey(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-api-key")

		cmd.Println(cmdutil.FormatJSON(res, "api_keys"))
	},
}

func init() {

	UpdateAPIKeyCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateAPIKeyCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateAPIKeyCmd)
}
