package incident_teams

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListIncidentTeamsCmd = &cobra.Command{
	Use:   "list_incident_teams",
	Short: "Get a list of all incident teams",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentTeamsApi(client.NewAPIClient())
		res, _, err := api.ListIncidentTeams(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list_incident_teams: %v", err)
		}

		cmdutil.PrintJSON(res, "teams")
	},
}

func init() {
	Cmd.AddCommand(ListIncidentTeamsCmd)
}
