package microsoft_teams_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListTenantBasedHandlesCmd = &cobra.Command{
	Use: "list-tenant-based-handles",

	Short: "Get all tenant-based handles",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		res, _, err := api.ListTenantBasedHandles(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-tenant-based-handles")

		cmd.Println(cmdutil.FormatJSON(res, "ms-teams-tenant-based-handle-info"))
	},
}

func init() {
	Cmd.AddCommand(ListTenantBasedHandlesCmd)
}
