package teams

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateTeamLinkCmd = &cobra.Command{
	Use:   "updateteamlink [team_id] [link_id]",
	Short: "Update a team link",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.UpdateTeamLink(client.NewContext(apiKey, appKey, site), args[0], args[1], datadogV2.TeamLinkCreateRequest{})
		if err != nil {
			log.Fatalf("failed to updateteamlink: %v", err)
		}

		cmdutil.PrintJSON(res, "teams")
	},
}

func init() {
	Cmd.AddCommand(UpdateTeamLinkCmd)
}
