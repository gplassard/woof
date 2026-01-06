package microsoft_teams_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateTenantBasedHandleCmd = &cobra.Command{
	Use: "update-tenant-based-handle [handle_id]",

	Short: "Update tenant-based handle",
	Long: `Update tenant-based handle
Documentation: https://docs.datadoghq.com/api/latest/microsoft-teams-integration/#update-tenant-based-handle`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MicrosoftTeamsTenantBasedHandleResponse
		var err error

		var body datadogV2.MicrosoftTeamsUpdateTenantBasedHandleRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateTenantBasedHandle(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-tenant-based-handle")

		cmd.Println(cmdutil.FormatJSON(res, "tenant-based-handle"))
	},
}

func init() {

	UpdateTenantBasedHandleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateTenantBasedHandleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateTenantBasedHandleCmd)
}
