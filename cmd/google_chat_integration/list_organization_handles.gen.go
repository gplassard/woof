package google_chat_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListOrganizationHandlesCmd = &cobra.Command{
	Use: "list-organization-handles [organization_binding_id]",

	Short: "Get all organization handles",
	Long: `Get all organization handles
Documentation: https://docs.datadoghq.com/api/latest/google-chat-integration/#list-organization-handles`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GoogleChatOrganizationHandlesResponse
		var err error

		api := datadogV2.NewGoogleChatIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListOrganizationHandles(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-organization-handles")

		fmt.Println(cmdutil.FormatJSON(res, "google-chat-organization-handle"))
	},
}

func init() {

	Cmd.AddCommand(ListOrganizationHandlesCmd)
}
