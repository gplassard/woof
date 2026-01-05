package incident_teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateIncidentTeamCmd = &cobra.Command{
	Use:     "update-incident-team [team_id] [payload]",
	Aliases: []string{"update"},
	Short:   "Update an existing incident team",
	Long: `Update an existing incident team
Documentation: https://docs.datadoghq.com/api/latest/incident-teams/#update-incident-team`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentTeamResponse
		var err error

		var body datadogV2.IncidentTeamUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewIncidentTeamsApi(client.NewAPIClient())
		res, _, err = api.UpdateIncidentTeam(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-incident-team")

		cmd.Println(cmdutil.FormatJSON(res, "teams"))
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentTeamCmd)
}
