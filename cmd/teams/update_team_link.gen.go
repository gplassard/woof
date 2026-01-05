package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateTeamLinkCmd = &cobra.Command{
	Use:     "update-team-link [team_id] [link_id] [payload]",
	Aliases: []string{"update-link"},
	Short:   "Update a team link",
	Args:    cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamLinkResponse
		var err error

		var body datadogV2.TeamLinkCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err = api.UpdateTeamLink(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to update-team-link")

		cmd.Println(cmdutil.FormatJSON(res, "team_links"))
	},
}

func init() {
	Cmd.AddCommand(UpdateTeamLinkCmd)
}
