package teams

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListTeamConnectionsCmd = &cobra.Command{
	Use:   "list-team-connections",
	Short: "List team connections",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.ListTeamConnections(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-team-connections: %v", err)
		}

		cmdutil.PrintJSON(res, "team_connection")
	},
}

func init() {
	Cmd.AddCommand(ListTeamConnectionsCmd)
}
