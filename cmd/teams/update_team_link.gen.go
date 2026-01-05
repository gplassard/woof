package teams

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateTeamLinkCmd = &cobra.Command{
	Use:     "update-team-link [team_id] [link_id]",
	Aliases: []string{"update-link"},
	Short:   "Update a team link",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.UpdateTeamLink(client.NewContext(apiKey, appKey, site), args[0], args[1], datadogV2.TeamLinkCreateRequest{})
		cmdutil.HandleError(err, "failed to update-team-link")

		cmd.Println(cmdutil.FormatJSON(res, "team_links"))
	},
}

func init() {
	Cmd.AddCommand(UpdateTeamLinkCmd)
}
