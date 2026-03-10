package google_chat_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSpaceByDisplayNameCmd = &cobra.Command{
	Use: "get-space-by-display-name [domain_name] [space_display_name]",

	Short: "Get space information by display name",
	Long: `Get space information by display name
Documentation: https://docs.datadoghq.com/api/latest/google-chat-integration/#get-space-by-display-name`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GoogleChatAppNamedSpaceResponse
		var err error

		api := datadogV2.NewGoogleChatIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetSpaceByDisplayName(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to get-space-by-display-name")

		fmt.Println(cmdutil.FormatJSON(res, "google-chat-app-named-space"))
	},
}

func init() {

	Cmd.AddCommand(GetSpaceByDisplayNameCmd)
}
