package incident_teams

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteIncidentTeamCmd = &cobra.Command{
	Use:   "deleteincidentteam [team_id]",
	Short: "Delete an existing incident team",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewIncidentTeamsApi(client.NewAPIClient())
		_, err := api.DeleteIncidentTeam(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to deleteincidentteam: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteIncidentTeamCmd)
}
