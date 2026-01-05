package microsoft_teams_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetTenantBasedHandleCmd = &cobra.Command{
	Use: "get-tenant-based-handle [handle_id]",

	Short: "Get tenant-based handle information",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		res, _, err := api.GetTenantBasedHandle(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-tenant-based-handle")

		cmd.Println(cmdutil.FormatJSON(res, "tenant-based-handle"))
	},
}

func init() {
	Cmd.AddCommand(GetTenantBasedHandleCmd)
}
