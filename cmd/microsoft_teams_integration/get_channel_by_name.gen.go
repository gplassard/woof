package microsoft_teams_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetChannelByNameCmd = &cobra.Command{
	Use:   "get_channel_by_name [tenant_name] [team_name] [channel_name]",
	Short: "Get channel information by name",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		res, _, err := api.GetChannelByName(client.NewContext(apiKey, appKey, site), args[0], args[1], args[2])
		if err != nil {
			log.Fatalf("failed to get_channel_by_name: %v", err)
		}

		cmdutil.PrintJSON(res, "ms-teams-channel-info")
	},
}

func init() {
	Cmd.AddCommand(GetChannelByNameCmd)
}
