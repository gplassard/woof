package incident_teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteIncidentTeamCmd = &cobra.Command{
	Use:     "delete-incident-team [team_id]",
	Aliases: []string{"delete"},
	Short:   "Delete an existing incident team",
	Long: `Delete an existing incident team
Documentation: https://docs.datadoghq.com/api/latest/incident-teams/#delete-incident-team`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewIncidentTeamsApi(client.NewAPIClient())
		_, err = api.DeleteIncidentTeam(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-incident-team")

	},
}

func init() {
	Cmd.AddCommand(DeleteIncidentTeamCmd)
}
