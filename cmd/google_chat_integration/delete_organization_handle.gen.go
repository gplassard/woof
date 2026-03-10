package google_chat_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteOrganizationHandleCmd = &cobra.Command{
	Use: "delete-organization-handle [organization_binding_id] [handle_id]",

	Short: "Delete organization handle",
	Long: `Delete organization handle
Documentation: https://docs.datadoghq.com/api/latest/google-chat-integration/#delete-organization-handle`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewGoogleChatIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteOrganizationHandle(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to delete-organization-handle")

	},
}

func init() {

	Cmd.AddCommand(DeleteOrganizationHandleCmd)
}
