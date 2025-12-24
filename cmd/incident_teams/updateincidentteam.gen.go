package incident_teams

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateIncidentTeamCmd = &cobra.Command{
	Use:   "updateincidentteam [team_id]",
	Short: "Update an existing incident team",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewIncidentTeamsApi(client.NewAPIClient())
		res, _, err := api.UpdateIncidentTeam(client.NewContext(apiKey, appKey, site), args[0], datadogV2.IncidentTeamUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to updateincidentteam: %v", err)
		}

		cmdutil.PrintJSON(res, "incident_teams")
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentTeamCmd)
}
