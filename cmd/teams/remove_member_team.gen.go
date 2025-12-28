package teams

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var RemoveMemberTeamCmd = &cobra.Command{
	Use:   "remove-member-team [super_team_id] [member_team_id]",
	Aliases: []string{ "remove-member", },
	Short: "Remove a member team",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		_, err := api.RemoveMemberTeam(client.NewContext(apiKey, appKey, site), args[0], args[1])
		if err != nil {
			log.Fatalf("failed to remove-member-team: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(RemoveMemberTeamCmd)
}
