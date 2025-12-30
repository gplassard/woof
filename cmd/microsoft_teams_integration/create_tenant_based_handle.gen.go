package microsoft_teams_integration

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateTenantBasedHandleCmd = &cobra.Command{
	Use: "create-tenant-based-handle",

	Short: "Create tenant-based handle",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		res, _, err := api.CreateTenantBasedHandle(client.NewContext(apiKey, appKey, site), datadogV2.MicrosoftTeamsCreateTenantBasedHandleRequest{})
		cmdutil.HandleError(err, "failed to create-tenant-based-handle")

		cmd.Println(cmdutil.FormatJSON(res, "tenant-based-handle"))
	},
}

func init() {
	Cmd.AddCommand(CreateTenantBasedHandleCmd)
}
