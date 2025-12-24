package incident_teams

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListIncidentTeamsCmd = &cobra.Command{
	Use:   "listincidentteams",
	Short: "Get a list of all incident teams",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewIncidentTeamsApi(client.NewAPIClient())
		res, _, err := api.ListIncidentTeams(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listincidentteams: %v", err)
		}

		cmdutil.PrintJSON(res, "incident_teams")
	},
}

func init() {
	Cmd.AddCommand(ListIncidentTeamsCmd)
}
