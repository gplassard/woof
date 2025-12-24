package microsoft_teams_integration

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetChannelByNameCmd = &cobra.Command{
	Use:   "getchannelbyname [tenant_name] [team_name] [channel_name]",
	Short: "Get channel information by name",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		res, _, err := api.GetChannelByName(client.NewContext(apiKey, appKey, site), args[0], args[1], args[2])
		if err != nil {
			log.Fatalf("failed to getchannelbyname: %v", err)
		}

		cmdutil.PrintJSON(res, "microsoft_teams_integration")
	},
}

func init() {
	Cmd.AddCommand(GetChannelByNameCmd)
}
