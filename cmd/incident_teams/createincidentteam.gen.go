package incident_teams

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateIncidentTeamCmd = &cobra.Command{
	Use:   "createincidentteam",
	Short: "Create a new incident team",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewIncidentTeamsApi(client.NewAPIClient())
		res, _, err := api.CreateIncidentTeam(client.NewContext(apiKey, appKey, site), datadogV2.IncidentTeamCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createincidentteam: %v", err)
		}

		cmdutil.PrintJSON(res, "incident_teams")
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentTeamCmd)
}
