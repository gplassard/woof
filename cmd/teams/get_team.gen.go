package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetTeamCmd = &cobra.Command{
	Use:     "get-team [team_id]",
	Aliases: []string{"get"},
	Short:   "Get a team",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.GetTeam(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-team")

		cmd.Println(cmdutil.FormatJSON(res, "team"))
	},
}

func init() {
	Cmd.AddCommand(GetTeamCmd)
}
