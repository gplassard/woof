package teams

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateTeamLinkCmd = &cobra.Command{
	Use:   "createteamlink [team_id]",
	Short: "Create a team link",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.CreateTeamLink(client.NewContext(apiKey, appKey, site), args[0], datadogV2.TeamLinkCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createteamlink: %v", err)
		}

		cmdutil.PrintJSON(res, "teams")
	},
}

func init() {
	Cmd.AddCommand(CreateTeamLinkCmd)
}
