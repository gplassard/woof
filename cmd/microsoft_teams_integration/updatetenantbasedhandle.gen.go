package microsoft_teams_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateTenantBasedHandleCmd = &cobra.Command{
	Use:   "updatetenantbasedhandle [handle_id]",
	Short: "Update tenant-based handle",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		res, _, err := api.UpdateTenantBasedHandle(client.NewContext(apiKey, appKey, site), args[0], datadogV2.MicrosoftTeamsUpdateTenantBasedHandleRequest{})
		if err != nil {
			log.Fatalf("failed to updatetenantbasedhandle: %v", err)
		}

		cmdutil.PrintJSON(res, "microsoft_teams_integration")
	},
}

func init() {
	Cmd.AddCommand(UpdateTenantBasedHandleCmd)
}
