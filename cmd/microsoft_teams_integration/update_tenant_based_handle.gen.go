package microsoft_teams_integration

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateTenantBasedHandleCmd = &cobra.Command{
	Use: "update-tenant-based-handle [handle_id]",

	Short: "Update tenant-based handle",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		res, _, err := api.UpdateTenantBasedHandle(client.NewContext(apiKey, appKey, site), args[0], datadogV2.MicrosoftTeamsUpdateTenantBasedHandleRequest{})
		cmdutil.HandleError(err, "failed to update-tenant-based-handle")

		cmd.Println(cmdutil.FormatJSON(res, "tenant-based-handle"))
	},
}

func init() {
	Cmd.AddCommand(UpdateTenantBasedHandleCmd)
}
