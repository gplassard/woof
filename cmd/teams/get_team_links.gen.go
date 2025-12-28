package teams

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetTeamLinksCmd = &cobra.Command{
	Use:   "get-team-links [team_id]",
	Short: "Get links for a team",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.GetTeamLinks(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get-team-links: %v", err)
		}

		cmdutil.PrintJSON(res, "team_links")
	},
}

func init() {
	Cmd.AddCommand(GetTeamLinksCmd)
}
