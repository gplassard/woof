package teams

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteTeamLinkCmd = &cobra.Command{
	Use:   "deleteteamlink [team_id] [link_id]",
	Short: "Remove a team link",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		_, err := api.DeleteTeamLink(client.NewContext(apiKey, appKey, site), args[0], args[1])
		if err != nil {
			log.Fatalf("failed to deleteteamlink: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteTeamLinkCmd)
}
