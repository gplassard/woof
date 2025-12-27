package incident_teams

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetIncidentTeamCmd = &cobra.Command{
	Use:   "get_incident_team [team_id]",
	Short: "Get details of an incident team",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentTeamsApi(client.NewAPIClient())
		res, _, err := api.GetIncidentTeam(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get_incident_team: %v", err)
		}

		cmdutil.PrintJSON(res, "incident_teams")
	},
}

func init() {
	Cmd.AddCommand(GetIncidentTeamCmd)
}
