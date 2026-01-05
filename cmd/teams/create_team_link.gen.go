package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateTeamLinkCmd = &cobra.Command{
	Use:     "create-team-link [team_id] [payload]",
	Aliases: []string{"create-link"},
	Short:   "Create a team link",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.TeamLinkCreateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.CreateTeamLink(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-team-link")

		cmd.Println(cmdutil.FormatJSON(res, "team_links"))
	},
}

func init() {
	Cmd.AddCommand(CreateTeamLinkCmd)
}
