package teams

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateTeamLinkCmd = &cobra.Command{
	Use:   "create-team-link [team_id]",
	Aliases: []string{ "create-link", },
	Short: "Create a team link",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.CreateTeamLink(client.NewContext(apiKey, appKey, site), args[0], datadogV2.TeamLinkCreateRequest{})
		cmdutil.HandleError(err, "failed to create-team-link")

		cmdutil.PrintJSON(res, "team_links")
	},
}

func init() {
	Cmd.AddCommand(CreateTeamLinkCmd)
}
