package teams

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteTeamConnectionsCmd = &cobra.Command{
	Use:   "delete_team_connections",
	Short: "Delete team connections",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		_, err := api.DeleteTeamConnections(client.NewContext(apiKey, appKey, site), datadogV2.TeamConnectionDeleteRequest{})
		if err != nil {
			log.Fatalf("failed to delete_team_connections: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteTeamConnectionsCmd)
}
