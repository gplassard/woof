package teams

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateTeamConnectionsCmd = &cobra.Command{
	Use:   "create-team-connections",
	Short: "Create team connections",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.CreateTeamConnections(client.NewContext(apiKey, appKey, site), datadogV2.TeamConnectionCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create-team-connections: %v", err)
		}

		cmdutil.PrintJSON(res, "team_connection")
	},
}

func init() {
	Cmd.AddCommand(CreateTeamConnectionsCmd)
}
