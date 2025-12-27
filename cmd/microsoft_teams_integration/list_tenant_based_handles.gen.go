package microsoft_teams_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListTenantBasedHandlesCmd = &cobra.Command{
	Use:   "list_tenant_based_handles",
	Short: "Get all tenant-based handles",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		res, _, err := api.ListTenantBasedHandles(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list_tenant_based_handles: %v", err)
		}

		cmdutil.PrintJSON(res, "microsoft_teams_integration")
	},
}

func init() {
	Cmd.AddCommand(ListTenantBasedHandlesCmd)
}
