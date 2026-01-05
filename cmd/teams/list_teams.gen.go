package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListTeamsCmd = &cobra.Command{
	Use:     "list-teams",
	Aliases: []string{"list"},
	Short:   "Get all teams",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamsResponse
		var err error

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err = api.ListTeams(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-teams")

		cmd.Println(cmdutil.FormatJSON(res, "team"))
	},
}

func init() {
	Cmd.AddCommand(ListTeamsCmd)
}
