package teams

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var SyncTeamsCmd = &cobra.Command{
	Use:   "sync_teams",
	Short: "Link Teams with GitHub Teams",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		_, err := api.SyncTeams(client.NewContext(apiKey, appKey, site), datadogV2.TeamSyncRequest{})
		if err != nil {
			log.Fatalf("failed to sync_teams: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(SyncTeamsCmd)
}
