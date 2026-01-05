package microsoft_teams_integration

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetChannelByNameCmd = &cobra.Command{
	Use: "get-channel-by-name [tenant_name] [team_name] [channel_name]",

	Short: "Get channel information by name",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		res, _, err := api.GetChannelByName(client.NewContext(apiKey, appKey, site), args[0], args[1], args[2])
		cmdutil.HandleError(err, "failed to get-channel-by-name")

		cmd.Println(cmdutil.FormatJSON(res, "ms-teams-channel-info"))
	},
}

func init() {
	Cmd.AddCommand(GetChannelByNameCmd)
}
