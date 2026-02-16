package microsoft_teams_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListTenantBasedHandlesCmd = &cobra.Command{
	Use: "list-tenant-based-handles",

	Short: "Get all tenant-based handles",
	Long: `Get all tenant-based handles
Documentation: https://docs.datadoghq.com/api/latest/microsoft-teams-integration/#list-tenant-based-handles`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MicrosoftTeamsTenantBasedHandlesResponse
		var err error

		optionalParams := datadogV2.NewListTenantBasedHandlesOptionalParameters()

		if cmd.Flags().Changed("tenant-id") {
			val, _ := cmd.Flags().GetString("tenant-id")
			optionalParams.WithTenantId(val)
		}

		if cmd.Flags().Changed("name") {
			val, _ := cmd.Flags().GetString("name")
			optionalParams.WithName(val)
		}

		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListTenantBasedHandles(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-tenant-based-handles")

		fmt.Println(cmdutil.FormatJSON(res, "ms-teams-tenant-based-handle-info"))
	},
}

func init() {

	ListTenantBasedHandlesCmd.Flags().String("tenant-id", "", "Your tenant id.")

	ListTenantBasedHandlesCmd.Flags().String("name", "", "Your tenant-based handle name.")

	Cmd.AddCommand(ListTenantBasedHandlesCmd)
}
