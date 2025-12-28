package incident_teams

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateIncidentTeamCmd = &cobra.Command{
	Use:   "create_incident_team",
	Short: "Create a new incident team",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentTeamsApi(client.NewAPIClient())
		res, _, err := api.CreateIncidentTeam(client.NewContext(apiKey, appKey, site), datadogV2.IncidentTeamCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create_incident_team: %v", err)
		}

		cmdutil.PrintJSON(res, "teams")
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentTeamCmd)
}
