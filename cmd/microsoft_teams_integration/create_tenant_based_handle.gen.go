package microsoft_teams_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateTenantBasedHandleCmd = &cobra.Command{
	Use: "create-tenant-based-handle [payload]",

	Short: "Create tenant-based handle",
	Long: `Create tenant-based handle
Documentation: https://docs.datadoghq.com/api/latest/microsoft-teams-integration/#create-tenant-based-handle`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MicrosoftTeamsTenantBasedHandleResponse
		var err error

		var body datadogV2.MicrosoftTeamsCreateTenantBasedHandleRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		res, _, err = api.CreateTenantBasedHandle(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-tenant-based-handle")

		cmd.Println(cmdutil.FormatJSON(res, "tenant-based-handle"))
	},
}

func init() {
	Cmd.AddCommand(CreateTenantBasedHandleCmd)
}
