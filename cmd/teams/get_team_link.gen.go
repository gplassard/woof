package teams

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetTeamLinkCmd = &cobra.Command{
	Use:   "get-team-link [team_id] [link_id]",
	Aliases: []string{ "get-link", },
	Short: "Get a team link",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.GetTeamLink(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to get-team-link")

		cmdutil.PrintJSON(res, "team_links")
	},
}

func init() {
	Cmd.AddCommand(GetTeamLinkCmd)
}
