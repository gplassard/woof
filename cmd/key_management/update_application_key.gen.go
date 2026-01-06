package key_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateApplicationKeyCmd = &cobra.Command{
	Use: "update-application-key [app_key_id]",

	Short: "Edit an application key",
	Long: `Edit an application key
Documentation: https://docs.datadoghq.com/api/latest/key-management/#update-application-key`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ApplicationKeyResponse
		var err error

		var body datadogV2.ApplicationKeyUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateApplicationKey(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-application-key")

		cmd.Println(cmdutil.FormatJSON(res, "application_keys"))
	},
}

func init() {

	UpdateApplicationKeyCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateApplicationKeyCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateApplicationKeyCmd)
}
