package google_chat_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateOrganizationHandleCmd = &cobra.Command{
	Use: "create-organization-handle [organization_binding_id]",

	Short: "Create organization handle",
	Long: `Create organization handle
Documentation: https://docs.datadoghq.com/api/latest/google-chat-integration/#create-organization-handle`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GoogleChatOrganizationHandleResponse
		var err error

		var body datadogV2.GoogleChatCreateOrganizationHandleRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewGoogleChatIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateOrganizationHandle(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-organization-handle")

		fmt.Println(cmdutil.FormatJSON(res, "google-chat-organization-handle"))
	},
}

func init() {

	CreateOrganizationHandleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateOrganizationHandleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateOrganizationHandleCmd)
}
