package incident_teams

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateIncidentTeamCmd = &cobra.Command{
	Use:     "update-incident-team [team_id]",
	Aliases: []string{"update"},
	Short:   "Update an existing incident team",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentTeamsApi(client.NewAPIClient())
		res, _, err := api.UpdateIncidentTeam(client.NewContext(apiKey, appKey, site), args[0], datadogV2.IncidentTeamUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-incident-team")

		cmd.Println(cmdutil.FormatJSON(res, "teams"))
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentTeamCmd)
}
