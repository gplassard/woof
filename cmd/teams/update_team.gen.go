package teams

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateTeamCmd = &cobra.Command{
	Use:   "update-team [team_id]",
	Short: "Update a team",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.UpdateTeam(client.NewContext(apiKey, appKey, site), args[0], datadogV2.TeamUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to update-team: %v", err)
		}

		cmdutil.PrintJSON(res, "team")
	},
}

func init() {
	Cmd.AddCommand(UpdateTeamCmd)
}
