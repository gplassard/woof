package key_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateAPIKeyCmd = &cobra.Command{
	Use: "create-api-key",

	Short: "Create an API key",
	Long: `Create an API key
Documentation: https://docs.datadoghq.com/api/latest/key-management/#create-api-key`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.APIKeyResponse
		var err error

		var body datadogV2.APIKeyCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateAPIKey(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-api-key")

		fmt.Println(cmdutil.FormatJSON(res, "a_p_i_key"))
	},
}

func init() {

	CreateAPIKeyCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateAPIKeyCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateAPIKeyCmd)
}
