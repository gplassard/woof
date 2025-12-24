package teams

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteTeamConnectionsCmd = &cobra.Command{
	Use:   "deleteteamconnections",
	Short: "Delete team connections",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		_, err := api.DeleteTeamConnections(client.NewContext(apiKey, appKey, site), datadogV2.TeamConnectionDeleteRequest{})
		if err != nil {
			log.Fatalf("failed to deleteteamconnections: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteTeamConnectionsCmd)
}
