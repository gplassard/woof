package google_chat_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetOrganizationHandleCmd = &cobra.Command{
	Use: "get-organization-handle [organization_binding_id] [handle_id]",

	Short: "Get organization handle",
	Long: `Get organization handle
Documentation: https://docs.datadoghq.com/api/latest/google-chat-integration/#get-organization-handle`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GoogleChatOrganizationHandleResponse
		var err error

		api := datadogV2.NewGoogleChatIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetOrganizationHandle(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to get-organization-handle")

		fmt.Println(cmdutil.FormatJSON(res, "google-chat-organization-handle"))
	},
}

func init() {

	Cmd.AddCommand(GetOrganizationHandleCmd)
}
