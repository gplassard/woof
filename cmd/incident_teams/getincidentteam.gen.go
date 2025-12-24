package incident_teams

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetIncidentTeamCmd = &cobra.Command{
	Use:   "getincidentteam [team_id]",
	Short: "Get details of an incident team",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewIncidentTeamsApi(client.NewAPIClient())
		res, _, err := api.GetIncidentTeam(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getincidentteam: %v", err)
		}

		cmdutil.PrintJSON(res, "incident_teams")
	},
}

func init() {
	Cmd.AddCommand(GetIncidentTeamCmd)
}
