package incident_teams

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListIncidentTeamsCmd = &cobra.Command{
	Use:     "list-incident-teams",
	Aliases: []string{"list"},
	Short:   "Get a list of all incident teams",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentTeamsApi(client.NewAPIClient())
		res, _, err := api.ListIncidentTeams(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-incident-teams")

		cmd.Println(cmdutil.FormatJSON(res, "teams"))
	},
}

func init() {
	Cmd.AddCommand(ListIncidentTeamsCmd)
}
