package microsoft_teams_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateTenantBasedHandleCmd = &cobra.Command{
	Use:   "create_tenant_based_handle",
	Short: "Create tenant-based handle",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		res, _, err := api.CreateTenantBasedHandle(client.NewContext(apiKey, appKey, site), datadogV2.MicrosoftTeamsCreateTenantBasedHandleRequest{})
		if err != nil {
			log.Fatalf("failed to create_tenant_based_handle: %v", err)
		}

		cmdutil.PrintJSON(res, "tenant-based-handle")
	},
}

func init() {
	Cmd.AddCommand(CreateTenantBasedHandleCmd)
}
